package alertcallouts

// AlertCallouts is a extension for the goldmark.

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	alertParser "github.com/ZMT-Creative/gm-alert-callouts/internal/parser"
	alertRenderer "github.com/ZMT-Creative/gm-alert-callouts/internal/renderer"
)

type AlertCalloutsOptions struct {
	alertRenderer.Icons
	alertRenderer.DisableFolding
}

// Meta is a extension for the goldmark.
var AlertCallouts = &AlertCalloutsOptions{}

// Extend implements goldmark.Extender.
func (e *AlertCalloutsOptions) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(alertParser.NewAlertsParser(), 799),
			util.Prioritized(alertParser.NewAlertsHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(alertRenderer.NewAlertsHTMLRenderer(e.DisableFolding), 0),
			util.Prioritized(alertRenderer.NewAlertsHeaderHTMLRendererWithIcons(e.Icons, e.DisableFolding), 0),
			util.Prioritized(alertRenderer.NewAlertsBodyHTMLRenderer(), 0),
		),
	)
}
