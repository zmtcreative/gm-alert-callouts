package parser

import (
	"github.com/ZMT-Creative/gm-alert-callouts/internal/ast"
	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type alertHeaderParser struct{}

var defaultAlertsHeaderParser = &alertHeaderParser{}

func NewAlertsHeaderParser() parser.BlockParser {
	return defaultAlertsHeaderParser
}

func (b *alertHeaderParser) Trigger() []byte {
	// end of Alerts beginning
	return []byte{']'}
}

func (b *alertHeaderParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	// this is always the first child of KindAlerts
	if parent.ChildCount() != 0 || parent.Kind() != constants.KindAlerts {
		return nil, parser.NoChildren
	}

	// ]
	reader.Advance(1)

	next := reader.Peek()
	if next == '-' || next == '+' {
		reader.Advance(1)
	}

	_, segment := reader.Position()
	line, _ := reader.PeekLine()

	w, _ := util.IndentWidth(line, reader.LineOffset())
	reader.Advance(w)

	_, segment = reader.Position()
	line, _ = reader.PeekLine()

	// remove \n from the title of the alert
	if len(line) > 0 && line[len(line)-1] == '\n' {
		segment.Stop = segment.Stop - 1
	}

	var titleLine string = string(line)

	alert := ast.NewAlertsHeader()

	if segment.Len() != 0 {
		segments := text.Segments{}
		segments.Append(segment)

		paragraph := gast.NewTextBlock()
		paragraph.SetLines(&segments)

		alert.AppendChild(alert, paragraph)

		var kind string = ""
		if t, ok := parent.AttributeString("kind"); ok {
			kind = string(t.([]uint8))
			alert.SetAttributeString("kind", kind)
		}
		alert.SetAttributeString("title", titleLine)
	} else {
		var kind string = ""
		if t, ok := parent.AttributeString("kind"); ok {
			kind = string(t.([]uint8))
			alert.SetAttributeString("kind", kind)
		}
	}

	// Set the 'shouldfold' attribute if it is set in the parent
	// We'll need this value in the rendering process to determine which HTML elements to use.
	if t, ok := parent.AttributeString("shouldfold"); ok {
		alert.SetAttributeString("shouldfold", t)
	}
	// Set the 'noicon' attribute if it is set in the parent
	// We'll need this value in the rendering process to supress icon use
	if t, ok := parent.AttributeString("noicon"); ok {
		alert.SetAttributeString("noicon", t)
	}

	return alert, parser.NoChildren
}

func (b *alertHeaderParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (b *alertHeaderParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *alertHeaderParser) CanInterruptParagraph() bool {
	return false
}

func (b *alertHeaderParser) CanAcceptIndentedLine() bool {
	return true
}
