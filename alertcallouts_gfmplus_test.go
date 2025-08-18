package alertcallouts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Test extension using GFMPlus icons and folding enabled
var mdGFMPlus = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseGFMPlusIcons(),
			WithFolding(true),
		),
	),
)

// TestGFMPlusPrimaryCallouts tests the primary callouts from the GFM Plus icon set
func TestGFMPlusPrimaryCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Primary Note callout",
			md: `> [!NOTE]
> This is a note callout with informational content.`,
			html: `<div class="callout callout-note iconset-gfmplus" data-callout="note"><div class="callout-title">
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
			html: `<div class="callout callout-tip iconset-gfmplus" data-callout="tip"><div class="callout-title">
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
			html: `<div class="callout callout-important iconset-gfmplus" data-callout="important"><div class="callout-title">
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
			html: `<div class="callout callout-warning iconset-gfmplus" data-callout="warning"><div class="callout-title">
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
			html: `<div class="callout callout-caution iconset-gfmplus" data-callout="caution"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-octagon-alert-icon lucide-octagon-alert"><path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/></svg><p class="callout-title-text">Caution</p>
</div>
<div class="callout-body"><p>This is a caution callout for dangerous situations.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMPlus, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMPlusAliasCallouts tests alias callouts that reference primary callouts
func TestGFMPlusAliasCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Info alias for Note",
			md: `> [!INFO]
> This uses the "info" alias but renders as a note callout.`,
			html: `<div class="callout callout-info iconset-gfmplus" data-callout="info"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>This uses the &quot;info&quot; alias but renders as a note callout.</p>
</div>
</div>`,
		},
		{
			desc: "Hint alias for Tip",
			md: `> [!HINT]
> This uses the "hint" alias but renders as a tip callout.`,
			html: `<div class="callout callout-hint iconset-gfmplus" data-callout="hint"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-lightbulb-icon lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg><p class="callout-title-text">Hint</p>
</div>
<div class="callout-body"><p>This uses the &quot;hint&quot; alias but renders as a tip callout.</p>
</div>
</div>`,
		},
		{
			desc: "Warn alias for Warning",
			md: `> [!WARN]
> This uses the "warn" alias but renders as a warning callout.`,
			html: `<div class="callout callout-warn iconset-gfmplus" data-callout="warn"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-triangle-alert-icon lucide-triangle-alert"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg><p class="callout-title-text">Warn</p>
</div>
<div class="callout-body"><p>This uses the &quot;warn&quot; alias but renders as a warning callout.</p>
</div>
</div>`,
		},
		{
			desc: "Error alias for Caution",
			md: `> [!ERROR]
> This uses the "error" alias but renders as a caution callout.`,
			html: `<div class="callout callout-error iconset-gfmplus" data-callout="error"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-octagon-alert-icon lucide-octagon-alert"><path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/></svg><p class="callout-title-text">Error</p>
</div>
<div class="callout-body"><p>This uses the &quot;error&quot; alias but renders as a caution callout.</p>
</div>
</div>`,
		},
		{
			desc: "Check alias for Success",
			md: `> [!CHECK]
> This uses the "check" alias but renders as a success callout.`,
			html: `<div class="callout callout-check iconset-gfmplus" data-callout="check"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-check-big-icon lucide-circle-check-big"><path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/></svg><p class="callout-title-text">Check</p>
