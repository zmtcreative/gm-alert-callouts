package renderer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/ast"
	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

// mockBufWriter implements util.BufWriter for testing
type mockBufWriter struct {
	*bytes.Buffer
}

func (m *mockBufWriter) Buffered() int {
	return m.Buffer.Len()
}

func (m *mockBufWriter) Available() int {
	return 1024 // arbitrary value for testing
}

func (m *mockBufWriter) Flush() error {
	return nil
}

func newMockBufWriter() *mockBufWriter {
	return &mockBufWriter{Buffer: &bytes.Buffer{}}
}

func TestNewAlertsHTMLRenderer(t *testing.T) {
	r := NewAlertsHTMLRenderer(FoldingEnabled(true), constants.ICONS_GFM)
	if r == nil {
		t.Fatal("NewAlertsHTMLRenderer returned nil")
	}

	alertsRenderer, ok := r.(*AlertsHTMLRenderer)
	if !ok {
		t.Fatal("NewAlertsHTMLRenderer did not return *AlertsHTMLRenderer")
	}

	if alertsRenderer.FoldingEnabled != FoldingEnabled(true) {
		t.Error("FoldingEnabled not set correctly")
	}

	if alertsRenderer.DefaultIcons != constants.ICONS_GFM {
		t.Error("DefaultIcons not set correctly")
	}
}

func TestAlertsHTMLRendererRegisterFuncs(t *testing.T) {
	r := NewAlertsHTMLRenderer(FoldingEnabled(false), constants.ICONS_NONE)

	// Create a mock registerer to capture what gets registered
	registrations := make(map[gast.NodeKind]renderer.NodeRendererFunc)
	mockRegisterer := &mockNodeRendererFuncRegisterer{
		registrations: registrations,
	}

	r.RegisterFuncs(mockRegisterer)

	if len(registrations) != 1 {
		t.Errorf("Expected 1 registration, got %d", len(registrations))
	}

	if _, exists := registrations[constants.KindAlerts]; !exists {
		t.Error("Expected KindAlerts to be registered")
	}
}

