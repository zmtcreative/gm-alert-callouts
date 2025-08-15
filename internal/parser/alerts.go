package parser

import (
	"regexp"
	"strings"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/ast"
	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type alertParser struct{}

var defaultAlertsParser = &alertParser{}

func NewAlertsParser() parser.BlockParser {
	return defaultAlertsParser
}

func (b *alertParser) Trigger() []byte {
	return []byte{'>'}
}

var regex = regexp.MustCompile(`^\[!(?P<kind>[\w]+)\](?:(?P<closed>-{0,1})|(?P<opened>[+]{0,1}))($|\s+(?P<title>.*))`)

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
	match := constants.FindNamedMatches(regex, string(subline))

	// If no match found, this is not an alert
	if len(match["kind"]) == 0 {
		return nil, parser.NoChildren
	}

	kind := []uint8(match["kind"])
	closed := []uint8(match["closed"])
	opened := []uint8(match["opened"])
	title := []uint8(match["title"])
	shouldFold := 1

	if (len(closed) == 0 && len(opened) == 0) {
		shouldFold = 0;
	}

	alert := ast.NewAlerts()

	alert.SetAttributeString("kind", kind)
	alert.SetAttributeString("closed", len(closed) != 0)
	alert.SetAttributeString("opened", len(opened) != 0)
	alert.SetAttributeString("title", title)
	alert.SetAttributeString("shouldfold", shouldFold)
	// fmt.Println("Alert kind:", string(kind), " | closed:", len(closed) != 0, " | opened:", len(opened) != 0, " | title:", string(title), " | shouldfold:", shouldFold)

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
