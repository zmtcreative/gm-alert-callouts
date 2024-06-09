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
This is a callout
</summary>
<div class="callout-content">
</div>
</details>
`},
	{
		desc: "Callout with a paragraph",
		md: `> [!info] This is a callout
> Paragraph
> over a few lines`,
		html: `<details data-callout="info" open>
<summary>
This is a callout
</summary>
<div class="callout-content">
<p>Paragraph
over a few lines</p>
</div>
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
This is a callout
</summary>
<div class="callout-content">
<p>paragraph
over a few lines</p>
<p>second paragraph with <em>some</em> syntax</p>
</div>
</details>
`},
	{ /* This could be improved so strip out the empty par */
		desc: "Callout without tilte and body",
		md:   `> [!info]`,

		html: `<details data-callout="info" open>
<summary>
Info
</summary>
<div class="callout-content">
</div>
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
title
</summary>
<div class="callout-content">
</div>
</details>
`},
	{
		desc: "Callout with list",
		md: `> [!info] title
> - item`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
<ul>
<li>item</li>
</ul>
</div>
</details>
`},
	{
		desc: "Callout without space before type",
		md:   `>[!info] title`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
</div>
</details>
`},
	{
		desc: "Nested callout",
		md: `>[!info] title
> > [!alert] does this work
> > oh yeah it does`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
<details data-callout="alert" open>
<summary>
does this work
</summary>
<div class="callout-content">
<p>oh yeah it does</p>
</div>
</details>
</div>
</details>`},
	{
		desc: "README example",
		md: `> [!info] Great new feature
> With lots of possibilities:
> - feature one
> - feature two`,
		html: `<details data-callout="info" open>
<summary>
Great new feature
</summary>
<div class="callout-content">
<p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</div>
</details>
`}, {
		desc: "Two nested callouts",
		md: `>[!info] title
> > [!alert] does this work
> > > [!info] Yes it does`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
<details data-callout="alert" open>
<summary>
does this work
</summary>
<div class="callout-content">
<details data-callout="info" open>
<summary>
Yes it does
</summary>
<div class="callout-content">
</div>
</details>
</div>
</details>
</div>
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
title
</summary>
<div class="callout-content">
<details data-callout="alert" open>
<summary>
does this work
</summary>
<div class="callout-content">
<p>text</p>
<p>text</p>
<ul>
<li>list</li>
<li>list</li>
</ul>
</div>
</details>
</div>
</details>`}, {
		desc: "Space before summary",
		md:   `>[!info]  title`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Two spaces before summary",
		md:   `>[!info]   title`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Three spaces before summary",
		md:   `>[!info]    title`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Four spaces before summary",
		md:   `>[!info]     title`,
		html: `<details data-callout="info" open>
<summary>
title
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Syntax in summary",
		md:   `>[!info] Title with *some* syntax [and](http://example.com) links`,
		html: `<details data-callout="info" open>
<summary>
Title with <em>some</em> syntax <a href="http://example.com">and</a> links
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Closed by default callout",
		md:   `>[!info]- I am closed`,
		html: `<details data-callout="info">
<summary>
I am closed
</summary>
<div class="callout-content">
</div>
</details>`}, {
		desc: "Closed by default callout",
		md: `>[!info]- I am closed
> And have some content`,
		html: `<details data-callout="info">
<summary>
I am closed
</summary>
<div class="callout-content">
<p>And have some content</p>
</div>
</details>`}, {
		desc: "README II",
		md: `> [!info]- The dash after the callout type makes it closed
> Which is useful for hiding details behind a dropdown
`,
		html: `<details data-callout="info">
<summary>
The dash after the callout type makes it closed
</summary>
<div class="callout-content">
<p>Which is useful for hiding details behind a dropdown</p>
</div>
</details>`}, {
		desc: "README III",
		md: `> [!warning]
> The callout type with capitalized first letter is used as the the callout
> title
`,
		html: `<details data-callout="warning" open>
<summary>
Warning
</summary>
<div class="callout-content">
<p>The callout type with capitalized first letter is used as the the callout
title</p>
</div>
</details>
`},		{desc: "example",
		md: `> [!info]- The dash after the callout type makes it closed
> Which is useful for hiding details behind a dropdown, especially if there's a lot of them
>
> Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.

> [!attention] Without a dash, callout is opened by default
> For information you want the reasers to see right away
`,
		html: `<details data-callout="info">
<summary>
The dash after the callout type makes it closed
</summary>
<div class="callout-content">
<p>Which is useful for hiding details behind a dropdown, especially if there's a lot of them</p>
<p>Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.</p>
</div>
</details>
<details data-callout="attention" open>
<summary>
Without a dash, callout is opened by default
</summary>
<div class="callout-content">
<p>For information you want the reasers to see right away</p>
</div>
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
