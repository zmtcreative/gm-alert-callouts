package renderer

import (
	"fmt"
	"strings"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type AlertsHTMLRenderer struct {
	html.Config
	Icons               map[string]string
	FoldingEnabled      bool
	CustomAlertsEnabled bool
	DefaultIcons        int
	AllowNOICON         bool
}

func NewAlertsHTMLRenderer(icons map[string]string, foldingEnabled bool, defaultIcons int, customAlertsEnabled bool, allowNOICON bool, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHTMLRenderer{
		Config:              html.NewConfig(),
		Icons:               icons,
		FoldingEnabled:      foldingEnabled,
		DefaultIcons:        defaultIcons,
		CustomAlertsEnabled: customAlertsEnabled,
		AllowNOICON:         allowNOICON,
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
	var alertType = ""
	var icon string = ""

	if t, ok := node.AttributeString("kind"); ok {
		if typeBytes, isBytes := t.([]uint8); isBytes {
			alertType = strings.ToLower(string(typeBytes))
		} else if typeStr, isStr := t.(string); isStr {
			alertType = strings.ToLower(typeStr)
		}
		icon = r.Icons[alertType]
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
	case constants.ICONS_GFM_STRICT:
		iconset = " iconset-gfm"
	case constants.ICONS_GFM_WITH_ALIASES:
		iconset = " iconset-gfm"
	case constants.ICONS_GFM_PLUS:
		iconset = " iconset-gfmplus"
	case constants.ICONS_OBSIDIAN:
		iconset = " iconset-obsidian"
	}

	startHTML := ""
	endHTML := ""
	var _ = icon

	if r.FoldingEnabled && shouldFold {
		startHTML = fmt.Sprintf(`<details class="callout callout-foldable callout-%s%s" data-callout="%s"%s>`, alertType, iconset, alertType, open)
		endHTML = "\n</details>\n"
	} else {
		startHTML = fmt.Sprintf(`<div class="callout callout-%s%s" data-callout="%s">`, alertType, iconset, alertType)
		endHTML = "\n</div>\n"
	}

	if entering {
		w.WriteString(startHTML)
	} else {
		w.WriteString(endHTML)
	}
	return gast.WalkContinue, nil
}
