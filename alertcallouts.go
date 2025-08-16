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

// AlertCallouts is a extension for the goldmark (backwards compatibility).
var AlertCallouts = &AlertCalloutsOptions{}

// Option is a functional option for configuring AlertCalloutsOptions.
type Option func(*AlertCalloutsOptions)

// WithIcons sets the icons map for alert callouts.
func WithIcons(icons map[string]string) Option {
	return func(opts *AlertCalloutsOptions) {
		opts.Icons = icons
	}
}

// WithDisableFolding disables the folding functionality for alert callouts.
func WithDisableFolding(disable bool) Option {
	return func(opts *AlertCalloutsOptions) {
		opts.DisableFolding = alertRenderer.DisableFolding(disable)
	}
}

// WithIcon adds a single icon to the icons map for alert callouts.
func WithIcon(kind, icon string) Option {
	return func(opts *AlertCalloutsOptions) {
		if opts.Icons == nil {
			opts.Icons = make(map[string]string)
		}
		opts.Icons[kind] = icon
	}
}

// NewAlertCallouts creates a new AlertCallouts extension with the given options.
// This follows the standard Goldmark extension initialization pattern.
func NewAlertCallouts(options ...Option) *AlertCalloutsOptions {
	opts := &AlertCalloutsOptions{
		Icons:          make(map[string]string),
		DisableFolding: false,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

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
