package alertcallouts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Test extension using GFMStrict icons and folding enabled
var mdGFMStrictWithFolding = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
			WithFolding(true),
		),
	),
)

// Test extension using GFMStrict icons and folding should be disabled by default
var mdGFMStrict = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
		),
	),
)

// TestGFMPlusPrimaryCallouts tests the primary callouts from the GFM Plus icon set
func TestGFMStrictPrimaryCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Primary Note callout",
			md: `> [!NOTE]
> This is a note callout with informational content.`,
			html: `<div class="callout callout-note iconset-gfm" data-callout="note"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>This is a note callout with informational content.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Tip callout",
			md: `> [!TIP]
> This is a tip callout with helpful suggestions.`,
			html: `<div class="callout callout-tip iconset-gfm" data-callout="tip"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-lightbulb-icon lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><p>This is a tip callout with helpful suggestions.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Important callout",
			md: `> [!IMPORTANT]
> This is an important callout highlighting crucial information.`,
			html: `<div class="callout callout-important iconset-gfm" data-callout="important"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</div>
<div class="callout-body"><p>This is an important callout highlighting crucial information.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Warning callout",
			md: `> [!WARNING]
> This is a warning callout about potential issues.`,
			html: `<div class="callout callout-warning iconset-gfm" data-callout="warning"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-triangle-alert-icon lucide-triangle-alert"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This is a warning callout about potential issues.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Caution callout",
			md: `> [!CAUTION]
> This is a caution callout for dangerous situations.`,
			html: `<div class="callout callout-caution iconset-gfm" data-callout="caution"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-octagon-alert-icon lucide-octagon-alert"><path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/></svg><p class="callout-title-text">Caution</p>
</div>
<div class="callout-body"><p>This is a caution callout for dangerous situations.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrict, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMStrictAliasCallouts tests alias callouts that reference primary callouts
func TestGFMStrictAliasCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "This uses the undefined INFO alias.",
			md: `> [!INFO]
> This uses the INFO alias.`,
			html: `<blockquote>
<p>[!INFO]
This uses the INFO alias.</p>
</blockquote>
`,
		},
		{
			desc: "Undefined HINT alias with a Tip custom title.",
			md: `> [!HINT] Tip
> Undefined HINT alias with a Tip custom title.`,
			html: `<blockquote>
<p>[!HINT] Tip
Undefined HINT alias with a Tip custom title.</p>
</blockquote>
`,
		},
		{
			desc: "This uses the undefined WARN alias.",
			md: `> [!WARN]
> This uses the undefined WARN alias.`,
			html: `<blockquote>
<p>[!WARN]
This uses the undefined WARN alias.</p>
</blockquote>
`,
		},
		{
			desc: "This uses the undefined ERROR alias.",
			md: `> [!ERROR]
> This uses the undefined ERROR alias.`,
			html: `<blockquote>
<p>[!ERROR]
This uses the undefined ERROR alias.</p>
</blockquote>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrict, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMStrictFoldingExamples tests the folding functionality with GFM Strict callouts
func TestGFMStrictFoldingExamples(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Closed by Default folding",
			md: `> [!IMPORTANT]-
> This important callout is closed by default due to the minus sign.`,
			html: `<details class="callout callout-foldable callout-important iconset-gfm" data-callout="important"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is closed by default due to the minus sign.</p>
</div>
</details>
`,
		},
		{
			desc: "Open by Default folding (Explicit)",
			md: `> [!IMPORTANT]+
> This important callout is explicitly marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-important iconset-gfm" data-callout="important" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is explicitly marked as open by default with the plus sign.</p>
</div>
</details>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrictWithFolding, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMStrictCustomTitles tests custom titles functionality with GFM Strict callouts
func TestGFMStrictCustomTitles(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Disallowed Undefined SUCCESS alias with custom title",
			md: `> [!SUCCESS] MiSsIoN AcCoMpLiShEd
> Disallowed Undefined SUCCESS alias with custom title.`,
			html: `<blockquote>
<p>[!SUCCESS] MiSsIoN AcCoMpLiShEd
Disallowed Undefined SUCCESS alias with custom title.</p>
</blockquote>
`,
		},
		{
			desc: "Disallowed Undefined FOO Callout",
			md: `> [!FOO]
> Disallowed Undefined FOO Callout.`,
			html: `<blockquote>
<p>[!FOO]
Disallowed Undefined FOO Callout.</p>
</blockquote>
`,
		},
		{
			desc: "Disallowed NOICON Callout",
			md: `> [!NoIcon]
> This is a disallowed NOICON callout.`,
			html: `<blockquote>
<p>[!NoIcon]
This is a disallowed NOICON callout.</p>
</blockquote>
`,
		},
		{
			desc: "Disallowed NOICON Callout with Unrecognized Custom Title",
			md: `> [!NoIcon] FooBar
> Disallowed NOICON Callout with Unrecognized Custom Title`,
			html: `<blockquote>
<p>[!NoIcon] FooBar
Disallowed NOICON Callout with Unrecognized Custom Title</p>
</blockquote>
`,
		},
		{
			desc: "Disallowed NOICON Callout with Recognized Custom Title",
			md: `> [!NoIcon] Warning
> Disallowed NOICON Callout with Recognized Custom Title`,
			html: `<blockquote>
<p>[!NoIcon] Warning
Disallowed NOICON Callout with Recognized Custom Title</p>
</blockquote>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrict, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}
