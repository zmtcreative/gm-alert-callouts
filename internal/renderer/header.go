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

	if t, ok := node.AttributeString("kind"); ok {
		kind = strings.ToLower(t.(string))
		icon = r.Icons[kind]
	}
	if t, ok := node.AttributeString("shouldfold"); ok {
		shouldFold = bool(t.(bool))
	}

	startHTML := ""
	endHTML := ""
	// startNoCustomTitle := ""
	// endNoCustomTitle := ""

	if r.FoldingEnabled && shouldFold {
		startHTML = fmt.Sprintf(`<summary class="callout-title">` + "\n")
		endHTML = "\n</summary>\n"
	} else {
		startHTML = `<div class="callout-title">` + "\n"
		endHTML = "\n</div>\n"
	}

	if entering {
		w.WriteString(startHTML)
		// var kind string = ""


		// if t, ok := node.AttributeString("kind"); ok {
		// 	kind = strings.ToLower(t.(string))
			// icon, ok := r.Icons[kind]
			if icon != "" {
				w.WriteString(icon)
				// Check if the kind indicates no icon should be rendered.
				// if it's not a "no icon" kind, we can try to find a default icon.
			} else if r.CustomAlertsEnabled {
				if kind != "noicon" {
					for _, v := range []string{"note", "info", "tip", "question", "default", "icon", "svg"} {
						deficon, ok := r.Icons[v]
						if ok {
							w.WriteString(deficon)
							break
						}
					}
				}
			}
			// else if !r.CustomAlertsEnabled {
			// 	startNoCustomTitle = `[!`
			// 	endNoCustomTitle = `]`
			// }
			w.WriteString(`<p class="callout-title-text">`)

			if _, ok := node.AttributeString("title"); ok {
				if !r.CustomAlertsEnabled && (kind != "" && icon == "") {
					w.WriteString(`[!`)
					// w.WriteString(r.titleCaser.String(kind))
					w.WriteString(strings.ToUpper(kind))
					w.WriteString(`] `)
				}
				// do nothing
			} else {
				if !r.CustomAlertsEnabled && (kind != "" && icon == "") {
					w.WriteString(`[!`)
				}
				if !r.CustomAlertsEnabled && kind == "noicon" {
					w.WriteString(strings.ToUpper(kind))
				} else if kind != "noicon" {
					w.WriteString(r.titleCaser.String(kind))
				}
				if !r.CustomAlertsEnabled && (kind != "" && icon == "") {
					w.WriteString(`]`)
				}
			}

		// }
	} else {
		w.WriteString(`</p>`)
		w.WriteString(endHTML)
	}
	return gast.WalkContinue, nil
}
