package alerts

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/ZMT-Creative/goldmark-gh-alerts/body"
	"github.com/ZMT-Creative/goldmark-gh-alerts/details"
	"github.com/ZMT-Creative/goldmark-gh-alerts/summary"
)

type GhAlerts struct {
	summary.Icons
}

// Meta is a extension for the goldmark.
var GhAlertsExtension = &GhAlerts{}

// Extend implements goldmark.Extender.
func (e *GhAlerts) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(details.NewAlertsParser(), 799),
			util.Prioritized(summary.NewAlertsHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(details.NewAlertsHTMLRenderer(), 0),
			util.Prioritized(summary.NewAlertsHeaderHTMLRendererWithIcons(e.Icons), 0),
            util.Prioritized(body.NewAlertsBodyHTMLRenderer(), 0),
		),
	)
}
