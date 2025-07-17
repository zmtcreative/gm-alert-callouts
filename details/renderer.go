package details

import (
	"strings"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"

	"fmt"
)

type AlertsHTMLRenderer struct {
	html.Config
}

func NewAlertsHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindAlerts, r.renderAlerts)
}

func (r *AlertsHTMLRenderer) renderAlerts(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	alertType := ""
	if t, ok := node.AttributeString("kind"); ok {
		alertType = strings.ToLower(string(t.([]uint8)))
	}

	start := fmt.Sprintf(`<div class="markdown-alert markdown-alert-%s" data-callout="%s">`, alertType, alertType)

	if entering {
		w.WriteString(start)
	} else {
		w.WriteString("</div>\n")
	}
	return gast.WalkContinue, nil
}
