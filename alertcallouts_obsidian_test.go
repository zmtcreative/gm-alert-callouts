package alertcallouts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Test extension using Obsidian icons and folding enabled
var mdObsidian = goldmark.New(
	goldmark.WithExtensions(
		NewAlertCallouts(
			UseObsidianIcons(),
			WithFolding(true),
			WithCustomAlerts(true),
			WithAllowNOICON(false),
		),
	),
)

// var mdObsidianWithNOICON = goldmark.New(
// 	goldmark.WithExtensions(
// 		NewAlertCallouts(
// 			UseObsidianIcons(),
// 			WithFolding(true),
// 			WithCustomAlerts(true),
// 			WithAllowNOICON(true),
// 		),
// 	),
// )

// TestObsidianPrimaryCallouts tests the primary callouts from the GFM Plus icon set
func TestObsidianPrimaryCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Primary Note callout",
			md: `> [!NOTE]
> This is a Note callout with informational content.`,
			html: `<div class="callout callout-note iconset-obsidian" data-callout="note"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>This is a Note callout with informational content.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Abstract Callout",
			md: `> [!ABSTRACT]
> This is an Abstract callout.`,
			html: `<div class="callout callout-abstract iconset-obsidian" data-callout="abstract"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-clipboard-list"><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><path d="M12 11h4"></path><path d="M12 16h4"></path><path d="M8 11h.01"></path><path d="M8 16h.01"></path></svg><p class="callout-title-text">Abstract</p>
</div>
<div class="callout-body"><p>This is an Abstract callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Info callout",
			md: `> [!INFO]
> This is an Info callout.`,
			html: `<div class="callout callout-info iconset-obsidian" data-callout="info"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-info"><circle cx="12" cy="12" r="10"></circle><path d="M12 16v-4"></path><path d="M12 8h.01"></path></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>This is an Info callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Todo callout",
			md: `> [!TODO]
> This is a Todo callout.`,
			html: `<div class="callout callout-todo iconset-obsidian" data-callout="todo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check-circle-2"><circle cx="12" cy="12" r="10"></circle><path d="m9 12 2 2 4-4"></path></svg><p class="callout-title-text">Todo</p>
</div>
<div class="callout-body"><p>This is a Todo callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Tip callout",
			md: `> [!TIP]
> This is a Tip callout.`,
			html: `<div class="callout callout-tip iconset-obsidian" data-callout="tip"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><p>This is a Tip callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Success callout",
			md: `> [!SUCCESS]
> This is a Success callout.`,
			html: `<div class="callout callout-success iconset-obsidian" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check"><path d="M20 6 9 17l-5-5"></path></svg><p class="callout-title-text">Success</p>
</div>
<div class="callout-body"><p>This is a Success callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Question callout",
			md: `> [!QUESTION]
> This is a Question callout.`,
			html: `<div class="callout callout-question iconset-obsidian" data-callout="question"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-help-circle"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><path d="M12 17h.01"></path></svg><p class="callout-title-text">Question</p>
</div>
<div class="callout-body"><p>This is a Question callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Warning callout",
			md: `> [!WARNING]
> This is a Warning callout.`,
			html: `<div class="callout callout-warning iconset-obsidian" data-callout="warning"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-alert-triangle"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"></path><path d="M12 9v4"></path><path d="M12 17h.01"></path></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This is a Warning callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Failure callout",
			md: `> [!FAILURE]
> This is a Failure callout.`,
			html: `<div class="callout callout-failure iconset-obsidian" data-callout="failure"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-x"><path d="M18 6 6 18"></path><path d="m6 6 12 12"></path></svg><p class="callout-title-text">Failure</p>
</div>
<div class="callout-body"><p>This is a Failure callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Danger callout",
			md: `> [!DANGER]
> This is a Danger callout.`,
			html: `<div class="callout callout-danger iconset-obsidian" data-callout="danger"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-zap"><path d="M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"></path></svg><p class="callout-title-text">Danger</p>
</div>
<div class="callout-body"><p>This is a Danger callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Bug callout",
			md: `> [!BUG]
> This is a Bug callout.`,
			html: `<div class="callout callout-bug iconset-obsidian" data-callout="bug"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-bug"><path d="m8 2 1.88 1.88"></path><path d="M14.12 3.88 16 2"></path><path d="M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"></path><path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6"></path><path d="M12 20v-9"></path><path d="M6.53 9C4.6 8.8 3 7.1 3 5"></path><path d="M6 13H2"></path><path d="M3 21c0-2.1 1.7-3.9 3.8-4"></path><path d="M20.97 5c0 2.1-1.6 3.8-3.5 4"></path><path d="M22 13h-4"></path><path d="M17.2 17c2.1.1 3.8 1.9 3.8 4"></path></svg><p class="callout-title-text">Bug</p>
</div>
<div class="callout-body"><p>This is a Bug callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Example callout",
			md: `> [!EXAMPLE]
