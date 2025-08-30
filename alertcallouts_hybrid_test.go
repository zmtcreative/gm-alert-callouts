package alertcallouts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Test extension using Hybrid icons, folding, custom alerts, and 'noicon' feature
var mdHybrid = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseHybridIcons(),
			WithFolding(true),
			WithCustomAlerts(true),
			WithIcon("noicon", `<span class="callout-title-noicon" style="display: none;"></span>`),
			// WithAllowNOICON(true),
		),
	),
)

// Test extension using Hybrid icons and folding enabled
var mdHybridWithDefaultIcon = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseHybridIcons(),
			WithFolding(true),
			WithCustomAlerts(true),
			// WithAllowNOICON(true),
			WithIcon("noicon", `<span class="callout-title-noicon" style="display: none;"></span>`),
			WithIcon("default", `<svg class="default-icon"></svg>`),
		),
	),
)

// TestHybridPrimaryCallouts tests the primary callouts from the GFM Plus icon set
func TestHybridPrimaryCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Primary Note callout",
			md: `> [!NOTE]
> This is a note callout with informational content.`,
			html: `<div class="callout callout-note iconset-hybrid" data-callout="note"><div class="callout-title">
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
			html: `<div class="callout callout-tip iconset-hybrid" data-callout="tip"><div class="callout-title">
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
			html: `<div class="callout callout-important iconset-hybrid" data-callout="important"><div class="callout-title">
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
			html: `<div class="callout callout-warning iconset-hybrid" data-callout="warning"><div class="callout-title">
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
			html: `<div class="callout callout-caution iconset-hybrid" data-callout="caution"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-octagon-alert-icon lucide-octagon-alert"><path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/></svg><p class="callout-title-text">Caution</p>
</div>
<div class="callout-body"><p>This is a caution callout for dangerous situations.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Bug Callout",
			md: `> [!bug]
> This is a Bug callout.`,
			html: `<div class="callout callout-bug iconset-hybrid" data-callout="bug"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-bug-icon lucide-bug"><path d="m8 2 1.88 1.88"/><path d="M14.12 3.88 16 2"/><path d="M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"/><path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6"/><path d="M12 20v-9"/><path d="M6.53 9C4.6 8.8 3 7.1 3 5"/><path d="M6 13H2"/><path d="M3 21c0-2.1 1.7-3.9 3.8-4"/><path d="M20.97 5c0 2.1-1.6 3.8-3.5 4"/><path d="M22 13h-4"/><path d="M17.2 17c2.1.1 3.8 1.9 3.8 4"/></svg><p class="callout-title-text">Bug</p>
</div>
<div class="callout-body"><p>This is a Bug callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Example Callout",
			md: `> [!Example]
> This is a Example callout.`,
			html: `<div class="callout callout-example iconset-hybrid" data-callout="example"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-notebook-pen-icon lucide-notebook-pen"><path d="M13.4 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7.4"/><path d="M2 6h4"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><path d="M21.378 5.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/></svg><p class="callout-title-text">Example</p>
</div>
<div class="callout-body"><p>This is a Example callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Failure Callout",
			md: `> [!FaIlUrE]
> This is a Failure callout.`,
			html: `<div class="callout callout-failure iconset-hybrid" data-callout="failure"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-x-icon lucide-circle-x"><circle cx="12" cy="12" r="10"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/></svg><p class="callout-title-text">Failure</p>
</div>
<div class="callout-body"><p>This is a Failure callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Question Callout",
			md: `> [!Question]
> This is a Question callout.`,
			html: `<div class="callout callout-question iconset-hybrid" data-callout="question"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-circle-question-icon lucide-message-circle-question"><path d="M7.9 20A9 9 0 1 0 4 16.1L2 22Z"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/></svg><p class="callout-title-text">Question</p>
</div>
<div class="callout-body"><p>This is a Question callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Quote Callout",
			md: `> [!quote]
> This is a Quote callout.`,
			html: `<div class="callout callout-quote iconset-hybrid" data-callout="quote"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-quote-icon lucide-quote"><path d="M16 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"/><path d="M5 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"/></svg><p class="callout-title-text">Quote</p>
</div>
<div class="callout-body"><p>This is a Quote callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Scroll Callout",
			md: `> [!sRoLL]
> This is a Scroll callout.`,
			html: `<div class="callout callout-sroll iconset-hybrid" data-callout="sroll"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg><p class="callout-title-text">Sroll</p>
</div>
<div class="callout-body"><p>This is a Scroll callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Success Callout",
			md: `> [!Success]
> This is a caution callout.`,
			html: `<div class="callout callout-success iconset-hybrid" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-check-big-icon lucide-circle-check-big"><path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/></svg><p class="callout-title-text">Success</p>
