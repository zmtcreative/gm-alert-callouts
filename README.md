# goldmark-gh-alerts

goldmark-gh-alerts is an extension for the
[Goldmark](http://github.com/yuin/goldmark) Markdown Rendering Package that allows you to use [GitHub
alerts](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts).

This is a fork of:
[thiagokokada/goldmark-gh-alerts](https://github.com/thiagokokada/goldmark-gh-alerts).

## Changes from Original Extension

This modified version of the GitHub Alerts extension adds `<div>` wrappers around the alert **Title** text and the alert **Body** text. This allows more detailed styling with CSS. A new `code-examples` folder containing a more detailed usage example has also been added (see [More Detailed Example](#more-detailed-example) below).

## State of the project

If you want to use it in your own project feel free, but I recommend either
pinning a commit or forking since the API is not guarantee to be stable.

## Example

### Basic example

#### **Go**

```go
var markdown = goldmark.New(
  goldmark.WithExtensions(
    &GhAlerts{
      Icons: map[string]string{"note": `<svg class="octicon octicon-info mr-2" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>`},
    },
  ),
)
```

#### **Markdown**

```markdown
> [!NOTE]
> Useful information that users should know, even when skimming content.
```

#### **HTML**

```html
<div class="gh-alert gh-alert-note">
  <div class="gh-alert-title">
    <p>
      <svg class="octicon octicon-info mr-2" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>
      Note
    </p>
  </div>
  <div class="gh-alert-body">
    <p>Useful information that users should know, even when skimming content.</p>
  </div>
</div>
```

### More Detailed Example

A more detailed code example is located in the `code-examples` folder. If you are on Windows you can
run the `run-ghalerts.ps1` script which will generate the HTML output of sample GitHub Alerts markdown
text. This will write the output to `example.html` and then start the default web browser to view it.

If you are on MacOS or Linux, just do the following in the `code-examples` folder:

```sh
go run ./ghalerts.go > example.html
open example.html
```

## License

```text
MIT License

Copyright (c) 2024 Adam Chovanec
Copyright (c) 2025 Thiago Kenji Okada
Copyright (c) 2025 ZMT Creative LLC

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
