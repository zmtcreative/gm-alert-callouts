package alertcallouts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Test extension using GFMStrict icons and no folding or custom alerts
var mdGFMStrict = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
			WithFolding(false),
			WithCustomAlerts(false),
		),
	),
)

// Test extension using GFMStrict icons and folding enabled
var mdGFMStrictWithFolding = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
			WithFolding(true),
			WithCustomAlerts(false),
		),
	),
)

// Test extension using GFMStrict icons and custom alerts enabled
var mdGFMStrictWithCustomAlerts = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
			WithFolding(false),
			WithCustomAlerts(true),
		),
	),
)

// Test extension using GFMStrict icons and custom alerts enabled
var mdGFMStrictWithFoldingAndCustomAlerts = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMStrictIcons(),
			WithFolding(true),
			WithCustomAlerts(true),
		),
	),
)



// ###############################################################################################
// STRICT: GFMAlerts with no extensions only supports these primary five (5) alert types
// ###############################################################################################

// TestGFMPlusPrimaryAlerts using the defalt GFMStrict with no extensions
// Using purely GFMStrict with no extra features, these should product functioning alerts
func TestGFMStrictPrimaryAlerts(t *testing.T) {
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

// ###############################################################################################
// STRICT: GFMAlerts with no extensions only support the primary five (5) alert types, so anything
//         else should output as generic blockquotes (i.e., the parser should ignore them and pass
//         controll back to Goldmark)
// ###############################################################################################

// TestGFMStrictUnrecognizedAlerts using the default GFMStrict with no extensions
// Using purely GFMStrict with no extra features, these should produce blockquote output, not alerts
func TestGFMStrictAliasAlerts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Undefined INFO alert.",
			md: `> [!INFO]
> This uses the INFO alias.`,
			html: `<blockquote>
<p>[!INFO]
This uses the INFO alias.</p>
</blockquote>
`,
		},
		{
			desc: "Undefined HINT alert - With Custom Title.",
			md: `> [!HINT] Tip
> Undefined HINT alias with a Tip custom title.`,
			html: `<blockquote>
<p>[!HINT] Tip
Undefined HINT alias with a Tip custom title.</p>
</blockquote>
`,
		},
		{
			desc: "Undefined WARN alert.",
			md: `> [!WARN]
> This uses the undefined WARN alias.`,
			html: `<blockquote>
<p>[!WARN]
This uses the undefined WARN alias.</p>
</blockquote>
`,
		},
		{
			desc: "Undefined ERROR alert.",
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

// TestGFMStrictNoFoldingNoCustomalerts using the default GFMStrict with no extensions
// Using purely GFMStrict with no extra features, these should produce blockquote output, not alerts
func TestGFMStrictNoFoldingNoCustomalerts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Closed by Default folding",
			md: `> [!TIP]-
> This tip callout is closed by default due to the minus sign.`,
			html: `<blockquote>
<p>[!TIP]-
This tip callout is closed by default due to the minus sign.</p>
</blockquote>`,
		},
		{
			desc: "Open by Default folding (Explicit)",
			md: `> [!IMPORTANT]+
> This important callout is explicitly marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!IMPORTANT]+
This important callout is explicitly marked as open by default with the plus sign.</p>
</blockquote>
`,
		},
		{
			desc: "Open by Default folding Custom Alert",
			md: `> [!ZEPHYR]+
> This custom callout is marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!ZEPHYR]+
This custom callout is marked as open by default with the plus sign.</p>
</blockquote>
`,
		},
		{
			desc: "Closed by Default folding Custom Alert with Recognized Title",
			md: `> [!Caution]- Warning
> This custom callout is marked as closed by default with the minus sign.`,
			html: `<blockquote>
<p>[!Caution]- Warning
This custom callout is marked as closed by default with the minus sign.</p>
</blockquote>
`,
		},
		{
			desc: "Open by Default folding Recognized Alert with Recognized Title",
			md: `> [!Caution]+ Warning
> This danger callout is marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!Caution]+ Warning
This danger callout is marked as open by default with the plus sign.</p>
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

// TestGFMStrictCustomTitles using the default GFMStrict with no extensions
// Using purely GFMStrict with no extra features, these should produce blockquote output, not alerts
func TestGFMStrictCustomTitles(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Unrecognized Alert Type - With Custom Title",
			md: `> [!SUCCESS] Mission Accomplished
> Disallowed Undefined SUCCESS alias with custom title.`,
			html: `<blockquote>
<p>[!SUCCESS] Mission Accomplished
Disallowed Undefined SUCCESS alias with custom title.</p>
</blockquote>
`,
		},
		{
			desc: "Unrecognized Alert Type - With Custom Title - HaCkEr StYlE TiTlE",
			md: `> [!SUCCESS] MiSsIoN AcCoMpLiShEd
> Disallowed Undefined SUCCESS alias with custom title.`,
			html: `<blockquote>
<p>[!SUCCESS] MiSsIoN AcCoMpLiShEd
Disallowed Undefined SUCCESS alias with custom title.</p>
</blockquote>
`,
		},
		{
			desc: "Unrecognized NOICON Alert - With Custom Title",
			md: `> [!NoIcon] FooBar
> Disallowed NOICON Callout with Unrecognized Custom Title`,
			html: `<blockquote>
<p>[!NoIcon] FooBar
Disallowed NOICON Callout with Unrecognized Custom Title</p>
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

// ###############################################################################################
// EXTENDED: GFMAlerts using Folding and/or CustomAlerts is available if needed, so we test these.
// ###############################################################################################

// TestGFMStrictWithFolding using the default GFMStrict with folding enabled
// Using GFMStrict with folding enabled, this should allow folding on the five recognized alert types
// but produce blockquotes for unrecognized alerts types or any alert type with a custom title
func TestGFMStrictWithFolding(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Recognized Alert Type - Folding CLOSED by default",
			md: `> [!Tip]-
> This tip callout is closed by default due to the minus sign.`,
			html: `<details class="callout callout-foldable callout-tip iconset-gfm" data-callout="tip"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-lightbulb-icon lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg><p class="callout-title-text">Tip</p>
</summary>
<div class="callout-body"><p>This tip callout is closed by default due to the minus sign.</p>
</div>
</details>`,
		},
		{
			desc: "Recognized Alert Type - Folding OPEN by default",
			md: `> [!IMPORTANT]+
> This important callout is explicitly marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-important iconset-gfm" data-callout="important" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is explicitly marked as open by default with the plus sign.</p>
</div>
</details>`,
		},
		{
			desc: "Unrecognized Alert Type - Folding OPEN by default",
			md: `> [!ZEPHYR]+
> This custom callout is marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!ZEPHYR]+
This custom callout is marked as open by default with the plus sign.</p>
</blockquote>
`,
		},
		{
			desc: "Unrecognized Alert Type - Folding OPEN by default - With Custom Title",
			md: `> [!ZEPHYR]+ Warning
> This custom callout is marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!ZEPHYR]+ Warning
This custom callout is marked as open by default with the plus sign.</p>
</blockquote>
`,
		},
		{
			desc: "Recognized Alert Type - Folding OPEN by default - With Custom Title",
			md: `> [!Caution]+ Warning
> This danger callout is marked as open by default with the plus sign.`,
			html: `<blockquote>
<p>[!Caution]+ Warning
This danger callout is marked as open by default with the plus sign.</p>
</blockquote>
`,
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

// TestGFMStrictWithCustomalerts using the default GFMStrict with custom alerts enabled
// Using GFMStrict with custom alerts enabled, this should allow custom alert types
// but produce blockquotes for any alert with folding symbols `+` or `-`
func TestGFMStrictWithCustomalerts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Custom Alert Type",
			md: `> [!Foo]
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Foo</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Custom Title",
			md: `> [!Foo] BarBaz
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Multi-word Custom Title",
			md: `> [!Foo] BarBaz FooBar
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz FooBar</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Multi-word Custom Title - With Markdown Formatting",
			md: `> [!Foo] BarBaz **FooBar** BingBong
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz <strong>FooBar</strong> BingBong</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With dash in name",
			md: `> [!Bar-Baz]
> Custom Alert Type.`,
			html: `<div class="callout callout-bar-baz iconset-gfm" data-callout="bar-baz"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Bar-Baz</p>
</div>
<div class="callout-body"><p>Custom Alert Type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With underscore in name",
			md: `> [!Bar_Baz]
> Custom Alert Type.`,
			html: `<div class="callout callout-bar_baz iconset-gfm" data-callout="bar_baz"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Bar_baz</p>
</div>
<div class="callout-body"><p>Custom Alert Type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - Folding Ignored",
			md: `> [!ZEPHYR]+
> This custom callout is marked as open by default with the plus sign.`,
			html: `<div class="callout callout-zephyr iconset-gfm" data-callout="zephyr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Zephyr</p>
</div>
<div class="callout-body"><p>This custom callout is marked as open by default with the plus sign.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type with Custom Title - Folding Ignored",
			md: `> [!ZEPHYR]- Warning
> This custom callout is marked as closed by default with the minus sign.`,
			html: `<div class="callout callout-zephyr iconset-gfm" data-callout="zephyr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This custom callout is marked as closed by default with the minus sign.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Unicode in name",
			md: `> [!你好]
> Unicode Alert.`,
			html: `<div class="callout callout-你好 iconset-gfm" data-callout="你好"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">你好</p>
</div>
<div class="callout-body"><p>Unicode Alert.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Unicode in name and title",
			md: `> [!你好] 世界
> Unicode Alert.`,
			html: `<div class="callout callout-你好 iconset-gfm" data-callout="你好"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">世界</p>
</div>
<div class="callout-body"><p>Unicode Alert.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrictWithCustomAlerts, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMStrictWithCustomalerts using the default GFMStrict with custom alerts enabled
// Using GFMStrict with custom alerts enabled, this should allow custom alert types
// but produce blockquotes for any alert with folding symbols `+` or `-`
func TestGFMStrictWithFoldingAndCustomalerts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Custom Alert Type",
			md: `> [!Foo]
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Foo</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Custom Title",
			md: `> [!Foo] BarBaz
> Custom alert type.`,
			html: `<div class="callout callout-foo iconset-gfm" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz</p>
</div>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Alert Type - With Multi-word Custom Title - Folding OPEN by default",
			md: `> [!Foo]+ BarBaz FooBar
> Custom alert type.`,
			html: `<details class="callout callout-foldable callout-foo iconset-gfm" data-callout="foo" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz FooBar</p>
</summary>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</details>`,
		},
		{
			desc: "Custom Alert Type - With Multi-word Custom Title - Folding CLOSED by default",
			md: `> [!Foo]- BarBaz BingBong
> Custom alert type.`,
			html: `<details class="callout callout-foldable callout-foo iconset-gfm" data-callout="foo"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">BarBaz BingBong</p>
</summary>
<div class="callout-body"><p>Custom alert type.</p>
</div>
</details>`,
		},
		{
			desc: "Custom Alert Type - With Unicode in name - Folding OPEN by default",
			md: `> [!你好]+
> Unicode Alert.`,
			html: `<details class="callout callout-foldable callout-你好 iconset-gfm" data-callout="你好" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">你好</p>
</summary>
<div class="callout-body"><p>Unicode Alert.</p>
</div>
</details>`,
		},
		{
			desc: "Custom Alert Type - With Unicode in name and title - Folding CLOSED by default",
			md: `> [!你好]- 世界
> Unicode Alert.`,
			html: `<details class="callout callout-foldable callout-你好 iconset-gfm" data-callout="你好"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">世界</p>
</summary>
<div class="callout-body"><p>Unicode Alert.</p>
</div>
</details>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMStrictWithFoldingAndCustomAlerts, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}
