package details

import (
	"strings"

	"github.com/ZMT-Creative/goldmark-gh-alerts/body"
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
        alertType = strings.ToLower(string(t.([]uint8)))
    }

    if entering {
        start := fmt.Sprintf(`<div class="gh-alert gh-alert-%s" data-callout="%s">`, alertType, alertType)
        w.WriteString(start)

        // Find the header and wrap subsequent children in a body node
        var header gast.Node
        var bodyChildren []gast.Node
        for c := node.FirstChild(); c != nil; {
            next := c.NextSibling()
            if c.Kind() == kinds.KindAlertsHeader {
                header = c
            } else {
                bodyChildren = append(bodyChildren, c)
            }
            c = next
        }

        // Re-parent children
        node.RemoveChildren(node)
        if header != nil {
            node.AppendChild(node, header)
        }

        if len(bodyChildren) > 0 {
            bodyNode := body.NewAlertsBody()
            for _, child := range bodyChildren {
                bodyNode.AppendChild(bodyNode, child)
            }
            node.AppendChild(node, bodyNode)
        }

    } else {
        w.WriteString("</div>\n")
    }
    return gast.WalkContinue, nil
}