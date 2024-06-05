package summary

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type CalloutHeaderHTMLRenderer struct {
	html.Config
}

func NewCalloutHeaderHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &CalloutHeaderHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *CalloutHeaderHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindCalloutHeader, r.renderCalloutHeader)
}

func (r *CalloutHeaderHTMLRenderer) renderCalloutHeader(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {

	if entering {
		w.WriteString(`<summary>
`)
	} else {
		w.WriteString(`</summary>
`)
	}
	return gast.WalkContinue, nil
}
