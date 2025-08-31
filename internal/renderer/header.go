package renderer

import (
	"fmt"
	"log"
	"strings"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	locale "github.com/jeandeaual/go-locale"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Icons map[string]string
type FoldingEnabled bool
type CustomAlertsEnabled bool

type AlertsHeaderHTMLRenderer struct {
	html.Config
	Icons               map[string]string
	FoldingEnabled      bool
	CustomAlertsEnabled bool
	DefaultIcons        int
	AllowNOICON         bool
	titleCaser          cases.Caser
}

func NewAlertsHeaderHTMLRendererWithIcons(icons Icons, foldingEnabled FoldingEnabled, defaultIcons int, customAlertsEnabled CustomAlertsEnabled, opts ...html.Option) renderer.NodeRenderer {
	return NewAlertsHeaderHTMLRenderer(icons, bool(foldingEnabled), defaultIcons, bool(customAlertsEnabled), true, opts...)
}

// NewAlertsHeaderHTMLRenderer is the constructor used during normal program operation.
// It gets the user's local and then calls the internal implementation 'newAlertsHeaderHTMLRenderer' below with all the necessary parameters
func NewAlertsHeaderHTMLRenderer(icons map[string]string, foldingEnabled bool, defaultIcons int, customAlertsEnabled bool, allowNOICON bool, opts ...html.Option) renderer.NodeRenderer {
	// Detect the user's OS-level locale.
	userLocale, err := locale.GetLocale()
	if err != nil {
		// This is unlikely, but provides a safe fallback.
		log.Println("Could not detect OS locale, falling back to Undetermined:", err)
		userLocale = "und"
	}

	// language.Parse is robust. It returns language.Undetermined on error,
	// which is a perfect default for title casing. The cases package will
	// handle this gracefully.
	tag, _ := language.Parse(userLocale)

	return newAlertsHeaderHTMLRenderer(icons, foldingEnabled, defaultIcons, customAlertsEnabled, allowNOICON, tag, opts...)
}

// newAlertsHeaderHTMLRenderer is an unexported constructor that allows injecting a language tag.
// This is the internal implementation used by the public constructors above and is essential for writing
// unit tests that can verify behavior across different languages.
func newAlertsHeaderHTMLRenderer(icons map[string]string, foldingEnabled bool, defaultIcons int, customAlertsEnabled bool, allowNOICON bool, tag language.Tag, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config:              html.NewConfig(),
		Icons:               icons,
		FoldingEnabled:      foldingEnabled,
		CustomAlertsEnabled: customAlertsEnabled,
		DefaultIcons:        defaultIcons,
		AllowNOICON:         allowNOICON,
		titleCaser:          cases.Title(tag, cases.Compact),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsHeaderHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(constants.KindAlertsHeader, r.renderAlertsHeader)
}

func (r *AlertsHeaderHTMLRenderer) renderAlertsHeader(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	shouldFold := false
	var kind string = ""
	var icon string = ""
	var noicon bool = false


	if t, ok := node.AttributeString("kind"); ok {
		kind = strings.ToLower(t.(string))
		icon = r.Icons[kind]
	}
	if t, ok := node.AttributeString("shouldfold"); ok {
		shouldFold = bool(t.(bool))
	}
	if t, ok := node.AttributeString("noicon"); ok {
		noicon = bool(t.(bool))
	}


	startHTML := ""
	endHTML := ""

	if r.FoldingEnabled && shouldFold {
		startHTML = fmt.Sprintf(`<summary class="callout-title">` + "\n")
		endHTML = "\n</summary>\n"
	} else {
		startHTML = `<div class="callout-title">` + "\n"
		endHTML = "\n</div>\n"
	}

	// if AllowNOICON is set AND the alert had the 'noicon-' or 'noicon_' prefix...
	if r.AllowNOICON && noicon {
		// We'll place an empty span here to represent the empty icon space, just in case we need to
		// use CSS on this spot in the output (not sure it's necessary, but just being thorough for now).
		startHTML += `<span class="callout-title-noicon" style="display: none;"></span>`
	} else {
		// if the icon value is not empty, use the icon
		// else if custom alerts are enabled, use a fallback icon from 'constants.FALLBACK_ICON_LIST'
		if icon != "" {
			startHTML += icon
		} else if r.CustomAlertsEnabled {
			found := false
			for _, v := range constants.FALLBACK_ICON_LIST {
				deficon, ok := r.Icons[v]
				if ok {
					startHTML += deficon
					found = true
					break
				}
			}
			// If all else fails, generate an empty span as a placeholder for the icon
			// Note: This should never happen with the built-in iconsets.
			//       However, for custom iconsets provided by the user, they may not use
			//       alert names that match the predefined set of fallbacks in '
			//       constants.FALLBACK_ICON_LIST'. To be consistent, I think
			//       SOMETHING should be output where the icon would go.
			if !found {
				startHTML += `<span class="callout-title-noicon" style="display: none;"></span>`
			}
		}
	}

	startHTML += `<p class="callout-title-text">`

	_, hasTitle := node.AttributeString("title")

	// If there is an icon or if custom alerts are enabled, render the kind or the title
	if icon != "" || r.CustomAlertsEnabled {
		// If title isn't set, use kind for the title
		// NOTE: if title IS set, it is rendered separately as a text node when we 'WalkContinue' at the end
		if !hasTitle {
			startHTML += r.titleCaser.String(kind)
		}
	} else {
		// If we've gotten here, this is an invalid callout
		// Note: Because the parser phase SHOULD catch any kind value that doesn't have
		//   an icon when CustomAlerts is disabled, this part should never run, but keeping it
		//   here for now.
		startHTML += `[!` + strings.ToUpper(kind) + `]`
		if hasTitle {
			startHTML += ` `
		}
	}

	if entering {
		w.WriteString(startHTML)
	} else {
		w.WriteString(`</p>`)
		w.WriteString(endHTML)
	}
	return gast.WalkContinue, nil
}
