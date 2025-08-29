package renderer

import (
	"fmt"
	"strings"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
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

func NewAlertsHeaderHTMLRenderer(icons map[string]string, foldingEnabled bool, defaultIcons int, customAlertsEnabled bool, allowNOICON bool, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config:              html.NewConfig(),
		Icons:               icons,
		FoldingEnabled:      foldingEnabled,
		CustomAlertsEnabled: customAlertsEnabled,
		DefaultIcons:        defaultIcons,
		AllowNOICON:         allowNOICON,
		titleCaser:          cases.Title(language.English, cases.Compact),
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

	// Create a decision (decide) variable for later
	var decide int = 0
	if _, ok := r.Icons["noicon"]; ok {
		decide += 1
	}
	if r.CustomAlertsEnabled {
		decide += 2
	}
	if r.AllowNOICON {
		decide += 4
	}

	if t, ok := node.AttributeString("kind"); ok {
		kind = strings.ToLower(t.(string))
		icon = r.Icons[kind]
	}
	if t, ok := node.AttributeString("shouldfold"); ok {
		shouldFold = bool(t.(bool))
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


	if kind == "noicon" {
		if icon != "" {
			startHTML += icon
		} else if r.AllowNOICON {
			startHTML += `<span style="display: none;"></span>`
		}
	} else {
		if icon != "" {
			startHTML += icon
		} else if r.CustomAlertsEnabled {
			for _, v := range []string{"default", "note", "info", "tip", "question", "icon", "svg"} {
				deficon, ok := r.Icons[v]
				if ok {
					startHTML += deficon
					break
				}
			}
		}
	}
	startHTML += `<p class="callout-title-text">`

	_, hasTitle := node.AttributeString("title")
	if kind == "noicon" {
		// Based on the decision (decide) made earlier, only 0 or 2 will show the kind as invalid
		// All other values indicate the recognized NOICON callout value and don't render
		//   any output here
		if decide == 0 || decide == 2 {
			startHTML += `[!` + strings.ToUpper(kind) + `]`
			if hasTitle {
				startHTML += ` `
			}
		}
	} else {
		// If there is an icon or if custom alerts are enabled, render the kind or the title
		if icon != "" || r.CustomAlertsEnabled {
			// If title isn't set, use kind for the title
			// NOTE: if title IS set, it is rendered separately as a text node when we 'WalkContinue' at the end
			if !hasTitle {
				startHTML += r.titleCaser.String(kind)
			}
		} else {
			// If we've gotten here, this is an invalid callout
			startHTML += `[!` + strings.ToUpper(kind) + `]`
			if hasTitle {
				startHTML += ` `
			}
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
