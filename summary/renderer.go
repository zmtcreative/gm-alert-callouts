package summary

import (
	"strings"

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
		w.WriteString("<summary>\n")
		var kind string = ""
		if t, ok := node.AttributeString("kind"); ok {
			kind = string(t.(string))
			kind = t.(string)
			w.WriteString(strings.Title(kind))
		}

	} else {
		w.WriteString("\n</summary>\n<div class=\"callout-content\">\n")
	}
	return gast.WalkContinue, nil
}
