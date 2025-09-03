package parser

import (
	"strings"
	"testing"

	"github.com/zmtcreative/gm-alert-callouts/internal/ast"
	"github.com/zmtcreative/gm-alert-callouts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestNewAlertsParser(t *testing.T) {
	p := NewAlertsParser(make([]string, 0), false, false)
	if p == nil {
		t.Fatal("NewAlertsParser() returned nil")
	}

	// Verify it returns a properly configured alertParser
	alertParser, ok := p.(*alertParser)
	if !ok {
		t.Fatal("NewAlertsParser() should return *alertParser")
	}

	if alertParser.CustomAlertsEnabled != false {
		t.Error("Expected CustomAlertsEnabled to be false")
	}

	if len(alertParser.IconList) != 0 {
		t.Error("Expected empty IconList")
	}
}

func TestAlertsParserTrigger(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}
	trigger := p.Trigger()
	expected := []byte{'>'}

	if len(trigger) != 1 || trigger[0] != expected[0] {
		t.Errorf("Expected trigger %v, got %v", expected, trigger)
	}
}

func TestAlertsParserProcess(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}

	testCases := []struct {
		name     string
		input    string
		expected bool
		advance  int
	}{
		{
			name:     "Valid blockquote start",
			input:    "> [!note] Test",
			expected: true,
			advance:  2, // '>' + ' '
		},
		{
			name:     "Valid blockquote with tab",
			input:    ">\t[!note] Test",
			expected: true,
			advance:  2, // '>' + '\t'
		},
		{
			name:     "Valid empty blockquote",
			input:    ">",
			expected: true,
			advance:  1,
		},
		{
			name:     "Invalid - no blockquote marker",
			input:    "[!note] Test",
			expected: false,
			advance:  0,
		},
		{
			name:     "Invalid - too much indentation",
			input:    "    > [!note] Test",
			expected: false,
			advance:  0,
		},
		{
			name:     "Invalid - 4 spaces indentation",
			input:    "    > [!note] Test",
			expected: false,
			advance:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := text.NewReader([]byte(tc.input))
			ok, advanceBy := p.process(reader)

			if ok != tc.expected {
				t.Errorf("Expected process result %v, got %v", tc.expected, ok)
			}

			if advanceBy != tc.advance {
				t.Errorf("Expected advance %d, got %d", tc.advance, advanceBy)
			}
		})
	}
}

func TestAlertsParserOpenNoCustomAlertsNoFolding(t *testing.T) {
	p := &alertParser{[]string{"note", "warning", "info", "tip"}, true, true}
	pc := parser.NewContext()

	testCases := []struct {
		name     string
		input    string
		expected bool
		kind     string
		title    string
		closed   bool
		fold     bool
	}{
		{
			name:     "Basic note alert",
			input:    "> [!note] Test Title",
			expected: true,
			kind:     "note",
			title:    "Test Title",
			closed:   false,
			fold:     false,
		},
		{
			name:     "Warning alert with no title",
			input:    "> [!warning]",
			expected: true,
			kind:     "warning",
			title:    "",
			closed:   false,
			fold:     false,
		},
		{
			name:     "Closed alert",
			input:    "> [!info]- Closed Alert",
			expected: true,
			kind:     "info",
			title:    "Closed Alert",
			closed:   true,
			fold:     true,
		},
		{
			name:     "Open alert",
			input:    "> [!tip]+ Open Alert",
			expected: true,
			kind:     "tip",
			title:    "Open Alert",
			closed:   false,
			fold:     true,
		},
		{
			name:     "Regular blockquote",
			input:    "> This is not an alert",
			expected: false,
			kind:     "",
			title:    "",
			closed:   false,
			fold:     false,
		},
		{
			name:     "Invalid alert syntax",
			input:    "> [!note incomplete",
			expected: false,
			kind:     "",
			title:    "",
			closed:   false,
			fold:     false,
		},
		{
			name:     "Empty line",
			input:    "> ",
			expected: false,
			kind:     "",
			title:    "",
			closed:   false,
			fold:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := text.NewReader([]byte(tc.input))
			parent := gast.NewDocument()

			node, state := p.Open(parent, reader, pc)

			if tc.expected {
				if node == nil {
					t.Fatalf("Expected node to be created, got nil")
				}

				if node.Kind() != constants.KindAlerts {
					t.Errorf("Expected KindAlerts, got %v", node.Kind())
				}

				if state != parser.HasChildren {
					t.Errorf("Expected HasChildren state, got %v", state)
				}

				// Check attributes
				if kind, ok := node.AttributeString("kind"); ok {
					if string(kind.([]uint8)) != tc.kind {
						t.Errorf("Expected kind %s, got %s", tc.kind, string(kind.([]uint8)))
					}
				} else {
					t.Error("Expected kind attribute")
				}

				if title, ok := node.AttributeString("title"); ok {
					actualTitle := string(title.([]uint8))
					if actualTitle != tc.title {
						t.Errorf("Expected title '%s', got '%s'", tc.title, actualTitle)
					}
				} else if tc.title != "" {
					t.Error("Expected title attribute")
				}

				if closed, ok := node.AttributeString("closed"); ok {
					if bool(closed.(bool)) != tc.closed {
						t.Errorf("Expected closed %v, got %v", tc.closed, bool(closed.(bool)))
					}
				}

				if shouldfold, ok := node.AttributeString("shouldfold"); ok {
					if bool(shouldfold.(bool)) != tc.fold {
						t.Errorf("Expected shouldfold %v, got %v", tc.fold, bool(shouldfold.(bool)))
					}
				}
			} else {
				if node != nil {
					t.Errorf("Expected nil node, got %v", node)
				}

				if state != parser.NoChildren {
					t.Errorf("Expected NoChildren state, got %v", state)
				}
			}
		})
	}
}

