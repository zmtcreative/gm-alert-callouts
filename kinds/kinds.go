package kinds

import gast "github.com/yuin/goldmark/ast"

// KindAlerts is the NodeKind for the alert block.
var KindAlerts = gast.NewNodeKind("Alerts")

// KindAlertsHeader is the NodeKind for the alert header.
var KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

// KindAlertsBody is the NodeKind for the alert body.
var KindAlertsBody = gast.NewNodeKind("AlertsBody")