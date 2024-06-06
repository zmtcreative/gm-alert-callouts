package summary

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
	"github.com/yuin/goldmark/text"

	"gitlab.com/staticnoise/goldmark-callout/details"
)

type calloutHeaderParser struct{}

var defaultCalloutHeaderParser = &calloutHeaderParser{}

func NewCalloutHeaderParser() parser.BlockParser {
	return defaultCalloutHeaderParser
}

func (b *calloutHeaderParser) Trigger() []byte {
	// end of Callout begining
	return []byte{']'}
}

func (b *calloutHeaderParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	// this is always the first child of KindCallout
	if parent.ChildCount() != 0 || parent.Kind() != details.KindCallout {
		return nil, parser.NoChildren
	}

	// ]
	reader.Advance(1)

  next := reader.Peek() 
  if next == '-' {
    reader.Advance(1)
  }

	_, segment := reader.Position()
	line, _ := reader.PeekLine()

	w, _ := util.IndentWidth(line, reader.LineOffset())
	reader.Advance(w)

	_, segment = reader.Position()
	line, _ = reader.PeekLine()

	if len(line) > 0 && line[len(line)-1] == '\n' {
		segment.Stop = segment.Stop - 1
	}

	segments := text.Segments{}
	segments.Append(segment)

	paragraph := gast.NewParagraph()
	paragraph.SetLines(&segments)

	callout := NewCalloutHeader()
	callout.AppendChild(callout, paragraph)

	return callout, parser.NoChildren
}

func (b *calloutHeaderParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (b *calloutHeaderParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *calloutHeaderParser) CanInterruptParagraph() bool {
	return false
}

func (b *calloutHeaderParser) CanAcceptIndentedLine() bool {
	return true
}
