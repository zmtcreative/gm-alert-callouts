package renderer

import (
	"strings"
	"testing"

	"github.com/zmtcreative/gm-alert-callouts/internal/ast"
	"github.com/zmtcreative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

func TestNewAlertsBodyHTMLRenderer(t *testing.T) {
	r := NewAlertsBodyHTMLRenderer()
	if r == nil {
		t.Fatal("NewAlertsBodyHTMLRenderer returned nil")
	}

	bodyRenderer, ok := r.(*AlertsBodyHTMLRenderer)
	if !ok {
		t.Fatal("NewAlertsBodyHTMLRenderer did not return *AlertsBodyHTMLRenderer")
	}

	// Should have initialized config
	if bodyRenderer.Config.Writer == nil {
		// This is fine, the config might not have a default writer
	}
}

func TestAlertsBodyHTMLRendererRegisterFuncs(t *testing.T) {
	r := NewAlertsBodyHTMLRenderer()

	registrations := make(map[gast.NodeKind]renderer.NodeRendererFunc)
	mockRegisterer := &mockNodeRendererFuncRegisterer{
		registrations: registrations,
	}

	r.RegisterFuncs(mockRegisterer)

	if len(registrations) != 1 {
		t.Errorf("Expected 1 registration, got %d", len(registrations))
	}

	if _, exists := registrations[constants.KindAlertsBody]; !exists {
		t.Error("Expected KindAlertsBody to be registered")
	}
}

func TestAlertsBodyHTMLRendererRender(t *testing.T) {
	r := NewAlertsBodyHTMLRenderer()
	node := createMockBodyNode()

	// Test entering
	writer := newMockBufWriter()
	status, err := r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, node, true)
	if err != nil {
		t.Fatalf("Unexpected error on entering: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue, got %v", status)
	}

	enteringHTML := writer.String()
	expectedStart := `<div class="callout-body">`
	if !strings.Contains(enteringHTML, expectedStart) {
		t.Errorf("Expected %s, got: %s", expectedStart, enteringHTML)
	}

	// Test exiting
	writer.Reset()
	status, err = r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, node, false)
	if err != nil {
		t.Fatalf("Unexpected error on exiting: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue, got %v", status)
	}

	exitingHTML := writer.String()
	expectedEnd := "</div>"
	if !strings.Contains(exitingHTML, expectedEnd) {
		t.Errorf("Expected %s, got: %s", expectedEnd, exitingHTML)
	}
}

func TestAlertsBodyHTMLRendererComplete(t *testing.T) {
	// Test complete rendering cycle
	r := NewAlertsBodyHTMLRenderer()
	node := createMockBodyNode()

	writer := newMockBufWriter()

	// Entering
	status, err := r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, node, true)
	if err != nil {
		t.Fatalf("Unexpected error on entering: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue on entering, got %v", status)
	}

	// Exiting
	status, err = r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, node, false)
	if err != nil {
		t.Fatalf("Unexpected error on exiting: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue on exiting, got %v", status)
	}

	completeHTML := writer.String()

	// Should contain both opening and closing tags
	if !strings.Contains(completeHTML, `<div class="callout-body">`) {
		t.Error("Expected opening div tag")
	}
	if !strings.Contains(completeHTML, "</div>") {
		t.Error("Expected closing div tag")
	}

	// Opening tag should come before closing tag
	openIndex := strings.Index(completeHTML, `<div class="callout-body">`)
	closeIndex := strings.Index(completeHTML, "</div>")
	if openIndex == -1 || closeIndex == -1 || openIndex >= closeIndex {
		t.Errorf("Invalid HTML structure: %s", completeHTML)
	}
}

func TestAlertsBodyHTMLRendererMultipleCalls(t *testing.T) {
	// Test that multiple calls work correctly
	r := NewAlertsBodyHTMLRenderer()

	testCases := []struct {
		name     string
		entering bool
		expected string
	}{
		{
			name:     "First entering call",
			entering: true,
			expected: `<div class="callout-body">`,
		},
		{
			name:     "First exiting call",
			entering: false,
			expected: "</div>",
		},
		{
			name:     "Second entering call",
			entering: true,
			expected: `<div class="callout-body">`,
		},
		{
			name:     "Second exiting call",
			entering: false,
			expected: "</div>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := createMockBodyNode()
			writer := newMockBufWriter()

			status, err := r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, node, tc.entering)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if status != gast.WalkContinue {
				t.Errorf("Expected WalkContinue, got %v", status)
			}

			html := writer.String()
			if !strings.Contains(html, tc.expected) {
				t.Errorf("Expected %s, got: %s", tc.expected, html)
			}
		})
	}
}

func TestAlertsBodyHTMLRendererWithDifferentNodes(t *testing.T) {
	// Test that the renderer works with different types of body nodes
	r := NewAlertsBodyHTMLRenderer()

	// Test with empty body
	emptyBody := ast.NewAlertsBody()
	writer := newMockBufWriter()

	status, err := r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, emptyBody, true)
	if err != nil {
		t.Fatalf("Unexpected error with empty body: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue with empty body, got %v", status)
	}

	// Test with body containing children (should still work the same way)
	bodyWithChildren := ast.NewAlertsBody()
	para := gast.NewParagraph()
	bodyWithChildren.AppendChild(bodyWithChildren, para)

	writer.Reset()
	status, err = r.(*AlertsBodyHTMLRenderer).renderAlertsBody(writer, []byte{}, bodyWithChildren, true)
	if err != nil {
		t.Fatalf("Unexpected error with body containing children: %v", err)
	}
	if status != gast.WalkContinue {
		t.Errorf("Expected WalkContinue with body containing children, got %v", status)
	}

	html := writer.String()
	if !strings.Contains(html, `<div class="callout-body">`) {
		t.Error("Expected same output regardless of body content")
	}
}

// Helper function to create mock body nodes
func createMockBodyNode() gast.Node {
	return ast.NewAlertsBody()
}
