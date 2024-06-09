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
<details data-callout="info" open>
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
```

### Closed callout

**Markdown**

```markdown
> [!info]- The dash after the callout type makes it closed
> Which is useful for hiding details behind a dropdown
```

**HTML**

```html
<details data-callout="info">
<summary>
The dash after the callout type makes it closed
</summary>
<div class="callout-content">
<p>Which is useful for hiding details behind a dropdown</p>
</div>
</details>
```

### Default title

**Markdown**

```markdown
> [!warning]
> The callout type with capitalized first letter is used as the the callout
> title
```

**HTML**

```html
<details data-callout="warning" open>
<summary>
Warning
</summary>
<div class="callout-content">
<p>The callout type with capitalized first letter is used as the the callout
title</p>
</div>
</details>
```

### Styling with CSS

For example, see https://codepen.io/staticnoise/pen/JjqJmmE.

## Differences with Obsidian and GitHub syntax

Obsidian and GitHub render callouts with a `div` elements. Obsidian uses
JavaScript for opening and closing of summaries. I decided to use tags `details`
and `summary` to allow opening and closing the the tags without JavaScript.

Obsidian allows arbitrary content in the callout title, including blockquotes,
images, links, and another callouts. We don't allow this, callout title can only
be a single line of paragraph.

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
