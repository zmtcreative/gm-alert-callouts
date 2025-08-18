package renderer

import (
	"strings"
	"testing"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/ast"
	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

func TestNewAlertsHeaderHTMLRenderer(t *testing.T) {
	r := NewAlertsHeaderHTMLRenderer(FoldingEnabled(true), constants.ICONS_GFM)
	if r == nil {
		t.Fatal("NewAlertsHeaderHTMLRenderer returned nil")
	}

	headerRenderer, ok := r.(*AlertsHeaderHTMLRenderer)
	if !ok {
		t.Fatal("NewAlertsHeaderHTMLRenderer did not return *AlertsHeaderHTMLRenderer")
	}

	if headerRenderer.FoldingEnabled != FoldingEnabled(true) {
		t.Error("FoldingEnabled not set correctly")
	}

	if headerRenderer.DefaultIcons != constants.ICONS_GFM {
		t.Error("DefaultIcons not set correctly")
	}
}

func TestNewAlertsHeaderHTMLRendererWithIcons(t *testing.T) {
	icons := Icons{
		"note":    "<svg>note-icon</svg>",
		"warning": "<svg>warning-icon</svg>",
	}

	r := NewAlertsHeaderHTMLRendererWithIcons(icons, FoldingEnabled(false), constants.ICONS_NONE)
	if r == nil {
		t.Fatal("NewAlertsHeaderHTMLRendererWithIcons returned nil")
	}

	headerRenderer, ok := r.(*AlertsHeaderHTMLRenderer)
	if !ok {
		t.Fatal("NewAlertsHeaderHTMLRendererWithIcons did not return *AlertsHeaderHTMLRenderer")
	}

	if headerRenderer.FoldingEnabled != FoldingEnabled(false) {
		t.Error("FoldingEnabled not set correctly")
	}

	if len(headerRenderer.Icons) != 2 {
		t.Errorf("Expected 2 icons, got %d", len(headerRenderer.Icons))
	}

	if headerRenderer.Icons["note"] != "<svg>note-icon</svg>" {
		t.Error("Note icon not set correctly")
	}
}

func TestAlertsHeaderHTMLRendererRegisterFuncs(t *testing.T) {
	r := NewAlertsHeaderHTMLRenderer(FoldingEnabled(false), constants.ICONS_NONE)

	registrations := make(map[gast.NodeKind]renderer.NodeRendererFunc)
	mockRegisterer := &mockNodeRendererFuncRegisterer{
		registrations: registrations,
	}

	r.RegisterFuncs(mockRegisterer)

	if len(registrations) != 1 {
		t.Errorf("Expected 1 registration, got %d", len(registrations))
	}

	if _, exists := registrations[constants.KindAlertsHeader]; !exists {
		t.Error("Expected KindAlertsHeader to be registered")
	}
}

func TestAlertsHeaderHTMLRendererBasicHeader(t *testing.T) {
	testCases := []struct {
		name          string
		folding       FoldingEnabled
		shouldFold    bool
		expectedStart string
		expectedEnd   string
	}{
		{
			name:          "Non-foldable header",
			folding:       FoldingEnabled(false),
			shouldFold:    false,
			expectedStart: `<div class="callout-title">`,
			expectedEnd:   "</div>",
		},
		{
			name:          "Foldable header",
			folding:       FoldingEnabled(true),
			shouldFold:    true,
			expectedStart: `<summary class="callout-title">`,
			expectedEnd:   "</summary>",
		},
		{
			name:          "Non-foldable with folding disabled",
			folding:       FoldingEnabled(false),
			shouldFold:    true,
			expectedStart: `<div class="callout-title">`,
			expectedEnd:   "</div>",
		},
		{
			name:          "Non-foldable alert with folding enabled",
			folding:       FoldingEnabled(true),
			shouldFold:    false,
			expectedStart: `<div class="callout-title">`,
			expectedEnd:   "</div>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewAlertsHeaderHTMLRenderer(tc.folding, constants.ICONS_NONE)

			node := createMockHeaderNode("note", tc.shouldFold, "")

			// Test entering
			writer := newMockBufWriter()
			status, err := r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error on entering: %v", err)
			}
			if status != gast.WalkContinue {
				t.Errorf("Expected WalkContinue, got %v", status)
			}

			enteringHTML := writer.String()
			if !strings.Contains(enteringHTML, tc.expectedStart) {
				t.Errorf("Expected %s, got: %s", tc.expectedStart, enteringHTML)
			}

			// Test exiting
			writer.Reset()
			status, err = r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, false)
			if err != nil {
				t.Fatalf("Unexpected error on exiting: %v", err)
			}
			if status != gast.WalkContinue {
				t.Errorf("Expected WalkContinue, got %v", status)
			}

			exitingHTML := writer.String()
			if !strings.Contains(exitingHTML, tc.expectedEnd) {
				t.Errorf("Expected %s, got: %s", tc.expectedEnd, exitingHTML)
			}
		})
	}
}

