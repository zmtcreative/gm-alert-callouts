package alerts

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		GhAlertsExtension,
	),
)

type TestCase struct {
	desc string
	md   string
	html string
}

var cases = [...]TestCase{
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
		html: `<div class="markdown-alert markdown-alert-note"><p class="markdown-alert-title"><svg class="octicon octicon-info mr-2" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>Note</p><p>Paragraph
over a few lines</p>
</div>`},
	{
		desc: "Alerts with two paragraphs",
		md: `> [!InFo]
> paragraph
> over a few lines
>
> second paragraph with *some* syntax
`,

		html: `<div class="markdown-alert markdown-alert-info"><p class="markdown-alert-title">Info</p><p>paragraph
over a few lines</p>
<p>second paragraph with <em>some</em> syntax</p>
</div>`},
	{
		desc: "Alerts without body",
		md:   `> [!info] title`,
		html: `<div class="markdown-alert markdown-alert-info"><p class="markdown-alert-title">title</p></div>`},
	{
		desc: "Alerts with list",
		md: `> [!info]
> - item`,
		html: `<div class="markdown-alert markdown-alert-info"><p class="markdown-alert-title">Info</p><ul>
<li>item</li>
</ul>
</div>`},
	{
		desc: "README example",
		md: `> [!info]
> With lots of possibilities:
> - feature one
> - feature two`,
		html: `<div class="markdown-alert markdown-alert-info"><p class="markdown-alert-title">Info</p><p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</div>`},
	{
		desc: "Not a alert",
		md: `[!info] title
`,
		html: `<p>[!info] title</p>
`}, {
		desc: "Syntax in summary",
		md:   `>[!info] Title with *some* syntax [and](http://example.com) links`,
		html: `<div class="markdown-alert markdown-alert-info"><p class="markdown-alert-title">Title with <em>some</em> syntax <a href="http://example.com">and</a> links</p></div>`}, {
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

func TestAlerts(t *testing.T) {
	for i, c := range cases {
		testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}
