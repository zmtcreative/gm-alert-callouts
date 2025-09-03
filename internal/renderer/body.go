package renderer

import (
	"github.com/zmtcreative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type AlertsBodyHTMLRenderer struct {
	html.Config
}

func NewAlertsBodyHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsBodyHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsBodyHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(constants.KindAlertsBody, r.renderAlertsBody)
}

func (r *AlertsBodyHTMLRenderer) renderAlertsBody(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		w.WriteString(`<div class="callout-body">`)
	} else {
		w.WriteString("</div>")
	}
	return gast.WalkContinue, nil
}
