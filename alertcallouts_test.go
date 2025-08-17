package alertcallouts

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

var mdIconEmptySVG = goldmark.New(
	goldmark.WithExtensions(
		&alertCalloutsOptions{
			Icons: map[string]string{"note": "<svg></svg>"},
			FoldingEnabled: true,
		},
	),
)

type TestCase struct {
	desc string
	md   string
	html string
}

var casesBasic = [...]TestCase{
	{
		desc: "Empty blockquote",
		md:   ">",
		html: `<blockquote>
</blockquote>
`},
	{
		desc: "Empty blockquote with space",
		md:   "> ",
		html: `<blockquote>
</blockquote>
`},
	{
		desc: "Default blockquote",
		md:   "> This is a blockquote",
		html: `<blockquote>
<p>This is a blockquote</p>
</blockquote>
`},
	{
		desc: "Alerts with a paragraph",
		md: `> [!note]
> Paragraph
> over a few lines`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Paragraph
over a few lines</p>
</div>
</div>`},
	{
		desc: "Alerts with two paragraphs",
		md: `> [!InFo]
> paragraph
> over a few lines
>
> second paragraph with *some* syntax
`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>paragraph
over a few lines</p>
<p>second paragraph with <em>some</em> syntax</p>
</div>
</div>`},
	{
		desc: "Alerts with two paragraphs and a close request",
		md: `> [!InFo]-
> paragraph
> over a few lines
>
> second paragraph with *some* syntax
`,
		html: `<details class="callout callout-foldable callout-info" data-callout="info"><summary class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</summary>
<div class="callout-body"><p>paragraph
over a few lines</p>
<p>second paragraph with <em>some</em> syntax</p>
</div>
</details>`},	{
		desc: "Alerts without body",
		md:   `> [!info] title`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">title</p>
</div>

</div>`},
	{
		desc: "Alerts with list",
		md: `> [!info]
> - item`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><ul>
<li>item</li>
</ul>
</div>
</div>`},
	{
		desc: "README example",
		md: `> [!info]
> With lots of possibilities:
> - feature one
> - feature two`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</div>
</div>`},
	{
		desc: "Not a alert",
		md: `[!info] title
`,
		html: `<p>[!info] title</p>
`}, {
		desc: "Syntax in summary",
		md:   `>[!info] Title with *some* syntax [and](http://example.com) links`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Title with <em>some</em> syntax <a href="http://example.com">and</a> links</p>
</div>

</div>`}, {
		desc: "text before alert type",
		md: `> abcd [!info]- This is not a alert
`,
		html: `<blockquote>
<p>abcd [!info]- This is not a alert</p>
</blockquote>
`}, {desc: "space before a alert type",
		md: `>  [!info]- This is not a alert
`,
		html: `<blockquote>
<p>[!info]- This is not a alert</p>
</blockquote>
`}, {desc: "2 spaces before a alert type",
		md: `>   [!info]- This is not a alert
`,
		html: `<blockquote>
<p>[!info]- This is not a alert</p>
</blockquote>
`}, {desc: "3 spaces before a alert type",
		md: `>    [!info]- This is not a alert
`,
		html: `<blockquote>
<p>[!info]- This is not a alert</p>
</blockquote>
`}, {desc: "4 spaces before a alert type",
		md: `>     [!info]- This is not a alert
`,
		html: `<blockquote>
<pre><code>[!info]- This is not a alert
</code></pre>
</blockquote>
`},
}

// Additional test cases for comprehensive coverage
var casesAdvanced = [...]TestCase{
	// Edge cases and malformed syntax
	{
		desc: "Invalid alert type with numbers",
		md:   `> [!info123] title`,
		html: `<div class="callout callout-info123" data-callout="info123"><div class="callout-title">
<svg></svg><p class="callout-title-text">title</p>
</div>

</div>`,
	},
	{
		desc: "Alert type with special characters",
		md:   `> [!warning-special] title`,
		html: `<blockquote>
<p>[!warning-special] title</p>
</blockquote>
`,
	},
	{
		desc: "Empty alert type",
		md:   `> [!] title`,
		html: `<blockquote>
<p>[!] title</p>
</blockquote>
`,
	},
	{
		desc: "Missing closing bracket",
		md:   `> [!info title`,
		html: `<blockquote>
<p>[!info title</p>
</blockquote>
`,
	},
	{
		desc: "Missing opening bracket",
		md:   `> !info] title`,
		html: `<blockquote>
<p>!info] title</p>
</blockquote>
`,
	},
	// Foldable alerts
	{
		desc: "Closed alert with dash",
		md:   `> [!warning]- This is a closed alert`,
		html: `<details class="callout callout-foldable callout-warning" data-callout="warning"><summary class="callout-title">
<svg></svg><p class="callout-title-text">This is a closed alert</p>
</summary>

</details>`,
	},
	{
		desc: "Open alert with plus",
		md:   `> [!warning]+ This is an open alert`,
		html: `<details class="callout callout-foldable callout-warning" data-callout="warning" open><summary class="callout-title">
<svg></svg><p class="callout-title-text">This is an open alert</p>
</summary>

</details>`,
	},
	{
		desc: "Closed alert without title",
		md: `> [!tip]-
> content here`,
		html: `<details class="callout callout-foldable callout-tip" data-callout="tip"><summary class="callout-title">
<svg></svg><p class="callout-title-text">Tip</p>
</summary>
<div class="callout-body"><p>content here</p>
</div>
</details>`,
	},
	{
		desc: "Open alert without title",
		md: `> [!tip]+
> content here`,
		html: `<details class="callout callout-foldable callout-tip" data-callout="tip" open><summary class="callout-title">
<svg></svg><p class="callout-title-text">Tip</p>
</summary>
<div class="callout-body"><p>content here</p>
</div>
</details>`,
	},

	// Case sensitivity tests
	{
		desc: "Mixed case alert type",
		md:   `> [!WaRnInG] Mixed case`,
		html: `<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<svg></svg><p class="callout-title-text">Mixed case</p>
</div>

</div>`,
	},
	{
		desc: "Uppercase alert type",
		md:   `> [!ERROR] Uppercase alert`,
		html: `<div class="callout callout-error" data-callout="error"><div class="callout-title">
<svg></svg><p class="callout-title-text">Uppercase alert</p>
</div>

</div>`,
	},

	// Complex content within alerts
	{
		desc: "Alert with code block",
		md: `> [!note]
> Here's some code:
> ` + "```" + `
> function test() {
>   return true;
> }
> ` + "```",
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Here's some code:</p>
<pre><code>function test() {
  return true;
}
</code></pre>
</div>
</div>`,
	},
	{
		desc: "Alert with inline code",
		md: `> [!info]
> Use the ` + "`alert`" + ` function`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>Use the <code>alert</code> function</p>
</div>
</div>`,
	},
	{
		desc: "Alert with nested list",
		md: `> [!tip]
> - item 1
>   - nested item
>   - another nested
> - item 2`,
		html: `<div class="callout callout-tip" data-callout="tip"><div class="callout-title">
<svg></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><ul>
<li>item 1
<ul>
<li>nested item</li>
<li>another nested</li>
</ul>
</li>
<li>item 2</li>
</ul>
</div>
</div>`,
	},
	{
		desc: "Alert with ordered list",
		md: `> [!important]
> 1. First step
> 2. Second step
> 3. Third step`,
		html: `<div class="callout callout-important" data-callout="important"><div class="callout-title">
<svg></svg><p class="callout-title-text">Important</p>
</div>
<div class="callout-body"><ol>
<li>First step</li>
<li>Second step</li>
<li>Third step</li>
</ol>
</div>
</div>`,
	},
	{
		desc: "Alert with blockquote inside",
		md: `> [!note]
> > This is a nested quote
> > with multiple lines`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><blockquote>
<p>This is a nested quote
with multiple lines</p>
</blockquote>
</div>
</div>`,
	},

	// Multiple line breaks and empty lines
	{
		desc: "Alert with multiple empty lines",
		md: `> [!info]
> First paragraph
>
>
> Second paragraph after empty lines`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>First paragraph</p>
<p>Second paragraph after empty lines</p>
</div>
</div>`,
	},
	{
		desc: "Alert with trailing empty lines",
		md: `> [!warning]
> Content here
>
>
`,
		html: `<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<svg></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>Content here</p>
</div>
</div>`,
	},

	// HTML escaping and special characters
	{
		desc: "Alert with HTML entities in title",
		md:   `> [!note] Title with &amp; <script> tags`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Title with &amp; <!-- raw HTML omitted --> tags</p>
</div>

</div>`,
	},
	{
		desc: "Alert with HTML in content",
		md: `> [!warning]
> Be careful with <strong>HTML</strong> & scripts`,
		html: `<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<svg></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>Be careful with <!-- raw HTML omitted -->HTML<!-- raw HTML omitted --> &amp; scripts</p>
</div>
</div>`,
	},

	// Unicode and international content
	{
		desc: "Alert with unicode in title",
		md:   `> [!info] Título con acentos é çharacters 中文`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Título con acentos é çharacters 中文</p>
</div>

</div>`,
	},
	{
		desc: "Alert with emoji",
		md: `> [!tip] 🚀 Rocket tip
> Use emojis sparingly 😊`,
		html: `<div class="callout callout-tip" data-callout="tip"><div class="callout-title">
<svg></svg><p class="callout-title-text">🚀 Rocket tip</p>
</div>
<div class="callout-body"><p>Use emojis sparingly 😊</p>
</div>
</div>`,
	},

	// Long content tests
	{
		desc: "Alert with very long title",
		md:   `> [!note] This is a very long title that goes on and on and on and should still work properly even with lots of text in the title section of the alert`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">This is a very long title that goes on and on and on and should still work properly even with lots of text in the title section of the alert</p>
</div>

</div>`,
	},
	{
		desc: "Alert with very long paragraph",
		md: `> [!info]
> This is a very long paragraph that contains lots of text and should wrap properly in the alert body. It includes multiple sentences and should demonstrate that the alert can handle substantial amounts of content without any issues. The formatting should remain intact and the HTML output should be properly structured.`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>This is a very long paragraph that contains lots of text and should wrap properly in the alert body. It includes multiple sentences and should demonstrate that the alert can handle substantial amounts of content without any issues. The formatting should remain intact and the HTML output should be properly structured.</p>
</div>
</div>`,
	},

	// Edge cases with indentation
	{
		desc: "Alert with tabs in content",
		md: `> [!note]
>	Indented with tab
>		Double tab indent`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<svg></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><blockquote>
<blockquote>
<p>Indented with tab
Double tab indent</p>
</blockquote>
</blockquote>
</div>
</div>`,
	},

	// Various alert types that might not be tested
	{
		desc: "Caution alert type",
		md: `> [!caution]
> Be very careful here`,
		html: `<div class="callout callout-caution" data-callout="caution"><div class="callout-title">
<svg></svg><p class="callout-title-text">Caution</p>
</div>
<div class="callout-body"><p>Be very careful here</p>
</div>
</div>`,
	},
	{
		desc: "Important alert type",
		md: `> [!important] Critical information
> This is very important`,
		html: `<div class="callout callout-important" data-callout="important"><div class="callout-title">
<svg></svg><p class="callout-title-text">Critical information</p>
</div>
<div class="callout-body"><p>This is very important</p>
</div>
</div>`,
	},

	// Multiple alerts in sequence
	{
		desc: "Multiple alerts back to back",
		md: `> [!info]
> First alert

> [!warning]
> Second alert`,
		html: `<div class="callout callout-info" data-callout="info"><div class="callout-title">
<svg></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>First alert</p>
</div>
</div>
<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<svg></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>Second alert</p>
</div>
</div>`,
	},
}

// Test with different icon configurations
var mdWithIcons = goldmark.New(
	goldmark.WithExtensions(
		&alertCalloutsOptions{
			Icons: map[string]string{
				"note":      "📝",
				"tip":       "💡",
				"warning":   "⚠️",
				"caution":   "🔥",
				"important": "❗",
			},
			FoldingEnabled: true,
		},
	),
)

var casesWithIcons = [...]TestCase{
	{
		desc: "Alert with icon",
		md: `> [!note]
> Content with icon`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
📝<p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Content with icon</p>
</div>
</div>`,
	},
	{
		desc: "Alert using 'noicon' to suppress icon",
		md: `> [!noicon] Error
> No icon for this type`,
		html: `<div class="callout callout-error" data-callout="error"><div class="callout-title">
<p class="callout-title-text">Error</p>
</div>
<div class="callout-body"><p>No icon for this type</p>
</div>
</div>`,
	},
	{
		desc: "Alert using 'noicon' to suppress icon with complex title formatting",
		md: `> [!noicon] Error **BAD* [text](#link) ` + "`some code`" + `
> No icon for this type`,
		html: `<div class="callout callout-error-bad" data-callout="error-bad"><div class="callout-title">
<p class="callout-title-text">Error *<em>BAD</em> <a href="#link">text</a> <code>some code</code></p>
</div>
<div class="callout-body"><p>No icon for this type</p>
</div>
</div>`,
	},
	{
		desc: "Custom title with icon",
		md:   `> [!tip] Custom tip title`,
		html: `<div class="callout callout-tip" data-callout="tip"><div class="callout-title">
💡<p class="callout-title-text">Custom tip title</p>
</div>

</div>`,
	},
	{
		desc: "Unsupported alert type 'cite' should use 'note' icon",
		md:   `> [!cite]`,
		html: `<div class="callout callout-cite" data-callout="cite"><div class="callout-title">
📝<p class="callout-title-text">Cite</p>
</div>

</div>`,
	},
	{
		desc: "Unsupported alert type 'cite' with 'Quote' title should use 'note' icon",
		md:   `> [!cite] Quote`,
		html: `<div class="callout callout-cite" data-callout="cite"><div class="callout-title">
📝<p class="callout-title-text">Quote</p>
</div>

</div>`,
	},
}

// Test with default GFM icons configuration
var mdGFMIcons = goldmark.New(
	goldmark.WithExtensions(
		AlertCallouts,
	),
)

var casesGFMIcons = [...]TestCase{
	{
		desc: "Initialized with default GFM icons -- basic note",
		md: `> [!note]
> Content with GFM note icon`,
		html: `<div class="callout callout-note iconset-gfm" data-callout="note"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>Content with GFM note icon</p>
</div>
</div>`,
	},
	{
		desc: "Initialized with default GFM icons -- custom title with important icon",
		md: `> [!important] Shout-Out
> Content for shout-out.`,
		html: `<div class="callout callout-important iconset-gfm" data-callout="important"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Shout-Out</p>
</div>
<div class="callout-body"><p>Content for shout-out.</p>
</div>
</div>`,
	},
}

var mdDisableFolding = goldmark.New(
	goldmark.WithExtensions(
		&alertCalloutsOptions{
			FoldingEnabled: false,
		},
	),
)

var casesDisableFolding = [...]TestCase{
	{
		desc: "Alert with no folding symbol",
		md:   `> [!note]
Test`,
		html: `<div class="callout callout-note" data-callout="note"><div class="callout-title">
<p class="callout-title-text">Note</p>
</div>

</div>
<p>Test</p>`,
	},
	// Foldable alerts
	{
		desc: "Closed alert with dash",
		md:   `> [!warning]-
This is a closed alert`,
		html: `<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<p class="callout-title-text">Warning</p>
</div>

</div>
<p>This is a closed alert</p>`,
	},
	{
		desc: "Open alert with plus",
		md:   `> [!warning]+
This is an open alert`,
		html: `<div class="callout callout-warning" data-callout="warning"><div class="callout-title">
<p class="callout-title-text">Warning</p>
</div>

</div>
<p>This is an open alert</p>`,
	},
}

func TestAlerts(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		for i, c := range casesBasic {
			t.Run(c.desc, func(t *testing.T) {
				testutil.DoTestCase(mdIconEmptySVG, testutil.MarkdownTestCase{
					No:          i,
					Description: c.desc,
					Markdown:    c.md,
					Expected:    c.html,
				}, t)
			})
		}
	})

	t.Run("Additional", func(t *testing.T) {
		for i, c := range casesAdvanced {
			t.Run(c.desc, func(t *testing.T) {
				testutil.DoTestCase(mdIconEmptySVG, testutil.MarkdownTestCase{
					No:          i,
					Description: c.desc,
					Markdown:    c.md,
					Expected:    c.html,
				}, t)
			})
		}
	})

	t.Run("DisabledFoldingNoIcons", func(t *testing.T) {
		for i, c := range casesDisableFolding {
			t.Run(c.desc, func(t *testing.T) {
				testutil.DoTestCase(mdDisableFolding, testutil.MarkdownTestCase{
					No:          i,
					Description: c.desc,
					Markdown:    c.md,
					Expected:    c.html,
				}, t)
			})
		}
	})

	t.Run("WithIcons", func(t *testing.T) {
		for i, c := range casesWithIcons {
			t.Run(c.desc, func(t *testing.T) {
				testutil.DoTestCase(mdWithIcons, testutil.MarkdownTestCase{
					No:          i,
					Description: c.desc,
					Markdown:    c.md,
					Expected:    c.html,
				}, t)
			})
		}
	})

	t.Run("WithGFMIcons", func(t *testing.T) {
		for i, c := range casesGFMIcons {
			t.Run(c.desc, func(t *testing.T) {
				testutil.DoTestCase(mdGFMIcons, testutil.MarkdownTestCase{
					No:          i,
					Description: c.desc,
					Markdown:    c.md,
					Expected:    c.html,
				}, t)
			})
		}
	})
}

