package constants

import (
	gast "github.com/yuin/goldmark/ast"
)

const (
	ICONS_NONE = iota
	ICONS_GFM
	ICONS_HYBRID
	ICONS_OBSIDIAN
)

var FALLBACK_ICON_LIST = []string{"default", "icon", "custom", "note", "info"}

// Node kinds for different alert components
var (
	// KindAlerts is the NodeKind for the alert block.
	KindAlerts = gast.NewNodeKind("Alerts")

	// KindAlertsHeader is the NodeKind for the alert header.
	KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

	// KindAlertsBody is the NodeKind for the alert body.
	KindAlertsBody = gast.NewNodeKind("AlertsBody")
)