</div>
<div class="callout-body"><p>This uses the &quot;check&quot; alias but renders as a success callout.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMPlus, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMPlusFoldingExamples tests the folding functionality with GFM Plus callouts
func TestGFMPlusFoldingExamples(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Closed by Default folding",
			md: `> [!TIP]-
> This tip callout is closed by default due to the minus sign.`,
			html: `<details class="callout callout-foldable callout-tip iconset-gfmplus" data-callout="tip"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-lightbulb-icon lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg><p class="callout-title-text">Tip</p>
</summary>
<div class="callout-body"><p>This tip callout is closed by default due to the minus sign.</p>
</div>
</details>`,
		},
		{
			desc: "Open by Default folding (Explicit)",
			md: `> [!IMPORTANT]+
> This important callout is explicitly marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-important iconset-gfmplus" data-callout="important" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is explicitly marked as open by default with the plus sign.</p>
</div>
</details>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMPlus, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestGFMPlusCustomTitles tests custom titles functionality with GFM Plus callouts
func TestGFMPlusCustomTitles(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Use Existing Icon with Custom Title",
			md: `> [!SUCCESS] Mission Accomplished
> You can override the default title with any custom text.`,
			html: `<div class="callout callout-success iconset-gfmplus" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-check-big-icon lucide-circle-check-big"><path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/></svg><p class="callout-title-text">Mission Accomplished</p>
</div>
<div class="callout-body"><p>You can override the default title with any custom text.</p>
</div>
</div>`,
		},
		{
			desc: "Custom Title rendered As-Is",
			md: `> [!SUCCESS] MiSsIoN AcCoMpLiShEd
> You can override the default title with any custom text.`,
			html: `<div class="callout callout-success iconset-gfmplus" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-check-big-icon lucide-circle-check-big"><path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/></svg><p class="callout-title-text">MiSsIoN AcCoMpLiShEd</p>
</div>
<div class="callout-body"><p>You can override the default title with any custom text.</p>
</div>
</div>`,
		},
		{
			desc: "Unknown Callout",
			md: `> [!FOO]
> You can use an unrecognized entry for the callout.`,
			html: `<div class="callout callout-foo iconset-gfmplus" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Foo</p>
</div>
<div class="callout-body"><p>You can use an unrecognized entry for the callout.</p>
</div>
</div>`,
		},
		{
			desc: "TLDR alias with custom title",
			md: `> [!TLDR] tl;dr
> This uses the "tldr" alias as before but uses the custom title of 'tl;dr' instead.
>
> (*see [Custom Titles](#custom-titles) for more examples of custom titles*)`,
			html: `<div class="callout callout-tldr iconset-gfmplus" data-callout="tldr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-scroll-text-icon lucide-scroll-text"><path d="M15 12h-5"/><path d="M15 8h-5"/><path d="M19 17V5a2 2 0 0 0-2-2H4"/><path d="M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"/></svg><p class="callout-title-text">tl;dr</p>
</div>
<div class="callout-body"><p>This uses the &quot;tldr&quot; alias as before but uses the custom title of 'tl;dr' instead.</p>
<p>(<em>see <a href="#custom-titles">Custom Titles</a> for more examples of custom titles</em>)</p>
</div>
</div>`,
		},
		{
			desc: "No Icon Callout with No Title",
			md: `> [!NoIcon]
> This creates a No Icon Callout with No Title, but will be styled using default styling.`,
			html: `<div class="callout callout-default iconset-gfmplus" data-callout="default"><div class="callout-title">
<p class="callout-title-text">Noicon</p>
</div>
<div class="callout-body"><p>This creates a No Icon Callout with No Title, but will be styled using default styling.</p>
</div>
</div>`,
		},
		{
			desc: "No Icon Callout with Unrecognized Custom Title",
			md: `> [!NoIcon] FooBar
> This creates a No Icon Callout with a custom title, but will be styled using the default styling`,
			html: `<div class="callout callout-foobar iconset-gfmplus" data-callout="foobar"><div class="callout-title">
<p class="callout-title-text">FooBar</p>
</div>
<div class="callout-body"><p>This creates a No Icon Callout with a custom title, but will be styled using the default styling</p>
</div>
</div>`,
		},
		{
			desc: "No Icon Callout with Recognized Custom Title",
			md: `> [!NoIcon] Warning
> This creates a Warning Callout without the Warning Icon, but will be styled using ` + "`" + `data-callout="warning"` + "`" + `
> rather than the default styling, because 'warning' is a defined callout name.`,
			html: `<div class="callout callout-warning iconset-gfmplus" data-callout="warning"><div class="callout-title">
<p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This creates a Warning Callout without the Warning Icon, but will be styled using <code>data-callout=&quot;warning&quot;</code>
rather than the default styling, because 'warning' is a defined callout name.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdGFMPlus, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}
