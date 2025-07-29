package body

import (
	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
	gast "github.com/yuin/goldmark/ast"
)

type AlertsBody struct {
    gast.BaseBlock
}

func (n *AlertsBody) Dump(source []byte, level int) {
    gast.DumpHelper(n, source, level, nil, nil)
}

// var KindAlertsBody = gast.NewNodeKind("AlertsBody")

func (n *AlertsBody) Kind() gast.NodeKind {
    return kinds.KindAlertsBody
}

func NewAlertsBody() *AlertsBody {
    return &AlertsBody{}
}