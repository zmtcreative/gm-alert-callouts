package details

import (
	gast "github.com/yuin/goldmark/ast"
)

type Alerts struct {
	gast.BaseBlock
}

func (n *Alerts) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindAlerts = gast.NewNodeKind("Alerts")

func (n *Alerts) Kind() gast.NodeKind {
	return KindAlerts
}

func NewAlerts() *Alerts {
	return &Alerts{}
}
