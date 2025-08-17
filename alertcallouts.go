package alertcallouts

// AlertCallouts is an extension for Goldmark.

import (
	_ "embed"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	alertParser "github.com/ZMT-Creative/gm-alert-callouts/internal/parser"
	alertRenderer "github.com/ZMT-Creative/gm-alert-callouts/internal/renderer"
	utils "github.com/ZMT-Creative/gm-alert-callouts/internal/utilities"
)

//go:embed assets/alertcallouts-gfmplus.icons
var alertCalloutsIconsGFMPlus string

//go:embed assets/alertcallouts-gfm.icons
var alertCalloutsIconsGFM string

//go:embed assets/alertcallouts-obsidian.icons
var alertCalloutsIconsObsidian string

var _ = alertCalloutsIconsGFMPlus
var _ = alertCalloutsIconsGFM
var _ = alertCalloutsIconsObsidian

type alertCalloutsOptions struct {
	alertRenderer.Icons
	alertRenderer.FoldingEnabled
	defaultIcons int
}

// AlertCallouts is a extension for Goldmark.
// This variable will initialize the extension with Folding Enabled and the basic GFM icon set
var AlertCallouts = &alertCalloutsOptions{
	Icons:          utils.CreateIconsMap(alertCalloutsIconsGFM),
	FoldingEnabled: true,
	defaultIcons:   constants.ICONS_GFM,
}

// Option is a functional option for configuring alertCalloutsOptions.
type Option func(*alertCalloutsOptions)

// WithIcons sets the icons map for alert callouts.
func WithIcons(icons map[string]string) Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = icons
	}
}

// WithIcon adds a single icon to the icons map for alert callouts.
func WithIcon(kind, icon string) Option {
	return func(opts *alertCalloutsOptions) {
		if opts.Icons == nil {
			opts.Icons = make(map[string]string)
		}
		opts.Icons[kind] = icon
	}
}

// UseGFMIcons sets the icon map to the GFM icon set.
func UseGFMIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = utils.CreateIconsMap(alertCalloutsIconsGFM)
		opts.defaultIcons = constants.ICONS_GFM
	}
}

// UseGFMPlusIcons sets the icon map to the GFM Plus icon set.
func UseGFMPlusIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = utils.CreateIconsMap(alertCalloutsIconsGFMPlus)
		opts.defaultIcons = constants.ICONS_GFM_PLUS
	}
}

// UseObsidianIcons sets the icon map to the Obsidian icon set.
func UseObsidianIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = utils.CreateIconsMap(alertCalloutsIconsObsidian)
		opts.defaultIcons = constants.ICONS_OBSIDIAN
	}
}

// WithFolding sets the folding functionality for alert callouts.
func WithFolding(enable bool) Option {
	return func(opts *alertCalloutsOptions) {
		opts.FoldingEnabled = alertRenderer.FoldingEnabled(enable)
	}
}

// NewAlertCallouts creates a new AlertCallouts extension with the given options.
// This follows the standard Goldmark extension initialization pattern.
func NewAlertCallouts(options ...Option) *alertCalloutsOptions {
	opts := &alertCalloutsOptions{
		Icons:          make(map[string]string),
		FoldingEnabled: true,
		defaultIcons:   constants.ICONS_NONE,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

// Extend implements goldmark.Extender.
func (e *alertCalloutsOptions) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(alertParser.NewAlertsParser(), 799),
			util.Prioritized(alertParser.NewAlertsHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(alertRenderer.NewAlertsHTMLRenderer(e.FoldingEnabled, e.defaultIcons), 0),
			util.Prioritized(alertRenderer.NewAlertsHeaderHTMLRendererWithIcons(e.Icons, e.FoldingEnabled, e.defaultIcons), 0),
			util.Prioritized(alertRenderer.NewAlertsBodyHTMLRenderer(), 0),
		),
	)
}

