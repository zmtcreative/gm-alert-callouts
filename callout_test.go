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
		html: `<details data-callout="info" open>
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
		html: `<details data-callout="info" open>
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

		html: `<details data-callout="info" open>
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

		html: `<details data-callout="info" open>
<summary>
<p></p>
</summary>
</details>
`},
	{
		desc: "No space between type and title",
		md: `> [!info]asdf
`,
		html: `<blockquote>
<p>[!info]asdf</p>
</blockquote>
`},
	{
		desc: "Callout without body",
		md:   `> [!info] title`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>
`},
	{
		desc: "Callout with list",
		md: `> [!info] title
> - item`,
		html: `<details data-callout="info" open>
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
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>
`},
	{
		desc: "Nested callout",
		md: `>[!info] title
> > [!alert] does this work
> > oh yeah it does`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
<details data-callout="alert" open>
<summary>
<p>does this work</p>
</summary>
<p>oh yeah it does</p>
</details>
</details>`},
	{
		desc: "README example",
		md: `> [!info] Great new feature
> With lots of possibilities:
> - feature one
> - feature two`,
		html: `<details data-callout="info" open>
<summary>
<p>Great new feature</p>
</summary>
<p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</details>
`}, {
		desc: "Two nested callouts",
		md: `>[!info] title
> > [!alert] does this work
> > > [!info] Yes it does`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
<details data-callout="alert" open>
<summary>
<p>does this work</p>
</summary>
<details data-callout="info" open>
<summary>
<p>Yes it does</p>
</summary>
</details>
</details>
</details>`},
	{
		desc: "Not a callout",
		md: `[!info] title
`,
		html: `<p>[!info] title</p>
`}, {
		desc: "Two nested callouts",
		md: `>[!info] title
> > [!alert] does this work
> > text
> >
> > text
> > - list
> > - list
> > `,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
<details data-callout="alert" open>
<summary>
<p>does this work</p>
</summary>
<p>text</p>
<p>text</p>
<ul>
<li>list</li>
<li>list</li>
</ul>
</details>
</details>`}, {
		desc: "Space before summary",
		md: `>[!info]  title`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>`},{
		desc: "Two spaces before summary",
		md: `>[!info]   title`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>`},{
		desc: "Three spaces before summary",
		md: `>[!info]    title`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>`},{
		desc: "Four spaces before summary",
		md: `>[!info]     title`,
		html: `<details data-callout="info" open>
<summary>
<p>title</p>
</summary>
</details>`}, {
		desc: "Syntax in summary",
    md: `>[!info] Title with *some* syntax [and](http://example.com) links`,
		html: `<details data-callout="info" open>
<summary>
<p>Title with <em>some</em> syntax <a href="http://example.com">and</a> links</p>
</summary>
</details>`},{
		desc: "Closed by default callout",
    md: `>[!info]- I am closed`,
		html: `<details data-callout="info">
<summary>
<p>I am closed</p>
</summary>
</details>`},{
		desc: "Closed by default callout",
    md: `>[!info]- I am closed
> And have some content`,
		html: `<details data-callout="info">
<summary>
<p>I am closed</p>
</summary>
<p>And have some content</p>
</details>`},{
		desc: "README II",
    md: `> [!info]- The dash after the callout type makes it closed
> Which is useful for hiding details behind a dropdown
`,
		html: `<details data-callout="info">
<summary>
<p>The dash after the callout type makes it closed</p>
</summary>
<p>Which is useful for hiding details behind a dropdown</p>
</details>`},
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
