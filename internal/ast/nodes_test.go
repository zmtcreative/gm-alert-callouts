package ast

import (
	"testing"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
)

func TestAlertsNode(t *testing.T) {
	t.Run("NewAlerts creates node", func(t *testing.T) {
		node := NewAlerts()
		if node == nil {
			t.Fatal("NewAlerts() returned nil")
		}
	})

	t.Run("Alerts implements correct kind", func(t *testing.T) {
		node := NewAlerts()
		if node.Kind() != constants.KindAlerts {
			t.Errorf("Expected KindAlerts, got %v", node.Kind())
		}
	})

	t.Run("Alerts is based on BaseBlock", func(t *testing.T) {
		node := NewAlerts()
		// Check that it has the expected base behavior
		if node.Type() != gast.TypeBlock {
			t.Error("Expected Alerts to have block type")
		}
	})

	t.Run("Alerts dump does not panic", func(t *testing.T) {
		node := NewAlerts()
		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Dump() panicked: %v", r)
			}
		}()
		node.Dump([]byte("test"), 0)
	})
}

func TestAlertsHeaderNode(t *testing.T) {
	t.Run("NewAlertsHeader creates node", func(t *testing.T) {
		node := NewAlertsHeader()
		if node == nil {
			t.Fatal("NewAlertsHeader() returned nil")
		}
	})

	t.Run("AlertsHeader implements correct kind", func(t *testing.T) {
		node := NewAlertsHeader()
		if node.Kind() != constants.KindAlertsHeader {
			t.Errorf("Expected KindAlertsHeader, got %v", node.Kind())
		}
	})

	t.Run("AlertsHeader is based on BaseBlock", func(t *testing.T) {
		node := NewAlertsHeader()
		// Check that it has the expected base behavior
		if node.Type() != gast.TypeBlock {
			t.Error("Expected AlertsHeader to have block type")
		}
	})

	t.Run("AlertsHeader dump does not panic", func(t *testing.T) {
		node := NewAlertsHeader()
		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Dump() panicked: %v", r)
			}
		}()
		node.Dump([]byte("test"), 0)
	})
}

func TestAlertsBodyNode(t *testing.T) {
	t.Run("NewAlertsBody creates node", func(t *testing.T) {
		node := NewAlertsBody()
		if node == nil {
			t.Fatal("NewAlertsBody() returned nil")
		}
	})

	t.Run("AlertsBody implements correct kind", func(t *testing.T) {
		node := NewAlertsBody()
		if node.Kind() != constants.KindAlertsBody {
			t.Errorf("Expected KindAlertsBody, got %v", node.Kind())
		}
	})

	t.Run("AlertsBody is based on BaseBlock", func(t *testing.T) {
		node := NewAlertsBody()
		// Check that it has the expected base behavior
		if node.Type() != gast.TypeBlock {
			t.Error("Expected AlertsBody to have block type")
		}
	})

	t.Run("AlertsBody dump does not panic", func(t *testing.T) {
		node := NewAlertsBody()
		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Dump() panicked: %v", r)
			}
		}()
		node.Dump([]byte("test"), 0)
	})
}

func TestNodeHierarchy(t *testing.T) {
	t.Run("Nodes can form hierarchy", func(t *testing.T) {
		alerts := NewAlerts()
		header := NewAlertsHeader()
		body := NewAlertsBody()

		// Test that we can build a proper node hierarchy
		alerts.AppendChild(alerts, header)
		alerts.AppendChild(alerts, body)

		if alerts.ChildCount() != 2 {
			t.Errorf("Expected 2 children, got %d", alerts.ChildCount())
		}

		// Check first child
		firstChild := alerts.FirstChild()
		if firstChild == nil {
			t.Fatal("Expected first child to exist")
		}
		if firstChild.Kind() != constants.KindAlertsHeader {
			t.Error("Expected first child to be AlertsHeader")
		}

		// Check last child
		lastChild := alerts.LastChild()
		if lastChild == nil {
			t.Fatal("Expected last child to exist")
		}
		if lastChild.Kind() != constants.KindAlertsBody {
			t.Error("Expected last child to be AlertsBody")
		}
	})
}

func TestNodeKindUniqueness(t *testing.T) {
	t.Run("Each node kind is unique", func(t *testing.T) {
		alertsKind := constants.KindAlerts
		headerKind := constants.KindAlertsHeader
		bodyKind := constants.KindAlertsBody

		// Check that all kinds are different
		if alertsKind == headerKind {
			t.Error("KindAlerts and KindAlertsHeader should be different")
		}
		if alertsKind == bodyKind {
			t.Error("KindAlerts and KindAlertsBody should be different")
		}
		if headerKind == bodyKind {
			t.Error("KindAlertsHeader and KindAlertsBody should be different")
		}

		// Check that they are all valid NodeKinds
		if alertsKind == gast.NodeKind(0) {
			t.Error("KindAlerts should not be zero value")
		}
		if headerKind == gast.NodeKind(0) {
			t.Error("KindAlertsHeader should not be zero value")
		}
		if bodyKind == gast.NodeKind(0) {
			t.Error("KindAlertsBody should not be zero value")
		}
	})
}
