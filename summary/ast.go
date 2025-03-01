package summary

import (
	gast "github.com/yuin/goldmark/ast"
)

type AlertsHeader struct {
	gast.BaseBlock
}

func (n *AlertsHeader) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

func (n *AlertsHeader) Kind() gast.NodeKind {
	return KindAlertsHeader
}

func NewAlertsHeader() *AlertsHeader {
	return &AlertsHeader{}
}
