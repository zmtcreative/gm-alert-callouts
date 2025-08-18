package alertcallouts

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
)

func TestAlertCalloutsIntegration(t *testing.T) {
	t.Run("Integration with Goldmark", func(t *testing.T) {
		ext := NewAlertCallouts(
			WithIcon("note", "<svg>note-icon</svg>"),
			WithFolding(true),
		)

		md := goldmark.New(goldmark.WithExtensions(ext))

		input := `> [!note]
> This is a test note`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		result := output.String()

		// Check that the extension is working
		if !strings.Contains(result, "callout") {
			t.Error("Expected output to contain callout class")
		}

		if !strings.Contains(result, "<svg>note-icon</svg>") {
			t.Error("Expected output to contain the note icon")
		}

		if !strings.Contains(result, "This is a test note") {
			t.Error("Expected output to contain the note content")
		}
	})

	t.Run("Backwards compatibility with existing variable", func(t *testing.T) {
		// Ensure the old AlertCallouts variable still works
		md := goldmark.New(goldmark.WithExtensions(AlertCallouts))

		input := `> [!note]
> Test content`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		// Should still produce alert output (even without icons)
		result := output.String()
		if !strings.Contains(result, "callout") {
			t.Error("Expected backwards compatibility to work")
		}
	})

	t.Run("Folding functionality with new initializer", func(t *testing.T) {
		ext := NewAlertCallouts(
			WithIcon("tip", "<svg>tip-icon</svg>"),
			WithFolding(true), // Folding enabled
		)

		md := goldmark.New(goldmark.WithExtensions(ext))

		input := `> [!tip]-
> This should be a closed foldable callout`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		result := output.String()

		// Check for folding elements
		if !strings.Contains(result, "<details") {
			t.Error("Expected output to contain details element for folding")
		}

		if !strings.Contains(result, "<summary") {
			t.Error("Expected output to contain summary element for folding")
		}

		if !strings.Contains(result, "callout-foldable") {
			t.Error("Expected output to contain foldable class")
		}
	})

	t.Run("No folding when disabled", func(t *testing.T) {
		ext := NewAlertCallouts(
			WithIcon("warning", "<svg>warning-icon</svg>"),
			WithFolding(false), // Folding disabled
		)

		md := goldmark.New(goldmark.WithExtensions(ext))

		input := `> [!warning]-
> This should not be foldable`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		result := output.String()

		// Check that folding elements are not present
		if strings.Contains(result, "<details") {
			t.Error("Expected output to NOT contain details element when folding disabled")
		}

		if strings.Contains(result, "<summary") {
			t.Error("Expected output to NOT contain summary element when folding disabled")
		}
	})

	t.Run("Multiple alerts in sequence", func(t *testing.T) {
		ext := NewAlertCallouts(
			WithIcons(map[string]string{
				"note":    "<svg>note</svg>",
				"warning": "<svg>warning</svg>",
			}),
			WithFolding(true),
		)

		md := goldmark.New(goldmark.WithExtensions(ext))

		input := `> [!note]
> First alert

> [!warning]
> Second alert`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		result := output.String()

		// Count occurrences of callout divs
		noteCount := strings.Count(result, "callout-note")
		warningCount := strings.Count(result, "callout-warning")

		if noteCount != 1 {
			t.Errorf("Expected 1 note callout, got %d", noteCount)
		}

		if warningCount != 1 {
			t.Errorf("Expected 1 warning callout, got %d", warningCount)
		}
	})

	t.Run("Mixed content with regular blockquotes", func(t *testing.T) {
		ext := NewAlertCallouts(
			WithIcon("info", "<svg>info</svg>"),
		)

		md := goldmark.New(goldmark.WithExtensions(ext))

		input := `> Regular blockquote

> [!info]
> Alert callout

> Another regular blockquote`

		var output strings.Builder
		err := md.Convert([]byte(input), &output)
		if err != nil {
			t.Fatalf("Failed to convert markdown: %v", err)
		}

		result := output.String()

		// Should have one callout and two blockquotes
		calloutCount := strings.Count(result, "callout-info")
		blockquoteCount := strings.Count(result, "<blockquote>")

		if calloutCount != 1 {
			t.Errorf("Expected 1 info callout, got %d", calloutCount)
		}

		if blockquoteCount != 2 {
			t.Errorf("Expected 2 regular blockquotes, got %d", blockquoteCount)
		}
	})
}