> This is an Example callout.`,
			html: `<div class="callout callout-example iconset-obsidian" data-callout="example"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-list"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg><p class="callout-title-text">Example</p>
</div>
<div class="callout-body"><p>This is an Example callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Quote callout",
			md: `> [!QUOTE]
> This is a Quote callout.`,
			html: `<div class="callout callout-quote iconset-obsidian" data-callout="quote"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-quote"><path d="M16 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"></path><path d="M5 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"></path></svg><p class="callout-title-text">Quote</p>
</div>
<div class="callout-body"><p>This is a Quote callout.</p>
</div>
</div>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdObsidian, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestObsidianAliasCallouts tests some alias callouts that reference primary callouts
func TestObsidianAliasCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Summary callout - Uses Abstract Icon",
			md: `> [!SUMMARY]
> This is a Summary callout (an alias for Abstract).`,
			html: `<div class="callout callout-summary iconset-obsidian" data-callout="summary"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-clipboard-list"><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><path d="M12 11h4"></path><path d="M12 16h4"></path><path d="M8 11h.01"></path><path d="M8 16h.01"></path></svg><p class="callout-title-text">Summary</p>
</div>
<div class="callout-body"><p>This is a Summary callout (an alias for Abstract).</p>
</div>
</div>
`,
		},
		{
			desc: "Hint callout - Uses Tip Icon",
			md: `> [!HINT]
> This is a Hint callout (an alias for Tip).`,
			html: `<div class="callout callout-hint iconset-obsidian" data-callout="hint"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Hint</p>
</div>
<div class="callout-body"><p>This is a Hint callout (an alias for Tip).</p>
</div>
</div>
`,
		},
		{
			desc: "Error callout - Uses Danger Icon",
			md: `> [!ERROR]
> This is an Error callout (an alias for Danger).`,
			html: `<div class="callout callout-error iconset-obsidian" data-callout="error"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-zap"><path d="M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"></path></svg><p class="callout-title-text">Error</p>
</div>
<div class="callout-body"><p>This is an Error callout (an alias for Danger).</p>
</div>
</div>
`,
		},
		{
			desc: "Check callout - Uses Success Icon",
			md: `> [!CHECK]
> This is a Check callout (an alias for Success).`,
			html: `<div class="callout callout-check iconset-obsidian" data-callout="check"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check"><path d="M20 6 9 17l-5-5"></path></svg><p class="callout-title-text">Check</p>
