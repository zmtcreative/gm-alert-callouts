# goldmark-callout

goldmark-callout is an extension for the
[goldmark](http://github.com/yuin/goldmark) that allows you to use [GitHub
alerts](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts)
or [Obsidian
callouts](https://help.obsidian.md/Editing+and+formatting/Callouts).

## Example

**Markdown**

```markdown
> [!info] Great new feature
> With lots of possibilities:
> - feature one
> - feature two
```

**HTML**

```html
<details class="obsidian-callout-info">
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

## Goals

- [x] Basic syntax
- [x] Markdown in callout titles
- [x] Arbitrary markdown in callout body
- [ ] Commonmark-compliant parsing of blockquotes
  - Missing check for whispace before `>`
- [ ] Parity with Obsidian
  - [ ] Required space (`> [!info]title` should not be valid)
  - [ ] Default titles (capitalized type)
  - [ ] Nested callouts (requires whitespace before `>` to pass)
- [ ] Tests (and render) of forced open, closed by default callouts
- [ ] Better README explaining the features
  - [ ] Closed-by-default and forced-open callouts
  - [ ] How to style callouts and add icon with CSS
- [ ] Test suite from goldmark to make sure we didn't screw anything up

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
