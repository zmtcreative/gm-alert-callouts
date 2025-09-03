package parser

import (
	"regexp"
	"slices"
	"strings"

	"github.com/zmtcreative/gm-alert-callouts/internal/ast"
	"github.com/zmtcreative/gm-alert-callouts/internal/constants"
	"github.com/zmtcreative/gm-alert-callouts/internal/utilities"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type alertParser struct{
	IconList []string
	FoldingEnabled bool
	CustomAlertsEnabled bool
}

var defaultAlertsParser = &alertParser{}
var _ = defaultAlertsParser

func NewAlertsParser(iconList []string, foldingEnabled bool, customAlertsEnabled bool) parser.BlockParser {
	return &alertParser{
		IconList:            iconList,
		FoldingEnabled:      foldingEnabled,
		CustomAlertsEnabled: customAlertsEnabled,
	}
}

func (b *alertParser) Trigger() []byte {
	return []byte{'>'}
}

// Regex updated to support Unicode in <kind> value
var regex = regexp.MustCompile(`^\[!(?P<kind>\p{L}[\p{L}\p{N}_-]*)\](?:(?P<closed>-{0,1})|(?P<opened>[+]{0,1}))($|\s+(?P<title>.*))`)

func (b *alertParser) process(reader text.Reader) (bool, int) {
	// This is slightly modified code from https://github.com/yuin/goldmark.git
	// Originally written by Yusuke Inuzuka, licensed under MIT License

	line, _ := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())
	if w > 3 || pos >= len(line) || line[pos] != '>' {
		return false, 0
	}

	advanceBy := 1

	if pos+advanceBy >= len(line) || line[pos+advanceBy] == '\n' {
		return true, advanceBy
	}
	if line[pos+advanceBy] == ' ' || line[pos+advanceBy] == '\t' {
		advanceBy++
	}

	if line[pos+advanceBy-1] == '\t' {
		reader.SetPadding(2)
	}

	return true, advanceBy
}

func (b *alertParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	// check if we are inside of a block quote
	ok, advanceBy := b.process(reader)
	if !ok {
		return nil, parser.NoChildren
	}

	line, _ := reader.PeekLine()

	// empty blockquote
	if len(line) <= advanceBy {
		return nil, parser.NoChildren
	}

	// right after `>` and up to one space
	subline := line[advanceBy:]
	match := utilities.FindNamedMatches(regex, string(subline))

	// If no match found, this is not an alert
	if len(match["kind"]) == 0 {
		return nil, parser.NoChildren
	}

	kind := []uint8(match["kind"])
	closed := []uint8(match["closed"])
	title := []uint8(match["title"])
	opened := []uint8(match["opened"])

	// Set the 'shouldFold' variable:
	//   If the markdown uses either '-' or '+' for folding we assume the user wants the alert to be foldable.
	//   If neither '-' or '+' are used, we assume the alert is not meant to be foldable.
	//   We only need 'opened' here to check if the alert is meant to be foldable.
	//   The 'closed' variable is legacy and is used by existing code elsewhere (we're not tinkering with it -- yet)
	shouldFold := 1
	if (len(closed) == 0 && len(opened) == 0) {
		shouldFold = 0;
	}

	// Check for our hybrid feature of prefixing kind with 'noicon-' or 'noicon_'
	//   If we find the prefix, strip it, leaving the original kind without the prefix and set the attribute 'noicon' on the node
	//   We need to do this first so the later logic is looking at the kind/lckind without the prefix
	noicon := 0
	lckind := strings.ToLower(string(kind))
	if strings.HasPrefix(lckind, "noicon-") || strings.HasPrefix(lckind, "noicon_") {
		kind = []uint8(lckind[7:])
		lckind = lckind[7:]
		noicon = 1
	}

	// If CustomAlerts is not in use, disallow anything like:
	//   - kind doesn't have an icon
	//   - custom title not allowed
	//   - folding symbols (+ and -) not allowed
	if !b.CustomAlertsEnabled {
		if !(slices.Contains(b.IconList, lckind)) {
			// We'll reject any kind that isn't in the current IconList
			return nil, parser.NoChildren
		} else if len(title) > 0 {
			// GFM does not support custom titles, so including a custom title is disallowed
			return nil, parser.NoChildren
		} else if !b.FoldingEnabled && (len(closed) != 0 || len(opened) != 0) {
			// GFM does not support folding, so we should disallow even recognized kind values
			// with the folding symbols (+ and -) to conform to the way GitHub does its alerts
			return nil, parser.NoChildren
		}
	} else if !b.FoldingEnabled {
		// IF Folding is not enabled, but CustomAlerts IS enabled, modify shouldFold, closed, and opened
		//   so the CustomAlerts still work but folding symbols are silently ignored.
		// We don't want to prevent Custom Alert kinds just because the + or - is added.
		//   but we need to ignore folding unless Folding is explicitly enabled (this is a design decision)
		shouldFold = 0
		closed = nil
		opened = nil
	} else if len(kind) < 1 {
		return nil, parser.NoChildren
	}

	alert := ast.NewAlerts()

	alert.SetAttributeString("kind", kind)
	alert.SetAttributeString("closed", len(closed) != 0)
	alert.SetAttributeString("title", title)
	alert.SetAttributeString("shouldfold", shouldFold != 0)
	alert.SetAttributeString("noicon", noicon != 0)

	i := strings.Index(string(line), "]")
	if i >= 0 {
		reader.Advance(i)
	}

	return alert, parser.HasChildren
}

func (b *alertParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	ok, advanceBy := b.process(reader)
	if !ok {
		return parser.Close
	}

	reader.Advance(advanceBy)

	return parser.Continue | parser.HasChildren
}

func (b *alertParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// Restructure nodes to have a proper AlertsHeader and AlertsBody hierarchy
	var header gast.Node
	var bodyChildren []gast.Node

	for c := node.FirstChild(); c != nil; {
		next := c.NextSibling()
		if c.Kind() == constants.KindAlertsHeader {
			header = c
		} else {
			bodyChildren = append(bodyChildren, c)
		}
		c = next
	}

	// Re-parent children
	node.RemoveChildren(node)
	if header != nil {
		node.AppendChild(node, header)
	}

	if len(bodyChildren) > 0 {
		bodyNode := ast.NewAlertsBody()
		for _, child := range bodyChildren {
			bodyNode.AppendChild(bodyNode, child)
		}
		node.AppendChild(node, bodyNode)
	}
}

func (b *alertParser) CanInterruptParagraph() bool {
	return true
}

func (b *alertParser) CanAcceptIndentedLine() bool {
	return false
}
