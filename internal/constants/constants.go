package constants

import (
	gast "github.com/yuin/goldmark/ast"
)

const (
	ICONS_NONE = iota
	ICONS_GFM_STRICT
	ICONS_GFM_WITH_ALIASES
	ICONS_GFM_PLUS
	ICONS_OBSIDIAN
)

// Node kinds for different alert components
var (
	// KindAlerts is the NodeKind for the alert block.
	KindAlerts = gast.NewNodeKind("Alerts")

	// KindAlertsHeader is the NodeKind for the alert header.
	KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

	// KindAlertsBody is the NodeKind for the alert body.
	KindAlertsBody = gast.NewNodeKind("AlertsBody")
)

