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

// TestObsidianPrimaryCallouts tests the primary callouts from the GFM Plus icon set
func TestObsidianPrimaryCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Primary Note callout",
			md: `> [!NOTE]
> This is a note callout with informational content.`,
			html: `<div class="callout callout-note iconset-obsidian" data-callout="note"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Note</p>
</div>
<div class="callout-body"><p>This is a note callout with informational content.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Tip callout",
			md: `> [!TIP]
> This is a tip callout with helpful suggestions.`,
			html: `<div class="callout callout-tip iconset-obsidian" data-callout="tip"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Tip</p>
</div>
<div class="callout-body"><p>This is a tip callout with helpful suggestions.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Important callout",
			md: `> [!IMPORTANT]
> This is an important callout highlighting crucial information.`,
			html: `<div class="callout callout-important iconset-obsidian" data-callout="important"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Important</p>
</div>
<div class="callout-body"><p>This is an important callout highlighting crucial information.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Warning callout",
			md: `> [!WARNING]
> This is a warning callout about potential issues.`,
			html: `<div class="callout callout-warning iconset-obsidian" data-callout="warning"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-alert-triangle"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"></path><path d="M12 9v4"></path><path d="M12 17h.01"></path></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This is a warning callout about potential issues.</p>
</div>
</div>
`,
		},
		{
			desc: "Primary Caution callout",
			md: `> [!CAUTION]
> This is a caution callout for dangerous situations.`,
			html: `<div class="callout callout-caution iconset-obsidian" data-callout="caution"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-alert-triangle"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"></path><path d="M12 9v4"></path><path d="M12 17h.01"></path></svg><p class="callout-title-text">Caution</p>
</div>
<div class="callout-body"><p>This is a caution callout for dangerous situations.</p>
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

// TestObsidianAliasCallouts tests alias callouts that reference primary callouts
func TestObsidianAliasCallouts(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Info alias for Note",
			md: `> [!INFO]
> This uses the "info" alias but renders as a note callout.`,
			html: `<div class="callout callout-info iconset-obsidian" data-callout="info"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-info"><circle cx="12" cy="12" r="10"></circle><path d="M12 16v-4"></path><path d="M12 8h.01"></path></svg><p class="callout-title-text">Info</p>
</div>
<div class="callout-body"><p>This uses the &quot;info&quot; alias but renders as a note callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Hint alias for Tip",
			md: `> [!HINT]
> This uses the "hint" alias but renders as a tip callout.`,
			html: `<div class="callout callout-hint iconset-obsidian" data-callout="hint"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-flame"><path d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg><p class="callout-title-text">Hint</p>
</div>
<div class="callout-body"><p>This uses the &quot;hint&quot; alias but renders as a tip callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Warn alias for Warning",
			md: `> [!WARN]
> This uses the "warn" alias but renders as a warning callout.`,
			html: `<div class="callout callout-warn iconset-obsidian" data-callout="warn"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-pencil"><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"></path><path d="m15 5 4 4"></path></svg><p class="callout-title-text">Warn</p>
</div>
<div class="callout-body"><p>This uses the &quot;warn&quot; alias but renders as a warning callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Error alias for Caution",
			md: `> [!ERROR]
> This uses the "error" alias but renders as a caution callout.`,
			html: `<div class="callout callout-error iconset-obsidian" data-callout="error"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-zap"><path d="M4 14a1 1 0 0 1-.78-1.63l9.9-10.2a.5.5 0 0 1 .86.46l-1.92 6.02A1 1 0 0 0 13 10h7a1 1 0 0 1 .78 1.63l-9.9 10.2a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14z"></path></svg><p class="callout-title-text">Error</p>
</div>
<div class="callout-body"><p>This uses the &quot;error&quot; alias but renders as a caution callout.</p>
</div>
</div>
`,
		},
		{
			desc: "Check alias for Success",
			md: `> [!CHECK]
> This uses the "check" alias but renders as a success callout.`,
			html: `<div class="callout callout-check iconset-obsidian" data-callout="check"><div class="callout-title">
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="svg-icon lucide-check"><path d="M20 6 9 17l-5-5"></path></svg><p class="callout-title-text">Check</p>
</div>
<div class="callout-body"><p>This uses the &quot;check&quot; alias but renders as a success callout.</p>
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
			desc: "Open by Default folding (Explicit)",
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
func TestObsidianCustomTitles(t *testing.T) {
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
			desc: "No Icon Callout with No Title",
			md: `> [!NoIcon]
> This creates a No Icon Callout with No Title, but will be styled using default styling.`,
			html: `<div class="callout callout-noicon iconset-obsidian" data-callout="noicon"><div class="callout-title">
<svg></svg><p class="callout-title-text"></p>
</div>
<div class="callout-body"><p>This creates a No Icon Callout with No Title, but will be styled using default styling.</p>
</div>
</div>
`,
		},
		{
			desc: "No Icon Callout with Unrecognized Custom Title",
			md: `> [!NoIcon] FooBar
> This creates a No Icon Callout with a custom title, but will be styled using the default styling`,
			html: `<div class="callout callout-noicon iconset-obsidian" data-callout="noicon"><div class="callout-title">
<svg></svg><p class="callout-title-text">FooBar</p>
</div>
<div class="callout-body"><p>This creates a No Icon Callout with a custom title, but will be styled using the default styling</p>
</div>
</div>
`,
		},
		{
			desc: "No Icon Callout with Recognized Custom Title",
			md: `> [!NoIcon] Warning
> This creates a Warning Callout without the Warning Icon, but will be styled using ` + "`" + `data-callout="warning"` + "`" + `
> rather than the default styling, because 'warning' is a defined callout name.`,
			html: `<div class="callout callout-warning iconset-obsidian" data-callout="warning"><div class="callout-title">
<svg></svg><p class="callout-title-text">Warning</p>
</div>
<div class="callout-body"><p>This creates a Warning Callout without the Warning Icon, but will be styled using <code>data-callout=&quot;warning&quot;</code>
rather than the default styling, because 'warning' is a defined callout name.</p>
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
