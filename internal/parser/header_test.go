package parser

import (
	"testing"

	"github.com/ZMT-Creative/gm-alert-callouts/internal/ast"
	"github.com/ZMT-Creative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestNewAlertsHeaderParser(t *testing.T) {
	p := NewAlertsHeaderParser()
	if p == nil {
		t.Fatal("NewAlertsHeaderParser() returned nil")
	}

	if p != defaultAlertsHeaderParser {
		t.Error("NewAlertsHeaderParser() should return defaultAlertsHeaderParser")
	}
}

func TestAlertsHeaderParserTrigger(t *testing.T) {
	p := &alertHeaderParser{}
	trigger := p.Trigger()
	expected := []byte{']'}

	if len(trigger) != 1 || trigger[0] != expected[0] {
		t.Errorf("Expected trigger %v, got %v", expected, trigger)
	}
}

func TestAlertsHeaderParserOpen(t *testing.T) {
	p := &alertHeaderParser{}
	pc := parser.NewContext()

	testCases := []struct {
		name      string
		parent    gast.Node
		input     string
		expectNil bool
		kind      string
		hasTitle  bool
	}{
		{
			name:      "Valid header with title",
			parent:    createMockAlertsParent("note", true),
			input:     "] Custom Title\n",
			expectNil: false,
			kind:      "note",
			hasTitle:  true,
		},
		{
			name:      "Valid header with folding marker",
			parent:    createMockAlertsParent("warning", true),
			input:     "]- Warning Title\n",
			expectNil: false,
			kind:      "warning",
			hasTitle:  true,
		},
		{
			name:      "Valid header with plus marker",
			parent:    createMockAlertsParent("info", true),
			input:     "]+ Info Title\n",
			expectNil: false,
			kind:      "info",
			hasTitle:  true,
		},
		{
			name:      "Valid header no title",
			parent:    createMockAlertsParent("tip", false),
			input:     "]\n",
			expectNil: false,
			kind:      "tip",
			hasTitle:  false,
		},
		{
			name:      "Invalid parent - not alerts",
			parent:    gast.NewDocument(),
			input:     "] Title\n",
			expectNil: true,
		},
		{
			name:      "Invalid parent - already has children",
			parent:    createMockAlertsParentWithChild("note"),
			input:     "] Title\n",
			expectNil: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := text.NewReader([]byte(tc.input))

			node, state := p.Open(tc.parent, reader, pc)

			if tc.expectNil {
				if node != nil {
					t.Errorf("Expected nil node, got %v", node)
				}
				if state != parser.NoChildren {
					t.Errorf("Expected NoChildren state, got %v", state)
				}
			} else {
				if node == nil {
					t.Fatalf("Expected node to be created, got nil")
				}

				if node.Kind() != constants.KindAlertsHeader {
					t.Errorf("Expected KindAlertsHeader, got %v", node.Kind())
				}

				if state != parser.NoChildren {
					t.Errorf("Expected NoChildren state, got %v", state)
				}

				// Check kind attribute
				if kind, ok := node.AttributeString("kind"); ok {
					if kind != tc.kind {
						t.Errorf("Expected kind %s, got %s", tc.kind, kind)
					}
				} else {
					t.Error("Expected kind attribute")
				}

				// Check shouldfold attribute if parent has it
				if shouldfold, ok := tc.parent.AttributeString("shouldfold"); ok {
					if headerShouldFold, exists := node.AttributeString("shouldfold"); !exists {
						t.Error("Expected shouldfold attribute to be copied from parent")
					} else if headerShouldFold != shouldfold {
						t.Errorf("Expected shouldfold to match parent: %v, got %v", shouldfold, headerShouldFold)
					}
				}

				// Check for title content
				if tc.hasTitle {
					if node.ChildCount() == 0 {
						t.Error("Expected header to have child content when title is present")
					}
				}
			}
		})
	}
}

func TestAlertsHeaderParserContinue(t *testing.T) {
	p := &alertHeaderParser{}
	pc := parser.NewContext()
	node := ast.NewAlertsHeader()
	reader := text.NewReader([]byte("any input"))

	state := p.Continue(node, reader, pc)
	if state != parser.Close {
		t.Errorf("Expected Close state, got %v", state)
	}
}