</div>
<div class="callout-body"><p>This is a caution callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary Summary Callout",
			md: `> [!SUMMARy]
> This is a Summary callout.`,
			html: `<div class="callout callout-summary iconset-hybrid" data-callout="summary"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-clipboard-list-icon lucide-clipboard-list"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M12 11h4"/><path d="M12 16h4"/><path d="M8 11h.01"/><path d="M8 16h.01"/></svg><p class="callout-title-text">Summary</p>
</div>
<div class="callout-body"><p>This is a Summary callout.</p>
</div>
</div>`,
		},
		{
			desc: "Primary ToDo Callout",
			md: `> [!TODO]
> This is a TODO callout.`,
			html: `<div class="callout callout-todo iconset-hybrid" data-callout="todo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-list-todo-icon lucide-list-todo"><rect x="3" y="5" width="6" height="6" rx="1"/><path d="m3 17 2 2 4-4"/><path d="M13 6h8"/><path d="M13 12h8"/><path d="M13 18h8"/></svg><p class="callout-title-text">Todo</p>
</div>
<div class="callout-body"><p>This is a TODO callout.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdHybrid, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestHybridAliasCallouts tests alias callouts that reference primary callouts
func TestHybridAliasCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Info alias for Note",
			md: `> [!INFO]
> This uses the "info" alias but renders as a note callout.`,
			html: `<div class="callout callout-info iconset-hybrid" data-callout="info"><div class="callout-title">
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
			html: `<div class="callout callout-hint iconset-hybrid" data-callout="hint"><div class="callout-title">
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
			html: `<div class="callout callout-warn iconset-hybrid" data-callout="warn"><div class="callout-title">
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
			html: `<div class="callout callout-error iconset-hybrid" data-callout="error"><div class="callout-title">
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
			html: `<div class="callout callout-check iconset-hybrid" data-callout="check"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-check-big-icon lucide-circle-check-big"><path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/></svg><p class="callout-title-text">Check</p>
</div>
<div class="callout-body"><p>This uses the &quot;check&quot; alias but renders as a success callout.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdHybrid, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestHybridFoldingExamples tests the folding functionality with GFM Plus callouts
func TestHybridFoldingExamples(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Closed by Default folding",
			md: `> [!TIP]-
> This tip callout is closed by default due to the minus sign.`,
			html: `<details class="callout callout-foldable callout-tip iconset-hybrid" data-callout="tip"><summary class="callout-title">
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
			html: `<details class="callout callout-foldable callout-important iconset-hybrid" data-callout="important" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-message-square-warning-icon lucide-message-square-warning"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is explicitly marked as open by default with the plus sign.</p>
</div>
</details>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdHybrid, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestHybridCustomTitles tests custom titles functionality with GFM Plus callouts
func TestHybridCustomTitles(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Use Existing Icon with Custom Title",
			md: `> [!SUCCESS] Mission Accomplished
> You can override the default title with any custom text.`,
			html: `<div class="callout callout-success iconset-hybrid" data-callout="success"><div class="callout-title">
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
			html: `<div class="callout callout-success iconset-hybrid" data-callout="success"><div class="callout-title">
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
			html: `<div class="callout callout-foo iconset-hybrid" data-callout="foo"><div class="callout-title">
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
			html: `<div class="callout callout-tldr iconset-hybrid" data-callout="tldr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-scroll-text-icon lucide-scroll-text"><path d="M15 12h-5"/><path d="M15 8h-5"/><path d="M19 17V5a2 2 0 0 0-2-2H4"/><path d="M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"/></svg><p class="callout-title-text">tl;dr</p>
</div>
<div class="callout-body"><p>This uses the &quot;tldr&quot; alias as before but uses the custom title of 'tl;dr' instead.</p>
<p>(<em>see <a href="#custom-titles">Custom Titles</a> for more examples of custom titles</em>)</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix with Recognized Callout - with No Title",
			md: `> [!noicon_Tip]
> NoIcon prefix with Recognized Callout - with No Title.`,
			html: `<div class="callout callout-tip iconset-hybrid" data-callout="tip"><div class="callout-title">
<span class="callout-title-noicon" style="display: none;"></span><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><p>NoIcon prefix with Recognized Callout - with No Title.</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix with Custom Callout - with Custom Title",
			md: `> [!NoIcon-FooBar] Bar Baz Bing
> NoIcon prefix with Custom Callout with Custom Title.`,
			html: `<div class="callout callout-foobar iconset-hybrid" data-callout="foobar"><div class="callout-title">
<span class="callout-title-noicon" style="display: none;"></span><p class="callout-title-text">Bar Baz Bing</p>
</div>
<div class="callout-body"><p>NoIcon prefix with Custom Callout with Custom Title.</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix with NO Callout - Disallowed",
			md: `> [!NoIcon-] Warning
> NoIcon prefix with NO Callout - disallowed - will generate a blockquote`,
			html: `<blockquote>
<p>[!NoIcon-] Warning
NoIcon prefix with NO Callout - disallowed - will generate a blockquote</p>
</blockquote>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdHybrid, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

func TestHybridCustomTitlesWithDefaultIcon(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Unknown Callout",
			md: `> [!FOO]
> You can use an unrecognized entry for the callout.`,
			html: `<div class="callout callout-foo iconset-hybrid" data-callout="foo"><div class="callout-title">
<svg class="default-icon"></svg><p class="callout-title-text">Foo</p>
</div>
<div class="callout-body"><p>You can use an unrecognized entry for the callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Custom Callout with Custom Title",
			md: `> [!BAR] Custom Title
> You can use an unrecognized entry for the callout.`,
			html: `<div class="callout callout-bar iconset-hybrid" data-callout="bar"><div class="callout-title">
<svg class="default-icon"></svg><p class="callout-title-text">Custom Title</p>
</div>
<div class="callout-body"><p>You can use an unrecognized entry for the callout.</p>
</div>
</div>
`,
		},
		{
			desc: "NOICON Callout with Custom Title",
			md: `> [!NOICON] Warning
> NOICON Callout with Custom Title.`,
			html: `<div class="callout callout-noicon iconset-hybrid" data-callout="noicon"><div class="callout-title">
<span class="callout-title-noicon" style="display: none;"></span><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>NOICON Callout with Custom Title.</p>
</div>
</div>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdHybridWithDefaultIcon, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}
