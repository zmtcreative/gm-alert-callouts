package details

import (
	gast "github.com/yuin/goldmark/ast"
)

type Callout struct {
	gast.BaseBlock
}

func (n *Callout) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindCallout = gast.NewNodeKind("Callout")

func (n *Callout) Kind() gast.NodeKind {
	return KindCallout
}

func NewCallout() *Callout {
	return &Callout{}
}