func TestAlertsParserContinue(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}
	pc := parser.NewContext()
	node := ast.NewAlerts()

	testCases := []struct {
		name     string
		input    string
		expected parser.State
	}{
		{
			name:     "Valid continuation",
			input:    "> Content line",
			expected: parser.Continue | parser.HasChildren,
		},
		{
			name:     "End of blockquote",
			input:    "Not a blockquote line",
			expected: parser.Close,
		},
		{
			name:     "Another valid line",
			input:    ">   Indented content",
			expected: parser.Continue | parser.HasChildren,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := text.NewReader([]byte(tc.input))
			state := p.Continue(node, reader, pc)

			if state != tc.expected {
				t.Errorf("Expected state %v, got %v", tc.expected, state)
			}
		})
	}
}

func TestAlertsParserClose(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}
	pc := parser.NewContext()

	// Create a mock alert node with header and body children
	alert := ast.NewAlerts()
	header := ast.NewAlertsHeader()
	header.SetAttributeString("kind", "note")

	// Add some regular paragraph nodes that should become body content
	para1 := gast.NewParagraph()
	para2 := gast.NewParagraph()

	alert.AppendChild(alert, header)
	alert.AppendChild(alert, para1)
	alert.AppendChild(alert, para2)

	// Verify initial state
	if alert.ChildCount() != 3 {
		t.Fatalf("Expected 3 children initially, got %d", alert.ChildCount())
	}

	// Close the parser
	reader := text.NewReader([]byte(""))
	p.Close(alert, reader, pc)

	// Verify restructuring
	if alert.ChildCount() != 2 {
		t.Errorf("Expected 2 children after close (header + body), got %d", alert.ChildCount())
	}

	// First child should be the header
	firstChild := alert.FirstChild()
	if firstChild.Kind() != constants.KindAlertsHeader {
		t.Errorf("Expected first child to be AlertsHeader, got %v", firstChild.Kind())
	}

	// Second child should be the body
	secondChild := firstChild.NextSibling()
	if secondChild == nil {
		t.Fatal("Expected second child")
	}
	if secondChild.Kind() != constants.KindAlertsBody {
		t.Errorf("Expected second child to be AlertsBody, got %v", secondChild.Kind())
	}

	// Body should contain the two paragraphs
	if secondChild.ChildCount() != 2 {
		t.Errorf("Expected body to have 2 children, got %d", secondChild.ChildCount())
	}
}

func TestAlertsParserCanInterruptParagraph(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}
	if !p.CanInterruptParagraph() {
		t.Error("Expected CanInterruptParagraph to return true")
	}
}

func TestAlertsParserCanAcceptIndentedLine(t *testing.T) {
	p := &alertParser{[]string{"note"}, false, false}
	if p.CanAcceptIndentedLine() {
		t.Error("Expected CanAcceptIndentedLine to return false")
	}
}

func TestAlertsParserIntegration(t *testing.T) {
	// Test the parser with a complete alert structure
	p := &alertParser{[]string{"warning"}, true, true}
	pc := parser.NewContext()
	parent := gast.NewDocument()

	input := `> [!warning]- Important Warning
> This is the content of the warning.
>
> - Item 1
> - Item 2`

	lines := strings.Split(input, "\n")
	reader := text.NewReader([]byte(lines[0]))

	// Test Open
	node, state := p.Open(parent, reader, pc)
	if node == nil {
		t.Fatal("Failed to open alert")
	}

	if state != parser.HasChildren {
		t.Errorf("Expected HasChildren state, got %v", state)
	}

	// Test Continue with subsequent lines
	for i := 1; i < len(lines); i++ {
		reader = text.NewReader([]byte(lines[i]))
		state = p.Continue(node, reader, pc)
		if state != (parser.Continue | parser.HasChildren) {
			t.Errorf("Line %d: Expected Continue|HasChildren, got %v", i, state)
		}
	}

	// Test with a line that should close the alert
	reader = text.NewReader([]byte("Not a blockquote"))
	state = p.Continue(node, reader, pc)
	if state != parser.Close {
		t.Errorf("Expected Close state, got %v", state)
	}
}

func TestRegexMatching(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "Basic note",
			input: "[!note]",
			expected: map[string]string{
				"kind":   "note",
				"closed": "",
				"opened": "",
				"title":  "",
			},
		},
		{
			name:  "Note with title",
			input: "[!note] My Title",
			expected: map[string]string{
				"kind":   "note",
				"closed": "",
				"opened": "",
				"title":  "My Title",
			},
		},
		{
			name:  "Closed alert",
			input: "[!warning]- Closed Warning",
			expected: map[string]string{
				"kind":   "warning",
				"closed": "-",
				"opened": "",
				"title":  "Closed Warning",
			},
		},
		{
			name:  "Open alert",
			input: "[!tip]+ Open Tip",
			expected: map[string]string{
				"kind":   "tip",
				"closed": "",
				"opened": "+",
				"title":  "Open Tip",
			},
		},
		{
			name:  "Just marker no space",
			input: "[!info]-",
			expected: map[string]string{
				"kind":   "info",
				"closed": "-",
				"opened": "",
				"title":  "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matches := regex.FindStringSubmatch(tc.input)
			if len(matches) == 0 {
				t.Fatal("No regex matches found")
			}

			subMatchMap := make(map[string]string)
			for i, name := range regex.SubexpNames() {
				if i != 0 && i < len(matches) && name != "" {
					subMatchMap[name] = matches[i]
				}
			}

			for key, expected := range tc.expected {
				if actual, exists := subMatchMap[key]; !exists || actual != expected {
					t.Errorf("Key %s: expected %s, got %s", key, expected, actual)
				}
			}
		})
	}
}