</div>
<div class="callout-body"><p>This is a Check callout (an alias for Success).</p>
</div>
</div>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdObsidian, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestObsidianFoldingExamples tests the folding functionality with GFM Plus callouts
func TestObsidianFoldingExamples(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Closed by Default folding",
			md: `> [!TIP]-
> This tip callout is closed by default due to the minus sign.`,
			html: `<details class="callout callout-foldable callout-tip iconset-obsidian" data-callout="tip"><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Tip</p>
</summary>
<div class="callout-body"><p>This tip callout is closed by default due to the minus sign.</p>
</div>
</details>
`,
		},
		{
			desc: "Open by Default folding",
			md: `> [!IMPORTANT]+
> This important callout is explicitly marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-important iconset-obsidian" data-callout="important" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Important</p>
</summary>
<div class="callout-body"><p>This important callout is explicitly marked as open by default with the plus sign.</p>
</div>
</details>
`,
		},
		{
			desc: "Open by Default folding Custom Alert",
			md: `> [!ZEPHYR]+
> This custom callout is marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-zephyr iconset-obsidian" data-callout="zephyr" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Zephyr</p>
</summary>
<div class="callout-body"><p>This custom callout is marked as open by default with the plus sign.</p>
</div>
</details>
`,
		},
		{
			desc: "Open by Default folding Custom Alert with Custom Title",
			md: `> [!ZEPHYR]+ Warning
> This custom callout is marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-zephyr iconset-obsidian" data-callout="zephyr" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Warning</p>
</summary>
<div class="callout-body"><p>This custom callout is marked as open by default with the plus sign.</p>
</div>
</details>`,
		},
		{
			desc: "Open by Default folding Recognized Alert with Recognized Title",
			md: `> [!Danger]+ Warning
> This danger callout is marked as open by default with the plus sign.`,
			html: `<details class="callout callout-foldable callout-danger iconset-obsidian" data-callout="danger" open><summary class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-zap"><path d="M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"></path></svg><p class="callout-title-text">Warning</p>
</summary>
<div class="callout-body"><p>This danger callout is marked as open by default with the plus sign.</p>
</div>
</details>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdObsidian, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

// TestObsidianCustomTitles tests custom titles functionality with GFM Plus callouts
func TestObsidianCustomAlerts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Use Existing Icon with Custom Title",
			md: `> [!SUCCESS] Mission Accomplished
> You can override the default title with any custom text.`,
			html: `<div class="callout callout-success iconset-obsidian" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check"><path d="M20 6 9 17l-5-5"></path></svg><p class="callout-title-text">Mission Accomplished</p>
</div>
<div class="callout-body"><p>You can override the default title with any custom text.</p>
</div>
</div>
`,
		},
		{
			desc: "Custom Title rendered As-Is",
			md: `> [!SUCCESS] MiSsIoN AcCoMpLiShEd
> You can override the default title with any custom text.`,
			html: `<div class="callout callout-success iconset-obsidian" data-callout="success"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check"><path d="M20 6 9 17l-5-5"></path></svg><p class="callout-title-text">MiSsIoN AcCoMpLiShEd</p>
</div>
<div class="callout-body"><p>You can override the default title with any custom text.</p>
</div>
</div>
`,
		},
		{
			desc: "Custom Callout using note icon",
			md: `> [!FOO]
> You can use an unrecognized entry for the callout.`,
			html: `<div class="callout callout-foo iconset-obsidian" data-callout="foo"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Foo</p>
</div>
<div class="callout-body"><p>You can use an unrecognized entry for the callout.</p>
</div>
</div>
`,
		},
		{
			desc: "TLDR alias with custom title",
			md: `> [!TLDR] tl;dr
> This uses the "tldr" alias as before but uses the custom title of 'tl;dr' instead.
>
> (*see [Custom Titles](#custom-titles) for more examples of custom titles*)`,
			html: `<div class="callout callout-tldr iconset-obsidian" data-callout="tldr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-clipboard-list"><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><path d="M12 11h4"></path><path d="M12 16h4"></path><path d="M8 11h.01"></path><path d="M8 16h.01"></path></svg><p class="callout-title-text">tl;dr</p>
</div>
<div class="callout-body"><p>This uses the &quot;tldr&quot; alias as before but uses the custom title of 'tl;dr' instead.</p>
<p>(<em>see <a href="#custom-titles">Custom Titles</a> for more examples of custom titles</em>)</p>
</div>
</div>
`,
		},
		{
			desc: "NoIcon prefix on Recognized Callout - with No Title",
			md: `> [!NoIcon-Tip]
> NoIcon prefix on Recognized Callout with No Title`,
			html: `<div class="callout callout-tip iconset-obsidian" data-callout="tip"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><p>NoIcon prefix on Recognized Callout with No Title</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix on Recognized Callout - with Custom Title",
			md: `> [!NoIcon-Tip] Special Tip
> NoIcon prefix on Recognized Callout with Custom Title`,
			html: `<div class="callout callout-tip iconset-obsidian" data-callout="tip"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Special Tip</p>
</div>
<div class="callout-body"><p>NoIcon prefix on Recognized Callout with Custom Title</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix on Custom Callout - with No Title",
			md: `> [!NoIcon-Zephyr]
> NoIcon prefix on Custom Callout with No Title`,
			html: `<div class="callout callout-zephyr iconset-obsidian" data-callout="zephyr"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Zephyr</p>
</div>
<div class="callout-body"><p>NoIcon prefix on Custom Callout with No Title</p>
</div>
</div>`,
		},
		{
			desc: "NOICON Callout - with Custom Title (just another custom alert)",
			md: `> [!NoIcon] Warning
> NOICON Callout with Custom Title`,
			html: `<div class="callout callout-noicon iconset-obsidian" data-callout="noicon"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>NOICON Callout with Custom Title</p>
</div>
</div>`,
		},
		{
			desc: "NoIcon prefix with NO CALLOUT - Disallowed - generates a blockquote",
			md: `> [!NoIcon-] Special Tip
> NoIcon prefix with NO CALLOUT - Disallowed - generates a blockquote`,
			html: `<blockquote>
<p>[!NoIcon-] Special Tip
NoIcon prefix with NO CALLOUT - Disallowed - generates a blockquote</p>
</blockquote>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.DoTestCase(mdObsidian, testutil.MarkdownTestCase{
				Description: tc.desc,
				Markdown:    tc.md,
				Expected:    tc.html,
			}, t)
		})
	}
}

