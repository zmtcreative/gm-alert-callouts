package summary

import (
	gast "github.com/yuin/goldmark/ast"
)

type CalloutHeader struct {
	gast.BaseBlock
}

func (n *CalloutHeader) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

var KindCalloutHeader = gast.NewNodeKind("CalloutHeader")

func (n *CalloutHeader) Kind() gast.NodeKind {
	return KindCalloutHeader
}

func NewCalloutHeader() *CalloutHeader {
	return &CalloutHeader{}
}
