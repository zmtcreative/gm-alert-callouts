package summary

import (
	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
	gast "github.com/yuin/goldmark/ast"
)

type AlertsHeader struct {
	gast.BaseBlock
}

func (n *AlertsHeader) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// var KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

func (n *AlertsHeader) Kind() gast.NodeKind {
	return kinds.KindAlertsHeader
}

func NewAlertsHeader() *AlertsHeader {
	return &AlertsHeader{}
}
