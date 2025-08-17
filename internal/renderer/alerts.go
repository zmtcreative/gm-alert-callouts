package renderer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type AlertsHTMLRenderer struct {
	html.Config
	FoldingEnabled
	DefaultIcons int
}

func NewAlertsHTMLRenderer(foldingEnabled FoldingEnabled, defaultIcons int, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHTMLRenderer{
		Config:         html.NewConfig(),
		FoldingEnabled: foldingEnabled,
		DefaultIcons:   defaultIcons,
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(constants.KindAlerts, r.renderAlerts)
}

func (r *AlertsHTMLRenderer) renderAlerts(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	alertType := ""
	rawTitle := ""
	if t, ok := node.AttributeString("kind"); ok {
		if typeBytes, isBytes := t.([]uint8); isBytes {
			alertType = strings.ToLower(string(typeBytes))
		} else if typeStr, isStr := t.(string); isStr {
			alertType = strings.ToLower(typeStr)
		}
		// Check if the alertType is "noicon" or one of its variants
		if constants.IsNoIconKind(alertType) {
			// If the alertType is "noicon", we want to use the "title" if it exists
			// If not, we can just use existing alertType as a fallback
			if rt, ok := node.AttributeString("title"); ok {
				if typeBytes, isBytes := rt.([]uint8); isBytes {
					rawTitle = strings.ToLower(string(typeBytes))
				} else if typeStr, isStr := rt.(string); isStr {
					rawTitle = strings.ToLower(typeStr)
				}
				// Create regular expressions for cleaning up the title for use in the HTML output
				reHTML := regexp.MustCompile(`<[^>]+>`)
				reWS := regexp.MustCompile(`\s+`)
				reMDInline := regexp.MustCompile(`\*\*|\*|__|_|~~|\\`)
				reMDLinks := regexp.MustCompile(`!?\[(.*?)\]\((.*?)\)`)
				reMDCode := regexp.MustCompile("`{1,3}[^`]*`{1,3}")
				// Clean up the raw title using the regular expressions
				rawTitle = reHTML.ReplaceAllString(rawTitle, "")
				rawTitle = reMDInline.ReplaceAllString(rawTitle, "")
				rawTitle = reMDLinks.ReplaceAllString(rawTitle, "")
				rawTitle = reMDCode.ReplaceAllString(rawTitle, "")
				rawTitle = strings.TrimSpace(rawTitle)
				rawTitle = reWS.ReplaceAllString(rawTitle, "-")
				// Set the alert type to the cleaned-up title
				alertType = rawTitle
			}
		}
	}

	open := " open"
	if t, ok := node.AttributeString("closed"); ok {
		if bool(t.(bool)) {
			open = ""
		}
	}
	shouldFold := false
	if t, ok := node.AttributeString("shouldfold"); ok {
		shouldFold = bool(t.(bool))
	}

	iconset := ""
	switch r.DefaultIcons {
	case constants.ICONS_GFM:
		iconset = " iconset-gfm"
	case constants.ICONS_GFM_PLUS:
		iconset = " iconset-gfmplus"
	case constants.ICONS_OBSIDIAN:
		iconset = " iconset-obsidian"
	}

	startHTML := ""
	endHTML := ""

	if bool(r.FoldingEnabled) && shouldFold {
		startHTML = fmt.Sprintf(`<details class="callout callout-foldable callout-%s%s" data-callout="%s"%s>`, alertType, iconset, alertType, open)
		endHTML = "\n</details>\n"
	} else {
		startHTML = fmt.Sprintf(`<div class="callout callout-%s%s" data-callout="%s">`, alertType, iconset, alertType)
		endHTML = "\n</div>\n"
	}

	if entering {
		// fmt.Fprintf(w, `<div class="gh-alert gh-alert-%s" data-callout="%s">`, alertType, alertType)
		w.WriteString(startHTML)
	} else {
		w.WriteString(endHTML)
	}
	return gast.WalkContinue, nil
}
