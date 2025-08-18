package alertcallouts

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Core functionality tests - essential alert behavior
func TestAlertCalloutsCore(t *testing.T) {
	mdTest := goldmark.New(
		goldmark.WithExtensions(
			&alertCalloutsOptions{
				Icons:          map[string]string{"note": "<svg></svg>", "info": "<svg></svg>", "warning": "<svg></svg>", "tip": "<svg></svg>"},
				FoldingEnabled: true,
			},
		),
	)

	testCases := []TestCase{
		{
			desc: "Basic alert",
			md: `> [!note]
> Paragraph content`,
			html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Paragraph content</p>
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
			desc: "Foldable alert closed",
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
			desc: "Foldable alert open",
			md: `> [!tip]+ Open alert
> Content here`,
			html: `<details class="callout callout-foldable callout-tip" data-callout="tip" open><summary class="callout-title">
<svg></svg><p class="callout-title-text">Open alert</p>
</summary>
<div class="callout-body"><p>Content here</p>
</div>
</details>`,
		},
		{
			desc: "Valid alert name - names CAN contain underscores",
			md: `> [!tip_one]
> Content here`,
			html: `<div class="callout callout-tip_one" data-callout="tip_one"><div class="callout-title">
<svg></svg><p class="callout-title-text">Tip_one</p>
</div>
<div class="callout-body"><p>Content here</p>
</div>
</div>`,
		},
		{
			desc: "Invalid alert name - names cannot contain dashes",
			md: `> [!tip-one]
> Content here`,
			html: `<blockquote>
<p>[!tip-one]
Content here</p>
</blockquote>`,
		},
		{
			desc: "Not an alert (regular blockquote)",
			md:   `> This is a blockquote`,
			html: `<blockquote>
<p>This is a blockquote</p>
</blockquote>
`,
		},
		{
			desc: "Invalid alert syntax",
			md:   `> [!info This is not a alert`,
			html: `<blockquote>
<p>[!info This is not a alert</p>
</blockquote>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdTest, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// Test AST node functionality
func TestASTNodeCreationCore(t *testing.T) {
	mdTest := goldmark.New(
		goldmark.WithExtensions(
			&alertCalloutsOptions{
				Icons:          map[string]string{"note": "<svg></svg>", "warning": "<svg></svg>", "info": "<svg></svg>"},
				FoldingEnabled: true,
			},
		),
	)

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
			err := mdTest.Convert([]byte(tc.md), &buf)
			if err != nil {
				t.Errorf("Failed to parse markdown: %s, error: %v", tc.md, err)
			}

			if buf.Len() == 0 {
				t.Errorf("No HTML generated for markdown: %s", tc.md)
			}
		})
	}
}

// Test extension registration
func TestExtensionRegistrationCore(t *testing.T) {
	ext := &alertCalloutsOptions{
		Icons:          map[string]string{"test": "icon"},
		FoldingEnabled: false,
	}

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

// Benchmark tests for performance
func BenchmarkSimpleAlertCore(b *testing.B) {
	mdTest := goldmark.New(
		goldmark.WithExtensions(
			&alertCalloutsOptions{
				Icons:          map[string]string{"info": "<svg></svg>"},
				FoldingEnabled: true,
			},
		),
	)

	md := `> [!info] Simple alert
> Content here`

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := mdTest.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}

func BenchmarkComplexAlertCore(b *testing.B) {
	mdTest := goldmark.New(
		goldmark.WithExtensions(
			&alertCalloutsOptions{
				Icons:          map[string]string{"warning": "<svg></svg>"},
				FoldingEnabled: true,
			},
		),
	)

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
		err := mdTest.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}
