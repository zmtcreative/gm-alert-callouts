package constants

import (
	"testing"

	gast "github.com/yuin/goldmark/ast"
)

func TestIconConstants(t *testing.T) {
	t.Run("Icon constants are distinct", func(t *testing.T) {
		constants := []int{ICONS_NONE, ICONS_GFM, ICONS_GFM_PLUS, ICONS_OBSIDIAN}

		// Check that all constants are different
		for i := 0; i < len(constants); i++ {
			for j := i + 1; j < len(constants); j++ {
				if constants[i] == constants[j] {
					t.Errorf("Constants at positions %d and %d have the same value: %d", i, j, constants[i])
				}
			}
		}
	})

	t.Run("Icon constants have expected values", func(t *testing.T) {
		if ICONS_NONE != 0 {
			t.Errorf("Expected ICONS_NONE to be 0, got %d", ICONS_NONE)
		}

		if ICONS_GFM <= ICONS_NONE {
			t.Errorf("Expected ICONS_GFM to be greater than ICONS_NONE, got %d", ICONS_GFM)
		}

		if ICONS_GFM_PLUS <= ICONS_GFM {
			t.Errorf("Expected ICONS_GFM_PLUS to be greater than ICONS_GFM, got %d", ICONS_GFM_PLUS)
		}

		if ICONS_OBSIDIAN <= ICONS_GFM_PLUS {
			t.Errorf("Expected ICONS_OBSIDIAN to be greater than ICONS_GFM_PLUS, got %d", ICONS_OBSIDIAN)
		}
	})
}

func TestNodeKinds(t *testing.T) {
	t.Run("KindAlerts is valid", func(t *testing.T) {
		if KindAlerts == gast.NodeKind(0) {
			t.Error("KindAlerts should not be zero value")
		}
	})

	t.Run("KindAlertsHeader is valid", func(t *testing.T) {
		if KindAlertsHeader == gast.NodeKind(0) {
			t.Error("KindAlertsHeader should not be zero value")
		}
	})

	t.Run("KindAlertsBody is valid", func(t *testing.T) {
		if KindAlertsBody == gast.NodeKind(0) {
			t.Error("KindAlertsBody should not be zero value")
		}
	})

	t.Run("All node kinds are unique", func(t *testing.T) {
		kinds := []gast.NodeKind{KindAlerts, KindAlertsHeader, KindAlertsBody}

		for i := 0; i < len(kinds); i++ {
			for j := i + 1; j < len(kinds); j++ {
				if kinds[i] == kinds[j] {
					t.Errorf("NodeKinds at positions %d and %d are the same: %v", i, j, kinds[i])
				}
			}
		}
	})

	t.Run("Node kinds have expected names", func(t *testing.T) {
		if KindAlerts.String() != "Alerts" {
			t.Errorf("Expected KindAlerts name to be 'Alerts', got '%s'", KindAlerts.String())
		}

		if KindAlertsHeader.String() != "AlertsHeader" {
			t.Errorf("Expected KindAlertsHeader name to be 'AlertsHeader', got '%s'", KindAlertsHeader.String())
		}

		if KindAlertsBody.String() != "AlertsBody" {
			t.Errorf("Expected KindAlertsBody name to be 'AlertsBody', got '%s'", KindAlertsBody.String())
		}
	})
}

func TestIconConstantUsage(t *testing.T) {
	t.Run("Can use constants for comparison", func(t *testing.T) {
		// Test that constants can be used in switch statements and comparisons
		testValue := ICONS_GFM

		var result string
		switch testValue {
		case ICONS_NONE:
			result = "none"
		case ICONS_GFM:
			result = "gfm"
		case ICONS_GFM_PLUS:
			result = "gfm_plus"
		case ICONS_OBSIDIAN:
			result = "obsidian"
		default:
			result = "unknown"
		}

		if result != "gfm" {
			t.Errorf("Expected 'gfm', got '%s'", result)
		}
	})

	t.Run("Constants work in conditional logic", func(t *testing.T) {
		if ICONS_GFM_PLUS == ICONS_NONE {
			t.Error("GFM_PLUS should not equal NONE")
		}

		if ICONS_OBSIDIAN < ICONS_GFM {
			t.Error("OBSIDIAN should be greater than GFM")
		}
	})
}