func TestAlertsHTMLRendererBasicAlert(t *testing.T) {
	testCases := []struct {
		name        string
		folding     FoldingEnabled
		defaultIcons int
		kind        string
		expectedDiv string
		expectedClass string
	}{
		{
			name:        "Basic note alert without folding",
			folding:     FoldingEnabled(false),
			defaultIcons: constants.ICONS_NONE,
			kind:        "note",
			expectedDiv: "div",
			expectedClass: `class="callout callout-note"`,
		},
		{
			name:        "Warning alert with GFM icons",
			folding:     FoldingEnabled(false),
			defaultIcons: constants.ICONS_GFM,
			kind:        "warning",
			expectedDiv: "div",
			expectedClass: `class="callout callout-warning iconset-gfm"`,
		},
		{
			name:        "Info alert with GFM Plus icons",
			folding:     FoldingEnabled(false),
			defaultIcons: constants.ICONS_GFM_PLUS,
			kind:        "info",
			expectedDiv: "div",
			expectedClass: `class="callout callout-info iconset-gfmplus"`,
		},
		{
			name:        "Tip alert with Obsidian icons",
			folding:     FoldingEnabled(false),
			defaultIcons: constants.ICONS_OBSIDIAN,
			kind:        "tip",
			expectedDiv: "div",
			expectedClass: `class="callout callout-tip iconset-obsidian"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewAlertsHTMLRenderer(tc.folding, tc.defaultIcons)

			// Create mock alert node
			node := createMockAlertNode(tc.kind, false, false)

			// Test entering
			writer := newMockBufWriter()

			status, err := r.(*AlertsHTMLRenderer).renderAlerts(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error on entering: %v", err)
			}
			if status != gast.WalkContinue {
				t.Errorf("Expected WalkContinue, got %v", status)
			}

			enteringHTML := writer.String()
			if !strings.Contains(enteringHTML, "<"+tc.expectedDiv) {
				t.Errorf("Expected <%s> element, got: %s", tc.expectedDiv, enteringHTML)
			}
			if !strings.Contains(enteringHTML, tc.expectedClass) {
				t.Errorf("Expected class %s, got: %s", tc.expectedClass, enteringHTML)
			}
			if !strings.Contains(enteringHTML, `data-callout="`+tc.kind+`"`) {
				t.Errorf("Expected data-callout attribute, got: %s", enteringHTML)
			}

			// Test exiting
			writer.Reset()
			status, err = r.(*AlertsHTMLRenderer).renderAlerts(writer, []byte{}, node, false)
			if err != nil {
				t.Fatalf("Unexpected error on exiting: %v", err)
			}
			if status != gast.WalkContinue {
				t.Errorf("Expected WalkContinue, got %v", status)
			}

			exitingHTML := writer.String()
			expectedEndTag := "</" + tc.expectedDiv + ">"
			if !strings.Contains(exitingHTML, expectedEndTag) {
				t.Errorf("Expected %s, got: %s", expectedEndTag, exitingHTML)
			}
		})
	}
}

func TestAlertsHTMLRendererFoldableAlerts(t *testing.T) {
	testCases := []struct {
		name        string
		folding     FoldingEnabled
		closed      bool
		shouldFold  bool
		expectedTag string
		expectedOpen string
	}{
		{
			name:        "Foldable closed alert",
			folding:     FoldingEnabled(true),
			closed:      true,
			shouldFold:  true,
			expectedTag: "details",
			expectedOpen: "",
		},
		{
			name:        "Foldable open alert",
			folding:     FoldingEnabled(true),
			closed:      false,
			shouldFold:  true,
			expectedTag: "details",
			expectedOpen: " open",
		},
		{
			name:        "Non-foldable alert with folding enabled",
			folding:     FoldingEnabled(true),
			closed:      false,
			shouldFold:  false,
			expectedTag: "div",
			expectedOpen: "",
		},
		{
			name:        "Alert with folding disabled",
			folding:     FoldingEnabled(false),
			closed:      true,
			shouldFold:  true,
			expectedTag: "div",
			expectedOpen: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewAlertsHTMLRenderer(tc.folding, constants.ICONS_NONE)

			node := createMockAlertNode("note", tc.closed, tc.shouldFold)

			writer := newMockBufWriter()

			// Test entering
			_, err := r.(*AlertsHTMLRenderer).renderAlerts(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			enteringHTML := writer.String()
			if !strings.Contains(enteringHTML, "<"+tc.expectedTag) {
				t.Errorf("Expected <%s> element, got: %s", tc.expectedTag, enteringHTML)
			}

			if tc.shouldFold && bool(tc.folding) {
				if !strings.Contains(enteringHTML, "callout-foldable") {
					t.Error("Expected callout-foldable class")
				}

				if tc.expectedOpen != "" {
					if !strings.Contains(enteringHTML, tc.expectedOpen) {
						t.Errorf("Expected %s attribute, got: %s", tc.expectedOpen, enteringHTML)
					}
				} else {
					if strings.Contains(enteringHTML, " open") {
						t.Error("Did not expect open attribute for closed alert")
					}
				}
			}
		})
	}
}

func TestAlertsHTMLRendererNoIconKind(t *testing.T) {
	r := NewAlertsHTMLRenderer(FoldingEnabled(false), constants.ICONS_NONE)

	testCases := []struct {
		name     string
		kind     string
		title    string
		expected string
	}{
		{
			name:     "No icon with title",
			kind:     "noicon",
			title:    "Custom Alert",
			expected: "custom-alert", // should use cleaned title
		},
		{
			name:     "No-icon with title",
			kind:     "no-icon",
			title:    "Another Custom",
			expected: "another-custom",
		},
		{
			name:     "Nil with title",
			kind:     "nil",
			title:    "Nil Alert",
			expected: "nil-alert",
		},
		{
			name:     "Null with title",
			kind:     "null",
			title:    "Null Alert",
			expected: "null-alert",
		},
		{
			name:     "No icon without title",
			kind:     "noicon",
			title:    "",
			expected: "noicon", // should fall back to kind
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := ast.NewAlerts()
			node.SetAttributeString("kind", []byte(tc.kind))
			if tc.title != "" {
				node.SetAttributeString("title", []byte(tc.title))
			}
			node.SetAttributeString("closed", false)
			node.SetAttributeString("shouldfold", false)

			writer := newMockBufWriter()

			_, err := r.(*AlertsHTMLRenderer).renderAlerts(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			html := writer.String()
			if !strings.Contains(html, "callout-"+tc.expected) {
				t.Errorf("Expected callout-%s class, got: %s", tc.expected, html)
			}
			if !strings.Contains(html, `data-callout="`+tc.expected+`"`) {
				t.Errorf("Expected data-callout=%s, got: %s", tc.expected, html)
			}
		})
	}
}

func TestAlertsHTMLRendererTitleCleaning(t *testing.T) {
	r := NewAlertsHTMLRenderer(FoldingEnabled(false), constants.ICONS_NONE)

	testCases := []struct {
		name     string
		title    string
		expected string
	}{
		{
			name:     "Title with HTML tags",
			title:    "Alert <strong>with</strong> HTML",
			expected: "alert-with-html",
		},
		{
			name:     "Title with markdown",
			title:    "Alert **with** *markdown*",
			expected: "alert-with-markdown",
		},
		{
			name:     "Title with links",
			title:    "Alert [with](link) content",
			expected: "alert-content", // Links are removed, leaving "Alert  content" -> "alert-content"
		},
		{
			name:     "Title with code",
			title:    "Alert `with code` content",
			expected: "alert-content", // Code is removed, leaving "Alert  content" -> "alert-content"
		},
		{
			name:     "Title with multiple spaces",
			title:    "Alert   with    spaces",
			expected: "alert-with-spaces",
		},
		{
			name:     "Complex title",
			title:    "**Complex** [Title](link) with `code` and <em>HTML</em>",
			expected: "complex-with-and-html", // All markdown/HTML removed, multiple spaces collapsed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := ast.NewAlerts()
			node.SetAttributeString("kind", []byte("noicon"))
			node.SetAttributeString("title", []byte(tc.title))
			node.SetAttributeString("closed", false)
			node.SetAttributeString("shouldfold", false)

			writer := newMockBufWriter()

			_, err := r.(*AlertsHTMLRenderer).renderAlerts(writer, []byte{}, node, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			html := writer.String()
			if !strings.Contains(html, "callout-"+tc.expected) {
				t.Errorf("Expected callout-%s class, got: %s", tc.expected, html)
			}
		})
	}
}

// Helper functions

func createMockAlertNode(kind string, closed bool, shouldFold bool) gast.Node {
	node := ast.NewAlerts()
	node.SetAttributeString("kind", []byte(kind))
	node.SetAttributeString("closed", closed)
	node.SetAttributeString("shouldfold", shouldFold)
	return node
}

type mockNodeRendererFuncRegisterer struct {
	registrations map[gast.NodeKind]renderer.NodeRendererFunc
}

func (m *mockNodeRendererFuncRegisterer) Register(kind gast.NodeKind, fn renderer.NodeRendererFunc) {
	m.registrations[kind] = fn
}
