package alertcallouts

// This file contains streamlined integration and core functionality tests.
// Detailed functionality tests are split into separate files:
// - alertcallouts_core_test.go: Core alert functionality
// - alertcallouts_options_test.go: Configuration and options
// - alertcallouts_integration_test.go: End-to-end integration tests
// Internal package unit tests are in their respective internal/ directories.

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

var mdIconEmptySVG = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			WithIcons(map[string]string{"note": "<svg></svg>"}),
			WithFolding(true),
		),
	),
)

type TestCase struct {
	desc string
	md   string
	html string
}

// Essential functionality tests - focused on the most important behaviors
func TestAlerts(t *testing.T) {
	essentialTests := []TestCase{
		{
			desc: "Basic alert",
			md: `> [!note]
> Simple alert content`,
			html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Simple alert content</p>
</div>
</div>`,
		},
		{
			desc: "Alert with custom title",
			md:   `> [!info] Custom Title`,
			html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Custom Title</p>
</div>

</div>`,
		},
		{
			desc: "Foldable alert",
			md: `> [!warning]- Closed alert
> Content here`,
			html: `<details class="callout callout-foldable callout-warning" data-callout="warning"><summary class="callout-title">
<svg></svg><p class="callout-title-text">Closed alert</p>
</summary>
<div class="callout-body"><p>Content here</p>
</div>
</details>`,
		},
		{
			desc: "Regular blockquote (not alert)",
			md:   `> This is a regular blockquote`,
			html: `<blockquote>
<p>This is a regular blockquote</p>
</blockquote>
`,
		},
		{
			desc: "Invalid alert syntax",
			md:   `> [!info invalid syntax`,
			html: `<blockquote>
<p>[!info invalid syntax</p>
</blockquote>
`,
		},
		{
			desc: "Alert with list",
			md: `> [!tip]
> - Item one
> - Item two`,
			html: `<div class="callout callout-tip" data-callout="tip"><div class="callout-title">
<svg></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><ul>
<li>Item one</li>
<li>Item two</li>
</ul>
</div>
</div>`,
		},
	}

	for _, tc := range essentialTests {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdIconEmptySVG, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// AST node creation test
func TestASTNodeCreation(t *testing.T) {
	testCases := []struct {
		name string
		md   string
	}{
		{"SimpleAlert", `> [!note] Test AST`},
		{"AlertWithBody", `> [!warning]
> Body content`},
		{"ClosedAlert", `> [!info]- Closed alert`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf strings.Builder
			err := mdIconEmptySVG.Convert([]byte(tc.md), &buf)
			if err != nil {
				t.Errorf("Failed to parse markdown: %s, error: %v", tc.md, err)
			}

			if buf.Len() == 0 {
				t.Errorf("No HTML generated for markdown: %s", tc.md)
			}
		})
	}
}

// Extension registration test
func TestExtensionRegistration(t *testing.T) {
	ext := NewAlertCallouts(
		WithIcons(map[string]string{"test": "icon"}),
		WithFolding(false),
	)

	md := goldmark.New(goldmark.WithExtensions(ext))

	var buf strings.Builder
	err := md.Convert([]byte(`> [!test] Extension test`), &buf)
	if err != nil {
		t.Errorf("Extension registration failed: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Extension registration failed: no output")
	}
}

// Performance benchmarks
func BenchmarkSimpleAlert(b *testing.B) {
	md := `> [!info] Simple alert
> Content here`

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := mdIconEmptySVG.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}

func BenchmarkComplexAlert(b *testing.B) {
	md := `> [!warning] Complex alert with content
> This is a paragraph with *emphasis*.
>
> - List item 1
> - List item 2
>
> ` + "```" + `
> code block
> ` + "```"

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := mdIconEmptySVG.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}
