# goldmark-callout

goldmark-callout is an extension for the
[goldmark](http://github.com/yuin/goldmark) that allows you to use [GitHub
alerts](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts)
or [Obsidian
callouts](https://help.obsidian.md/Editing+and+formatting/Callouts).

## State of the project

Under active development, many breaking changes all the time.

## Example

### Basic example

**Markdown**

```markdown
> [!info] Great new feature
> With lots of possibilities:
> - feature one
> - feature two
```

**HTML**

```html
<details class="callout" data-callout="info" open>
<summary>
<p>Great new feature</p>
</summary>
<p>With lots of possibilities:</p>
<ul>
<li>feature one</li>
<li>feature two</li>
</ul>
</details>
```

### Closed callout

**Markdown**

```markdown
> [!info]- The dash after the callout type makes it closed
> Which is useful for hiding details behind a dropdown
```

**HTML**

```html
<details class="callout" data-callout="info">
<summary>
<p>The dash after the callout type makes it closed</p>
</summary>
<p>Which is useful for hiding details behind a dropdown</p>
</details>
```

## Differences with Obsidian and GitHub syntax

Obsidian and GitHub render callouts with a `div` elements. Obsidian uses
JavaScript for opening and closing of summaries. I decided to use tags `details`
and `summary` to allow opening and closing the the tags without JavaScript.

Obsidian allows arbitrary content in the callout title, including blockquotes,
images, links, and another callouts. We don't allow this, callout title can only
be a single line of paragraph.



## Goals

- [x] Basic syntax
- [x] Markdown in callout titles
- [x] Arbitrary markdown in callout body
- [x] Commonmark-compliant parsing of blockquotes
- [ ] Parity with Obsidian
  - [x] Required space (`> [!info]title` should not be valid)
  - [ ] Default titles (capitalized type)
  - [x] Nested callouts (requires whitespace before `>` to pass)
  - [ ] Use `data-callout` to convey callout type instead of a class
- [ ] Tests (and render) of forced open, closed by default callouts
- [ ] Better README explaining the features
  - [ ] Closed-by-default and forced-open callouts
  - [ ] How to style callouts and add icon with CSS
  - [x] Explain differences between Obsidian implementation and this
- [ ] Test suite from goldmark to make sure we didn't screw anything up
- [ ] Integration test with Goldmark

## License

```
MIT License

Copyright (c) 2024 Adam Chovanec

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
