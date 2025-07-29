package details

import (
	"strings"

	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
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
	reg.Register(kinds.KindAlerts, r.renderAlerts)
}

func (r *AlertsHTMLRenderer) renderAlerts(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
    alertType := ""
    if t, ok := node.AttributeString("kind"); ok {
        if typeBytes, isBytes := t.([]uint8); isBytes {
            alertType = strings.ToLower(string(typeBytes))
        } else if typeStr, isStr := t.(string); isStr {
            alertType = strings.ToLower(typeStr)
        }
    }

    if entering {
        fmt.Fprintf(w, `<div class="gh-alert gh-alert-%s" data-callout="%s">`, alertType, alertType)
    } else {
        w.WriteString("</div>\n")
    }
    return gast.WalkContinue, nil
}