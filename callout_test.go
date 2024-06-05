package callout

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

var markdown = goldmark.New(
	goldmark.WithExtensions(
		CalloutExtention,
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
		desc: "Callout",
		md:   "> [!info] This is a callout",
		html: `<details class="obsidian-callout-info">
<summary>
<p>This is a callout</p>
</summary>
</details>
`},
	{
		desc: "Callout with a paragraph",
		md: `> [!info] This is a callout
> Paragraph
> over a few lines`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>This is a callout</p>
</summary>
<p>Paragraph
over a few lines</p>
</details>
`},
	{
		desc: "Callout with two paragraphs",
		md: `> [!info] This is a callout
> paragraph
> over a few lines
>
> second paragraph with *some* syntax
`,

		html: `<details class="obsidian-callout-info">
<summary>
<p>This is a callout</p>
</summary>
<p>paragraph
over a few lines</p>
<p>second paragraph with <em>some</em> syntax</p>
</details>
`},
	{ /* This could be improved so strip out the empty par */
		desc: "Callout without tilte and body",
		md:   `> [!info]`,

		html: `<details class="obsidian-callout-info">
<summary>
<p></p>
</summary>
</details>
`},
	{ /* TODO require space for parity with Obsidian */
		desc: "No space between type and title",
		md: `> [!info]asdf
`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>asdf</p>
</summary>
</details>
`},
	// 		html: `<blockquote>
	// <p>[!info]asdf</p>
	// </blockquote>
	// `},
	{
		desc: "Callout without body",
		md:   `> [!info] title`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>title</p>
</summary>
</details>
`},
	{
		desc: "Callout with list",
		md: `> [!info] title
> - item`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>title</p>
</summary>
<ul>
<li>item</li>
</ul>
</details>
`},
	{
		desc: "Callout without space before type",
		md:   `>[!info] title`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>title</p>
</summary>
</details>
`},
	{
		desc: "Nested callout",
		md: `>[!info] title
> > [!alert] does this work
> > not yet`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>title</p>
</summary>
<details class="obsidian-callout-alert">
<summary>
<p>does this work</p>
</summary>
<p>not yet</p>
</details>
</details>`},
	{
		desc: "README example",
		md: `> [!info] Great new feature
> With lots of possibilities:
> - feature one
> - feature two`,
		html: `<details class="obsidian-callout-info">
<summary>
<p>Great new feature</p>
</summary>
<p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</details>
`},
}

func TestCallout(t *testing.T) {
	for i, c := range cases {
		testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}
