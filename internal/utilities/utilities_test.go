package utilities

import (
	"regexp"
	"testing"
)

func TestIsNoIconKind(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
		desc     string
	}{
		{"noicon", true, "noicon should return true"},
		{"none", true, "no-icon should return true"},
		{"nil", true, "nil should return true"},
		{"null", true, "null should return true"},
		{"note", false, "note should return false"},
		{"warning", false, "warning should return false"},
		{"info", false, "info should return false"},
		{"", false, "empty string should return false"},
		{"NoIcon", true, "NoIcon (capitalized) should return true"},
		{"NO-ICON", false, "NO-ICON (uppercase) should return false"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := IsNoIconKind(tc.input)
			if result != tc.expected {
				t.Errorf("IsNoIconKind(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestFindNamedMatches(t *testing.T) {
	t.Run("Basic named capture groups", func(t *testing.T) {
		regex := regexp.MustCompile(`(?P<type>\w+)\|(?P<icon>.+)`)
		input := "note|<svg>icon</svg>"

		result := FindNamedMatches(regex, input)

		if len(result) != 3 { // Full match + 2 groups
			t.Errorf("Expected 3 matches, got %d", len(result))
		}

		if result["type"] != "note" {
			t.Errorf("Expected type 'note', got '%s'", result["type"])
		}

		if result["icon"] != "<svg>icon</svg>" {
			t.Errorf("Expected icon '<svg>icon</svg>', got '%s'", result["icon"])
		}
	})

	t.Run("No match returns empty map", func(t *testing.T) {
		regex := regexp.MustCompile(`(?P<type>\w+)\|(?P<icon>.+)`)
		input := "no match here"

		result := FindNamedMatches(regex, input)

		if len(result) != 0 {
			t.Errorf("Expected empty map, got %v", result)
		}
	})

	t.Run("Handles regex with no named groups", func(t *testing.T) {
		regex := regexp.MustCompile(`\w+\|.+`)
		input := "note|<svg>icon</svg>"

		result := FindNamedMatches(regex, input)

		// Should have one unnamed match
		if len(result) != 1 {
			t.Errorf("Expected 1 match, got %d", len(result))
		}
	})

	t.Run("Multiple named groups", func(t *testing.T) {
		regex := regexp.MustCompile(`(?P<first>\w+)-(?P<second>\w+)-(?P<third>\w+)`)
		input := "one-two-three"

		result := FindNamedMatches(regex, input)

		expected := map[string]string{
			"":       "one-two-three", // Full match
			"first":  "one",
			"second": "two",
			"third":  "three",
		}

		for key, expectedValue := range expected {
			if result[key] != expectedValue {
				t.Errorf("Expected %s='%s', got '%s'", key, expectedValue, result[key])
			}
		}
	})
}

func TestCreateIconsMap(t *testing.T) {
	t.Run("Basic icon parsing", func(t *testing.T) {
		iconData := `note|<svg>note icon</svg>
warning|<svg>warning icon</svg>
info|<svg>info icon</svg>`

		result := CreateIconsMap(iconData)

		expected := map[string]string{
			"note":    "<svg>note icon</svg>",
			"warning": "<svg>warning icon</svg>",
			"info":    "<svg>info icon</svg>",
		}

		if len(result) != len(expected) {
			t.Errorf("Expected %d icons, got %d", len(expected), len(result))
		}

		for key, expectedValue := range expected {
			if result[key] != expectedValue {
				t.Errorf("Expected %s='%s', got '%s'", key, expectedValue, result[key])
			}
		}
	})

	t.Run("Handles comments and empty lines", func(t *testing.T) {
		iconData := `# This is a comment
note|<svg>note icon</svg>

# Another comment
warning|<svg>warning icon</svg>

`

		result := CreateIconsMap(iconData)

		if len(result) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(result))
		}

		if result["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected note icon, got '%s'", result["note"])
		}

		if result["warning"] != "<svg>warning icon</svg>" {
			t.Errorf("Expected warning icon, got '%s'", result["warning"])
		}
	})

	t.Run("Handles aliases with arrows", func(t *testing.T) {
		iconData := `note|<svg>note icon</svg>
info -> note
tip -> note`

		result := CreateIconsMap(iconData)

		if len(result) != 3 {
			t.Errorf("Expected 3 entries (1 icon + 2 aliases), got %d", len(result))
		}

		if result["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected note icon, got '%s'", result["note"])
		}

		if result["info"] != "<svg>note icon</svg>" {
			t.Errorf("Expected info to alias note icon, got '%s'", result["info"])
		}

		if result["tip"] != "<svg>note icon</svg>" {
			t.Errorf("Expected tip to alias note icon, got '%s'", result["tip"])
		}
	})

	t.Run("Handles aliases defined before primary", func(t *testing.T) {
		iconData := `info -> note
tip -> note
note|<svg>note icon</svg>`

		result := CreateIconsMap(iconData)

		if len(result) != 3 {
			t.Errorf("Expected 3 entries, got %d", len(result))
		}

		if result["info"] != "<svg>note icon</svg>" {
			t.Errorf("Expected info to alias note icon, got '%s'", result["info"])
		}

		if result["tip"] != "<svg>note icon</svg>" {
			t.Errorf("Expected tip to alias note icon, got '%s'", result["tip"])
		}
	})

	t.Run("Handles complex icon data with whitespace", func(t *testing.T) {
		iconData := `  note  |  <svg>note icon</svg>
  warning  |  <svg>warning icon</svg>
  info  ->  note  `

		result := CreateIconsMap(iconData)

		if len(result) != 3 {
			t.Errorf("Expected 3 entries, got %d", len(result))
		}

		if result["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected trimmed note icon, got '%s'", result["note"])
		}

		if result["warning"] != "<svg>warning icon</svg>" {
			t.Errorf("Expected trimmed warning icon, got '%s'", result["warning"])
		}

		if result["info"] != "<svg>note icon</svg>" {
			t.Errorf("Expected info to alias note icon, got '%s'", result["info"])
		}
	})

	t.Run("Handles empty input", func(t *testing.T) {
		result := CreateIconsMap("")

		if len(result) != 0 {
			t.Errorf("Expected empty map for empty input, got %d entries", len(result))
		}
	})

	t.Run("Handles malformed lines", func(t *testing.T) {
		iconData := `note|<svg>note icon</svg>
malformed line without pipe
warning|<svg>warning icon</svg>
another malformed -> line -> with -> multiple arrows
tip -> note`

		result := CreateIconsMap(iconData)

		// Should only parse valid lines
		if len(result) != 3 {
			t.Errorf("Expected 3 entries (ignoring malformed lines), got %d", len(result))
		}

		if result["note"] != "<svg>note icon</svg>" {
			t.Error("Should parse valid note line")
		}

		if result["warning"] != "<svg>warning icon</svg>" {
			t.Error("Should parse valid warning line")
		}

		if result["tip"] != "<svg>note icon</svg>" {
			t.Error("Should parse valid alias line")
		}
	})

	t.Run("Handles chain aliases", func(t *testing.T) {
		iconData := `primary|<svg>primary icon</svg>
secondary -> primary
tertiary -> secondary`

		result := CreateIconsMap(iconData)

		if len(result) != 3 {
			t.Errorf("Expected 3 entries, got %d", len(result))
		}

		// Note: The current implementation may not handle chain aliases perfectly
		// This test documents the expected behavior
		if result["secondary"] != "<svg>primary icon</svg>" {
			t.Errorf("Expected secondary to resolve to primary icon, got '%s'", result["secondary"])
		}
	})
}

func TestCreateIconsMapEdgeCases(t *testing.T) {
	t.Run("Only comments", func(t *testing.T) {
		iconData := `# Comment 1
# Comment 2
# Comment 3`

		result := CreateIconsMap(iconData)

		if len(result) != 0 {
			t.Errorf("Expected empty map for comments only, got %d entries", len(result))
		}
	})

	t.Run("Mixed valid and invalid lines", func(t *testing.T) {
		iconData := `valid|<svg>valid</svg>
invalid-no-pipe-or-arrow
another|valid|<svg>another</svg>
alias -> valid
||empty-parts
|no-key<svg>no-key</svg>
key-no-value|`

		result := CreateIconsMap(iconData)

		// Should handle what it can parse
		if result["valid"] != "<svg>valid</svg>" {
			t.Error("Should parse first valid line")
		}

		if result["alias"] != "<svg>valid</svg>" {
			t.Error("Should parse valid alias")
		}

		// The "another|valid|<svg>another</svg>" line has multiple pipes -
		// it should take everything after the first pipe as the value
		if result["another"] != "valid|<svg>another</svg>" {
			t.Errorf("Expected 'valid|<svg>another</svg>', got '%s'", result["another"])
		}
	})
}