func TestAlertsHeaderHTMLRendererWithIcons(t *testing.T) {
	icons := Icons{
		"note":    "<svg class='note-icon'>Note</svg>",
		"warning": "<svg class='warning-icon'>Warning</svg>",
		"info":    "<svg class='info-icon'>Info</svg>",
		"default": "<svg class='default-icon'>Default</svg>",
	}

	testCases := []struct {
		name         string
		kind         string
		expectedIcon string
	}{
		{
			name:         "Note with specific icon",
			kind:         "note",
			expectedIcon: "<svg class='note-icon'>Note</svg>",
		},
		{
			name:         "Warning with specific icon",
			kind:         "warning",
			expectedIcon: "<svg class='warning-icon'>Warning</svg>",
		},
		{
			name:         "Unknown kind with default fallback",
			kind:         "unknown",
			expectedIcon: "<svg class='note-icon'>Note</svg>", // Falls back to note first
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewAlertsHeaderHTMLRendererWithIcons(icons, FoldingEnabled(false), constants.ICONS_NONE)

			node := createMockHeaderNode(tc.kind, false, "")

			writer := newMockBufWriter()
			_, err := r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			html := writer.String()
			if !strings.Contains(html, tc.expectedIcon) {
				t.Errorf("Expected icon %s, got: %s", tc.expectedIcon, html)
			}
		})
	}
}

func TestAlertsHeaderHTMLRendererNoIconKinds(t *testing.T) {
	icons := Icons{
		"note":    "<svg>note-icon</svg>",
		"default": "<svg>default-icon</svg>",
	}

	testCases := []string{"noicon", "no_icon", "none", "nil", "null"}

	for _, kind := range testCases {
		t.Run("No icon for "+kind, func(t *testing.T) {
			r := NewAlertsHeaderHTMLRendererWithIcons(icons, FoldingEnabled(false), constants.ICONS_NONE)

			node := createMockHeaderNode(kind, false, "")

			writer := newMockBufWriter()
			_, err := r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			html := writer.String()
			// Should not contain any icons for noicon kinds
			if strings.Contains(html, "<svg>") {
				t.Errorf("Expected no icon for kind %s, but found icon in: %s", kind, html)
			}
		})
	}
}

func TestAlertsHeaderHTMLRendererTitleCasing(t *testing.T) {
	r := NewAlertsHeaderHTMLRendererWithIcons(Icons{}, FoldingEnabled(false), constants.ICONS_NONE)

	testCases := []struct {
		name         string
		kind         string
		hasTitle     bool
		expectedText string
	}{
		{
			name:         "Note kind without title",
			kind:         "note",
			hasTitle:     false,
			expectedText: "Note", // should be title cased
		},
		{
			name:         "Warning kind without title",
			kind:         "warning",
			hasTitle:     false,
			expectedText: "Warning",
		},
		{
			name:         "Multi-word kind",
			kind:         "important-info",
			hasTitle:     false,
			expectedText: "Important-Info", // should be title cased
		},
		{
			name:         "With custom title",
			kind:         "note",
			hasTitle:     true,
			expectedText: "", // should not add kind text when title exists
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var title string
			if tc.hasTitle {
				title = "Custom Title"
			}

			node := createMockHeaderNode(tc.kind, false, title)

			writer := newMockBufWriter()
			_, err := r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			html := writer.String()
			if tc.expectedText != "" {
				if !strings.Contains(html, tc.expectedText) {
					t.Errorf("Expected text %s, got: %s", tc.expectedText, html)
				}
			}

			// Should always contain the title text element
			if !strings.Contains(html, `<p class="callout-title-text">`) {
				t.Error("Expected callout-title-text element")
			}
		})
	}
}

func TestAlertsHeaderHTMLRendererIconFallback(t *testing.T) {
	icons := Icons{
		"info": "<svg>info-icon</svg>", // has info but not note
	}

	r := NewAlertsHeaderHTMLRendererWithIcons(icons, FoldingEnabled(false), constants.ICONS_NONE)

	// Test fallback to 'note' when specific kind not found
	node := createMockHeaderNode("unknown", false, "")

	writer := newMockBufWriter()
	_, err := r.(*AlertsHeaderHTMLRenderer).renderAlertsHeader(writer, []byte{}, node, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	html := writer.String()
	// Since 'unknown' kind is not in icons, and 'note' is not available either,
	// and 'info' is available, it should fall back to 'info'
	if !strings.Contains(html, "<svg>info-icon</svg>") {
		t.Errorf("Expected fallback to info icon, got: %s", html)
	}
}

// Helper function to create mock header nodes
func createMockHeaderNode(kind string, shouldFold bool, title string) gast.Node {
	node := ast.NewAlertsHeader()
	node.SetAttributeString("kind", kind)
	node.SetAttributeString("shouldfold", shouldFold)

	if title != "" {
		node.SetAttributeString("title", title)
		// Add a text block child to simulate title content
		textBlock := gast.NewTextBlock()
		node.AppendChild(node, textBlock)
	}

	return node
}