func TestAlertsHeaderParserClose(t *testing.T) {
	p := &alertHeaderParser{}
	pc := parser.NewContext()
	node := ast.NewAlertsHeader()
	reader := text.NewReader([]byte(""))

	// Close should not panic and should complete without error
	p.Close(node, reader, pc)
}

func TestAlertsHeaderParserCanInterruptParagraph(t *testing.T) {
	p := &alertHeaderParser{}
	if p.CanInterruptParagraph() {
		t.Error("Expected CanInterruptParagraph to return false")
	}
}

func TestAlertsHeaderParserCanAcceptIndentedLine(t *testing.T) {
	p := &alertHeaderParser{}
	if !p.CanAcceptIndentedLine() {
		t.Error("Expected CanAcceptIndentedLine to return true")
	}
}

func TestAlertsHeaderParserTitleHandling(t *testing.T) {
	p := &alertHeaderParser{}
	pc := parser.NewContext()

	testCases := []struct {
		name       string
		input      string
		expectText bool
	}{
		{
			name:       "Title with content",
			input:      "] This is a title\n",
			expectText: true,
		},
		{
			name:       "Title with leading spaces",
			input:      "]   Spaced title\n",
			expectText: true,
		},
		{
			name:       "Empty title",
			input:      "]\n",
			expectText: false,
		},
		{
			name:       "Title with marker",
			input:      "]- Folded title\n",
			expectText: true,
		},
		{
			name:       "Title with plus",
			input:      "]+ Open title\n",
			expectText: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parent := createMockAlertsParent("test", true)
			reader := text.NewReader([]byte(tc.input))

			node, _ := p.Open(parent, reader, pc)
			if node == nil {
				t.Fatal("Expected node to be created")
			}

			hasChildren := node.ChildCount() > 0
			if hasChildren != tc.expectText {
				t.Errorf("Expected hasChildren=%v, got %v", tc.expectText, hasChildren)
			}

			if tc.expectText && hasChildren {
				// Should have a TextBlock child
				child := node.FirstChild()
				if child.Kind() != gast.KindTextBlock {
					t.Errorf("Expected TextBlock child, got %v", child.Kind())
				}
			}
		})
	}
}

// Helper function to create a mock alerts parent node
func createMockAlertsParent(kind string, shouldFold bool) gast.Node {
	parent := ast.NewAlerts()
	parent.SetAttributeString("kind", []byte(kind))
	parent.SetAttributeString("shouldfold", shouldFold)
	return parent
}

// Helper function to create a mock alerts parent node with existing child
func createMockAlertsParentWithChild(kind string) gast.Node {
	parent := ast.NewAlerts()
	parent.SetAttributeString("kind", []byte(kind))
	// Add a child to make ChildCount() > 0
	child := gast.NewParagraph()
	parent.AppendChild(parent, child)
	return parent
}

func TestAlertsHeaderParserIntegration(t *testing.T) {
	// Test the complete flow from alert parser to header parser
	alertsParser := &alertParser{}
	headerParser := &alertHeaderParser{}
	pc := parser.NewContext()

	// Start with alert parsing
	input := "> [!note]- Custom Title"
	reader := text.NewReader([]byte(input))
	parent := gast.NewDocument()

	alertNode, state := alertsParser.Open(parent, reader, pc)
	if alertNode == nil {
		t.Fatal("Failed to create alert node")
	}

	if state != parser.HasChildren {
		t.Fatalf("Expected HasChildren state, got %v", state)
	}

	// Now test header parsing
	headerInput := "]- Custom Title\n"
	headerReader := text.NewReader([]byte(headerInput))

	headerNode, headerState := headerParser.Open(alertNode, headerReader, pc)
	if headerNode == nil {
		t.Fatal("Failed to create header node")
	}

	if headerState != parser.NoChildren {
		t.Errorf("Expected NoChildren state for header, got %v", headerState)
	}

	// Verify attributes are properly transferred
	if kind, ok := headerNode.AttributeString("kind"); !ok {
		t.Error("Expected kind attribute in header node")
	} else if kind != "note" {
		t.Errorf("Expected kind='note', got %s", kind)
	}

	if shouldFold, ok := headerNode.AttributeString("shouldfold"); !ok {
		t.Error("Expected shouldfold attribute in header node")
	} else if !shouldFold.(bool) {
		t.Error("Expected shouldfold=true")
	}
}