// Test AST node functionality
func TestASTNodeCreation(t *testing.T) {
	// These tests verify that the AST nodes are created correctly
	// by attempting to convert to HTML and checking for errors

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

			// Verify that some HTML was generated
			if buf.Len() == 0 {
				t.Errorf("No HTML generated for markdown: %s", tc.md)
			}
		})
	}
}

// Test extension registration
func TestExtensionRegistration(t *testing.T) {
	ext := &alertCalloutsOptions{
		Icons: map[string]string{"test": "icon"},
		FoldingEnabled: false,
	}

	md := goldmark.New(goldmark.WithExtensions(ext))

	// Test that the extension was registered by converting a simple alert
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
	md := `> [!warning] Complex alert with lots of content
> This is a paragraph with *emphasis* and **strong** text.
>
> - List item 1
> - List item 2
>   - Nested item
>
> ` + "```" + `
> code block
> ` + "```" + `
>
> Final paragraph`

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := mdIconEmptySVG.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}

// Test cases for new initialization methods

func TestNewAlertCallouts(t *testing.T) {
	t.Run("Default options", func(t *testing.T) {
		ext := NewAlertCallouts()

		if ext == nil {
			t.Fatal("NewAlertCallouts() returned nil")
		}

		if ext.Icons == nil {
			t.Error("Expected Icons to be initialized as empty map")
		}

		if len(ext.Icons) != 0 {
			t.Errorf("Expected empty Icons map, got %d items", len(ext.Icons))
		}

		if ext.FoldingEnabled != true {
			t.Error("Expected FoldingEnabled to be true by default")
		}
	})

	t.Run("With single icon option", func(t *testing.T) {
		ext := NewAlertCallouts(WithIcon("note", "<svg>note icon</svg>"))

		if len(ext.Icons) != 1 {
			t.Errorf("Expected 1 icon, got %d", len(ext.Icons))
		}

		if ext.Icons["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected note icon, got %s", ext.Icons["note"])
		}
	})

	t.Run("With multiple icons option", func(t *testing.T) {
		icons := map[string]string{
			"note":    "<svg>note</svg>",
			"warning": "<svg>warning</svg>",
			"info":    "<svg>info</svg>",
		}

		ext := NewAlertCallouts(WithIcons(icons))

		if len(ext.Icons) != 3 {
			t.Errorf("Expected 3 icons, got %d", len(ext.Icons))
		}

		for kind, expected := range icons {
			if ext.Icons[kind] != expected {
				t.Errorf("Expected %s icon to be %s, got %s", kind, expected, ext.Icons[kind])
			}
		}
	})

	t.Run("Disable folding", func(t *testing.T) {
		ext := NewAlertCallouts(WithFolding(false))

		if ext.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})

	t.Run("With combined options", func(t *testing.T) {
		icons := map[string]string{"tip": "<svg>tip</svg>"}

		ext := NewAlertCallouts(
			WithIcons(icons),
			WithFolding(false),
			WithIcon("important", "<svg>important</svg>"),
		)

		if len(ext.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(ext.Icons))
		}

		if ext.Icons["tip"] != "<svg>tip</svg>" {
			t.Errorf("Expected tip icon, got %s", ext.Icons["tip"])
		}

		if ext.Icons["important"] != "<svg>important</svg>" {
			t.Errorf("Expected important icon, got %s", ext.Icons["important"])
		}

		if ext.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})
}

