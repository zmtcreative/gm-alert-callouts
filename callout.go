package callout

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"gitlab.com/staticnoise/goldmark-callout/details"
	"gitlab.com/staticnoise/goldmark-callout/summary"
)

type callout struct{}

// Meta is a extension for the goldmark.
var CalloutExtention = &callout{}

// Extend implements goldmark.Extender.
func (e *callout) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(details.NewCalloutParser(), 799),
			util.Prioritized(summary.NewCalloutHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(details.NewCalloutHTMLRenderer(), 0),
			util.Prioritized(summary.NewCalloutHeaderHTMLRenderer(), 0),
		),
	)
}
