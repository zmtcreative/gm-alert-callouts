package details

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"

	"regexp"
	"strings"
)

type calloutParser struct{}

var defaultCalloutParser = &calloutParser{}

func NewCalloutParser() parser.BlockParser {
	return defaultCalloutParser
}

func (b *calloutParser) Trigger() []byte {
	return []byte{'>'}
}

var regex = regexp.MustCompile("\\[!(?P<kind>[\\w]+)\\](?P<closed>-{0,1})(?P<open>\\+{0,1})\\s*(?P<title>.*)")

func (b *calloutParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	line, _ := reader.PeekLine()

	if !regex.Match(line) {
		return nil, parser.NoChildren
	}

	match := regex.FindSubmatch(line)

	kind := match[1]
	closed := match[2]
	open := match[3]
	// title := match[4]

	callout := NewCallout()

	callout.SetAttributeString("kind", kind)
	callout.SetAttributeString("closed", len(closed) != 0)
	callout.SetAttributeString("open", len(open) != 0)
	// callout.SetAttributeString("title", title)

	i := strings.Index(string(line), "]")
	reader.Advance(i)

	return callout, parser.HasChildren
}

func (b *calloutParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, _ := reader.PeekLine()

	if line[0] == '>' {
		reader.Advance(1)
		return parser.Continue | parser.HasChildren
	}

	return parser.Close
}

func (b *calloutParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *calloutParser) CanInterruptParagraph() bool {
	return true
}

func (b *calloutParser) CanAcceptIndentedLine() bool {
	return false
}