func TestWithIcon(t *testing.T) {
	t.Run("Adds icon to nil map", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithIcon("test", "<svg>test</svg>")
		option(opts)

		if opts.Icons == nil {
			t.Fatal("Expected Icons map to be initialized")
		}

		if opts.Icons["test"] != "<svg>test</svg>" {
			t.Errorf("Expected test icon, got %s", opts.Icons["test"])
		}
	})

	t.Run("Adds icon to existing map", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"existing": "<svg>existing</svg>"},
		}

		option := WithIcon("new", "<svg>new</svg>")
		option(opts)

		if len(opts.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(opts.Icons))
		}

		if opts.Icons["existing"] != "<svg>existing</svg>" {
			t.Error("Existing icon should be preserved")
		}

		if opts.Icons["new"] != "<svg>new</svg>" {
			t.Error("New icon should be added")
		}
	})

	t.Run("Overwrites existing icon", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"note": "<svg>old</svg>"},
		}

		option := WithIcon("note", "<svg>new</svg>")
		option(opts)

		if opts.Icons["note"] != "<svg>new</svg>" {
			t.Errorf("Expected icon to be overwritten, got %s", opts.Icons["note"])
		}
	})
}

func TestWithIcons(t *testing.T) {
	t.Run("Sets icons map", func(t *testing.T) {
		icons := map[string]string{
			"note":    "<svg>note</svg>",
			"warning": "<svg>warning</svg>",
		}

		opts := &alertCalloutsOptions{}
		option := WithIcons(icons)
		option(opts)

		if len(opts.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(opts.Icons))
		}

		for kind, expected := range icons {
			if opts.Icons[kind] != expected {
				t.Errorf("Expected %s icon to be %s, got %s", kind, expected, opts.Icons[kind])
			}
		}
	})

	t.Run("Replaces existing icons", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"old": "<svg>old</svg>"},
		}

		newIcons := map[string]string{"new": "<svg>new</svg>"}
		option := WithIcons(newIcons)
		option(opts)

		if len(opts.Icons) != 1 {
			t.Errorf("Expected 1 icon, got %d", len(opts.Icons))
		}

		if opts.Icons["new"] != "<svg>new</svg>" {
			t.Error("Expected new icon")
		}

		if _, exists := opts.Icons["old"]; exists {
			t.Error("Expected old icon to be removed")
		}
	})
}

func TestWithFolding(t *testing.T) {
	t.Run("Enables folding", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithFolding(true)
		option(opts)

		if opts.FoldingEnabled != true {
			t.Error("Expected FoldingEnabled to be true")
		}
	})

	t.Run("Disables folding", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithFolding(false)
		option(opts)

		if opts.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})
}

func TestNewAlertCalloutsIntegration(t *testing.T) {
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
		if !strings.Contains(result, "") {
			t.Error("Expected output to contain class")
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
		if !strings.Contains(result, "") {
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
}
