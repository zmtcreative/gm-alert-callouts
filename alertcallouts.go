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

//go:embed assets/alertcallouts-gfm-strict.icons
var alertCalloutsIconsGFMStrict string

//go:embed assets/alertcallouts-gfm-with-aliases.icons
var alertCalloutsIconsGFMWithAliases string

//go:embed assets/alertcallouts-gfmplus.icons
var alertCalloutsIconsGFMPlus string

//go:embed assets/alertcallouts-obsidian.icons
var alertCalloutsIconsObsidian string

var _ = alertCalloutsIconsGFMStrict
var _ = alertCalloutsIconsGFMWithAliases
var _ = alertCalloutsIconsGFMPlus
var _ = alertCalloutsIconsObsidian

// Config holds all configuration options for alert callouts rendering.
// This struct is passed to renderer constructors to avoid long parameter lists
// and make it easy to add new options without breaking function signatures.
type Config struct {
	Icons               map[string]string // Icon map for different alert types
	FoldingEnabled      bool              // Whether folding functionality is enabled
	CustomAlertsEnabled bool              // Whether custom alert types are allowed
	DefaultIcons        int               // Which default icon set to use (constants.ICONS_*)
	AllowNOICON         bool              // Whether to allow NOICON alert types (example of new option)
}

type alertCalloutsOptions struct {
	config Config
}

// Option is a functional option for configuring alertCalloutsOptions.
type Option func(*alertCalloutsOptions)

// WithIcons sets the icons map for alert callouts.
func WithIcons(icons map[string]string) Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.Icons = icons
	}
}

// WithIcon adds a single icon to the icons map for alert callouts.
func WithIcon(kind, icon string) Option {
	return func(opts *alertCalloutsOptions) {
		if opts.config.Icons == nil {
			opts.config.Icons = make(map[string]string)
		}
		opts.config.Icons[kind] = icon
	}
}

// UseGFMIcons sets the icon map to the GFM (GitHub Flavored Markdown) icon set.
// DEPRECATED: Use UseGFMStrictIcons instead.
func UseGFMIcons() Option {
	return UseGFMStrictIcons()
}

func UseGFMStrictIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.Icons = utils.CreateIconsMap(alertCalloutsIconsGFMStrict)
		opts.config.DefaultIcons = constants.ICONS_GFM_STRICT
		opts.config.FoldingEnabled = false
		opts.config.CustomAlertsEnabled = false
		opts.config.AllowNOICON = false
	}
}

func UseGFMWithAliasesIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.Icons = utils.CreateIconsMap(alertCalloutsIconsGFMWithAliases)
		opts.config.DefaultIcons = constants.ICONS_GFM_WITH_ALIASES
		opts.config.FoldingEnabled = false
		opts.config.CustomAlertsEnabled = false
		opts.config.AllowNOICON = false
	}
}

// UseGFMPlusIcons sets the icon map to the GFM Plus icon set (essentially a melding of GFM and Obsidian).
func UseGFMPlusIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.Icons = utils.CreateIconsMap(alertCalloutsIconsGFMPlus)
		opts.config.DefaultIcons = constants.ICONS_GFM_PLUS
		opts.config.FoldingEnabled = true
		opts.config.CustomAlertsEnabled = true
		opts.config.AllowNOICON = true
		// opts.config.Icons["noicon"] = `<span style="display: none;"></span>`
	}
}

// UseObsidianIcons sets the icon map to the Obsidian-style icon set.
func UseObsidianIcons() Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.Icons = utils.CreateIconsMap(alertCalloutsIconsObsidian)
		opts.config.DefaultIcons = constants.ICONS_OBSIDIAN
		opts.config.FoldingEnabled = true
		opts.config.CustomAlertsEnabled = true
		opts.config.AllowNOICON = false
	}
}

// WithFolding sets the folding functionality for alert callouts.
func WithFolding(enable bool) Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.FoldingEnabled = enable
	}
}

func WithCustomAlerts(enable bool) Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.CustomAlertsEnabled = enable
	}
}

// WithAllowNOICON sets whether to allow NOICON alert types (example of new option).
func WithAllowNOICON(enable bool) Option {
	return func(opts *alertCalloutsOptions) {
		opts.config.AllowNOICON = enable
		if enable {
			opts.config.Icons["noicon"] = `<svg></svg>`
		}
	}
}

// CreateIconsMap creates a map of icon names to their SVG data from the given icon data string.
// This is a public wrapper around the internal utilities function, allowing users to create
// custom icon maps from their own icon data files.
func CreateIconsMap(iconData string) map[string]string {
	return utils.CreateIconsMap(iconData)
}

// AlertCallouts will initialize the extension with Folding Enabled and the basic GFM icon set
// This can be initialized using the `goldmark.WithExtensions(alertcallouts.AlertCallouts)` syntax
var AlertCallouts = NewAlertCallouts(
	UseGFMWithAliasesIcons(),
	WithFolding(true),
)

// NewAlertCallouts creates a new AlertCallouts extension with the given options.
// This follows the standard Goldmark extension initialization pattern.
func NewAlertCallouts(options ...Option) *alertCalloutsOptions {
	opts := &alertCalloutsOptions{
		config: Config{
			Icons:               make(map[string]string),
			FoldingEnabled:      true,
			CustomAlertsEnabled: true,
			DefaultIcons:        constants.ICONS_NONE,
			AllowNOICON:         true, // Default to true for backward compatibility
		},
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

// GetConfig returns the internal configuration for testing purposes.
// This method should not be used in production code.
func (e *alertCalloutsOptions) GetConfig() *Config {
	return &e.config
}

// GetIconKeys returns a slice of all icon keys from the Config.Icons map.
// This allows access to the icon names without the SVG data, useful for
// validation, searching, or listing available alert types without the overhead
// of passing large SVG strings.
func (c *Config) GetIconKeys() []string {
	if c.Icons == nil {
		return []string{}
	}

	keys := make([]string, 0, len(c.Icons))
	for key := range c.Icons {
		keys = append(keys, key)
	}

	return keys
}

// Extend implements goldmark.Extender.
func (e *alertCalloutsOptions) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(alertParser.NewAlertsParser(e.config.GetIconKeys(), e.config.CustomAlertsEnabled, e.config.AllowNOICON), 799),
			util.Prioritized(alertParser.NewAlertsHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(alertRenderer.NewAlertsHTMLRenderer(e.config.Icons, e.config.FoldingEnabled, e.config.DefaultIcons, e.config.CustomAlertsEnabled, e.config.AllowNOICON), 0),
			util.Prioritized(alertRenderer.NewAlertsHeaderHTMLRenderer(e.config.Icons, e.config.FoldingEnabled, e.config.DefaultIcons, e.config.CustomAlertsEnabled, e.config.AllowNOICON), 0),
			util.Prioritized(alertRenderer.NewAlertsBodyHTMLRenderer(), 0),
		),
	)
}

