package details

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"

	"fmt"
)

type CalloutHTMLRenderer struct {
	html.Config
}

func NewCalloutHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &CalloutHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *CalloutHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindCallout, r.renderCallout)
}

func (r *CalloutHTMLRenderer) renderCallout(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	var calloutType string = ""
	if t, ok := node.AttributeString("kind"); ok {
		calloutType = string(t.([]uint8))
	}

	open := " open"
	if t, ok := node.AttributeString("closed"); ok {
		if bool(t.(bool)) {
			open = ""
		}
	}

	start := fmt.Sprintf(`<details data-callout="%s"%s>
`, calloutType, open)

	if entering {
		w.WriteString(start)
	} else {
		w.WriteString("</div>\n</details>\n")
	}
	return gast.WalkContinue, nil
}
