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
type IsFoldable bool

type AlertsHeaderHTMLRenderer struct {
	html.Config
	Icons
	IsFoldable
	titleCaser cases.Caser
}

func NewAlertsHeaderHTMLRendererWithIcons(icons Icons, isFoldable IsFoldable, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config:     html.NewConfig(),
		Icons:      icons,
		IsFoldable: isFoldable,
		titleCaser: cases.Title(language.English, cases.Compact),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func NewAlertsHeaderHTMLRenderer(isFoldable IsFoldable, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config: html.NewConfig(),
		IsFoldable: isFoldable,
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
	if t, ok := node.AttributeString("shouldfold"); ok {
		shouldFold = bool(t.(bool))
	}

	startHTML := ""
	endHTML := ""

	if bool(r.IsFoldable) && shouldFold {
		startHTML = fmt.Sprintf(`<summary class="gh-alert-title callout-title">` + "\n")
		endHTML = "\n</summary>\n"
	} else {
		startHTML = `<div class="gh-alert-title callout-title">` + "\n"
		endHTML = "\n</div>\n"
	}

	if entering {
		w.WriteString(startHTML)
		var kind string = ""

		if t, ok := node.AttributeString("kind"); ok {
			kind = strings.ToLower(t.(string))
			icon, ok := r.Icons[kind]
			// iconOutput := ""
			if ok {
				w.WriteString(icon)
				// iconOutput = fmt.Sprintf(`<div class="callout-icon icon-%s">` + icon + `</div>`, kind)
				// Check if the kind indicates no icon should be rendered.
				// if it's not a "no icon" kind, we can try to find a default icon.
			} else if !constants.IsNoIconKind(kind) {
				for _, v := range []string{"note", "info", "default"} {
					icon, ok = r.Icons[v]
					if ok {
						w.WriteString(icon)
						// iconOutput = fmt.Sprintf(`<div class="callout-icon icon-%s">` + icon + `</div>`, v)
						break
					}
				}
			}
			// if iconOutput != "" {
			// 	w.WriteString(iconOutput)
			// }
			w.WriteString(`<p class="callout-title-text">`)
			if _, ok := node.AttributeString("title"); ok {
				// do nothing
			} else {
				w.WriteString(r.titleCaser.String(kind))
			}
		}
	} else {
		w.WriteString(`</p>`)
		w.WriteString(endHTML)
	}
	return gast.WalkContinue, nil
}
