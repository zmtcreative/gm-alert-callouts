package alertcallouts

// AlertCallouts is an extension for Goldmark.

import (
	_ "embed"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	alertParser "github.com/ZMT-Creative/gm-alert-callouts/internal/parser"
	alertRenderer "github.com/ZMT-Creative/gm-alert-callouts/internal/renderer"
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
	Icons:          createIconsMap(alertCalloutsIconsGFM),
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
		opts.Icons = createIconsMap(alertCalloutsIconsGFM)
		opts.defaultIcons = constants.ICONS_GFM
	}
}

// UseGFMPlusIcons sets the icon map to the GFM Plus icon set.
func UseGFMPlusIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = createIconsMap(alertCalloutsIconsGFMPlus)
		opts.defaultIcons = constants.ICONS_GFM_PLUS
	}
}

// UseObsidianIcons sets the icon map to the Obsidian icon set.
func UseObsidianIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.Icons = createIconsMap(alertCalloutsIconsObsidian)
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

// createIconsMap creates a map of icon names to their SVG data from the given icon data string.
func createIconsMap(icondata string) map[string]string {
	iconmap := make(map[string]string)

	// Parse the embedded alert callouts data
	lines := strings.Split(icondata, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Check if it's an alias definition (contains ->)
		if strings.Contains(line, "->") {
			parts := strings.SplitN(line, "->", 2)
			if len(parts) == 2 {
				alias := strings.TrimSpace(parts[0])
				primary := strings.TrimSpace(parts[1])
				// Set alias to reference the primary icon (will be set after core icons are loaded)
				if svg, exists := iconmap[primary]; exists {
					iconmap[alias] = svg
				} else {
					// Store for later processing if primary doesn't exist yet
					// This handles the case where aliases are defined before their primary keys
					defer func(alias, primary string) {
						if svg, exists := iconmap[primary]; exists {
							iconmap[alias] = svg
						}
					}(alias, primary)
				}
			}
			continue
		}

		// Parse core icon definition (key|svg)
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			svg := strings.TrimSpace(parts[1])
			iconmap[key] = svg
		}
	}

	// Second pass to handle any aliases that couldn't be resolved in first pass
	lines = strings.Split(icondata, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "->") {
			parts := strings.SplitN(line, "->", 2)
			if len(parts) == 2 {
				alias := strings.TrimSpace(parts[0])
				primary := strings.TrimSpace(parts[1])
				if svg, exists := iconmap[primary]; exists {
					iconmap[alias] = svg
				}
			}
		}
	}

	return iconmap
}
