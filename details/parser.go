package details

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"

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

func (b *calloutParser) process(reader text.Reader) (bool, int) {
  // This is slighlty modified code from https://github.com/yuin/goldmark.git
  // Originally written by Yusuke Inuzuka, licensed under MIT License

	line, _ := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())
	if w > 3 || pos >= len(line) || line[pos] != '>' {
		return false, 0
	}

	advance_by := 1

	if pos >= len(line) || line[pos+advance_by] == '\n' {
		return true, advance_by
	}
	if line[pos+advance_by] == ' ' || line[pos+advance_by] == '\t' {
		advance_by++
	}

	if line[pos+advance_by-1] == '\t' {
		reader.SetPadding(2)
	}

	return true, advance_by
}

func (b *calloutParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {

	line, _ := reader.PeekLine()
	match := regex.FindSubmatch(line)

	if !regex.Match(line) {
		return nil, parser.NoChildren
	}
	ok, _ := b.process(reader)
	if !ok {
		return nil, parser.NoChildren
	}

	kind := match[1]
	closed := match[2]
	open := match[3]

	callout := NewCallout()

	callout.SetAttributeString("kind", kind)
	callout.SetAttributeString("closed", len(closed) != 0)
	callout.SetAttributeString("open", len(open) != 0)

	i := strings.Index(string(line), "]")
	// here's the bug
	// here we don't account for shifts in process
	reader.Advance(i)

	return callout, parser.HasChildren
}

func (b *calloutParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	ok, advance_by := b.process(reader)
	if !ok {
		return parser.Close
	}

  reader.Advance(advance_by)

	return parser.Continue | parser.HasChildren
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
