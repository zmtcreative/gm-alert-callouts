# Goldmark GitHub Alerts and Obsidian Callouts

[![Go Reference](https://pkg.go.dev/badge/github.com/ZMT-Creative/gm-alert-callouts.svg)](https://pkg.go.dev/github.com/ZMT-Creative/gm-alert-callouts)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ZMT-Creative/gm-alert-callouts)
![GitHub License](https://img.shields.io/github/license/ZMT-Creative/gm-alert-callouts)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/ZMT-Creative/gm-alert-callouts)
![GitHub Tag](https://img.shields.io/github/v/tag/ZMT-Creative/gm-alert-callouts?include_prereleases&sort=semver)

The `gm-alert-callouts` package is an extension for the
[Goldmark](http://github.com/yuin/goldmark) Markdown Rendering Package that allows you to use
[GitHub alerts](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts).
It also supports Obsidian-style callouts, including the Open/Close feature using
the `+` (default open) and `-` (default closed) characters appended to the marker after
the `]` (right square bracket).

> [!NOTE]
>
> This extension does **not** include any icons -- it just provides the parsing and
> rendering functionality to create the alerts/callouts. The user must provide the list of valid
> alert/callout names with a mapped string (`map[string]string{}`) containing the alert/callout
> identifier as the key (*e.g., `note`, `important`, etc.*) and the icon as the string value. The
> icon is usually an SVG in HTML format (*e.g., `<svg>...svg-definiton...</svg>`*), but can be any
> string that a browser or application can render (*e.g., a Unicode glyph or an HTML entity code*).

Throughout this document and the code itself, the terms `alert(s)` and `callout(s)` are used
interchangeably. GitHub refers to these as `Alerts` while Obsidian refers to them as `Callouts` --
for the purposes of this extension they mean the same thing.

## State of the Project

This ZMT-Creative project is a hard fork of:
[thiagokokada/goldmark-gh-alerts](https://github.com/thiagokokada/goldmark-gh-alerts).

As stated on Thiago Okada's original project's page, his `goldmark-gh-alerts` extension package has
been created primarily to support another of his projects and is not meant for general usage.

ZMT-Creative's package is also being used for a specific (*currently private*) project to create a
standalone Markdown Reader application.

If you want to use this package in your own project feel free, but it is recommended that you
should either pin a commit or fork since the API is not guarantee to be stable at this time.

## Changes from Original Extension

This modified version of the `goldmark-gh-alerts` extension adds `<div>` wrappers around the alert
**Title** text and the alert **Body** text. This allows more detailed styling with CSS.

A new `examples` folder containing a more detailed usage example has also been added (see [More
Detailed Example](#more-detailed-example-code) below).

By default, the extension also supports Obsidian-style foldable callouts, and uses the `<details>`
and `<summary>` HTML elements to wrap the callout. More details on this feature are explained below
and can be seen in the code example.

## Enabling the Extension in Goldmark

Install the extension:

```sh
go get github.com/ZMT-Creative/gm-alert-callouts
```

Starting with version 0.5.0, the extension supports a more idiomatic Go initialization pattern
using functional options. This provides better extensibility and follows Goldmark conventions:

In your code:

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    // Create extension with functional options
    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcon("note", `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
          viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
          stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><path d="M12 16v-4"></path>
          <path d="M12 8h.01"></path></svg>`),
        alertcallouts.WithIcon("warning", `<svg xmlns="http://www.w3.org/2000/svg" width="24"
          height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
          stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 14.5A2.5 2.5 0 0 0 11
          12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0
          1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"></path></svg>`),
        alertcallouts.WithIcon("tip", `<svg xmlns="http://www.w3.org/2000/svg" width="24"
          height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
          stroke-linecap="round" stroke-linejoin="round"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9
          1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"></path>
          <path d="M9 18h6"></path><path d="M10 22h4"></path></svg>`),
        alertcallouts.WithFolding(true), // Enable folding (default)
    )

    markdown := goldmark.New(goldmark.WithExtensions(extension))

    mdSource := `# Alert Examples
> [!NOTE]
> This is a note with an icon.

> [!WARNING]-
> This is a closed warning callout.

> [!TIP]+
> This is an open tip callout.`

    var buf bytes.Buffer
    if err := markdown.Convert([]byte(mdSource), &buf); err != nil {
        panic(err)
    }

    fmt.Printf(`<html><head></head><body>%s</body></html>`, buf.String())
}
```

### Available Functional Options

| Function | Description |
| :------- | :---------- |
| `WithIcon(kind, icon string)` | Adds a single icon for the specified alert type |
| `WithIcons(icons map[string]string)` | Sets the complete icons map (replaces any existing icons) |
| `WithFolding(enable bool)` | Enables or disables folding functionality (enabled by default) |

### Alternative Icon Configuration

You can also configure multiple icons at once:

```go
icons := map[string]string{
    "note":      "<svg>...</svg>",
    "warning":   "<svg>...</svg>",
    "important": "<svg>...</svg>",
    "tip":       "<svg>...</svg>",
    "caution":   "<svg>...</svg>",
}

extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithIcons(icons),
    alertcallouts.WithFolding(true),
)
```

### Backward Compatibility

The simple initialization method is still fully supported:

```go
// Using the global variable (simplest)
// This will initalize the extension with Folding enabled but no icons
markdown := goldmark.New(goldmark.WithExtensions(alertcallouts.AlertCallouts))

// This is equivalent to the newer init method called with no options
markdown := goldmark.New(goldmark.WithExtension(alertcallouts.NewAlertCallouts()))
```

## Standard Alert/Callout Style

### **Markdown Example 1**

This is a standard GitHub-style Alert (also used for non-folding Obsidian Callouts):

```markdown
> [!IMPORTANT]
> This is a GitHub important alert!
```

#### **HTML for Example 1**

```html
<div class="callout callout-important" data-callout="important">
  <div class="callout-title">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
    class="lucide lucide-message-square-warning-icon lucide-message-square-warning">
    <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/>
    <path d="M12 13h.01"/></svg>
    <p class="callout-title-text">Important</p>
  </div>
  <div class="callout-body">
    <p>This is a GitHub important alert!</p>
  </div>
</div>
```

## Foldable Alert/Callouts

When a `-` or `+` are appended directly after the marker (e.g., `> [!TIP]-`) the extension will
alter the HTML output to use the `<details>` and `<summary>` elements. This allows us to open and close
the callout and style it in CSS. The `-` creates a closed-by-default callout and the `+` creates an
open-by-default callout.

If no symbol is used, the HTML output uses just `<div>` elements (as noted in the first example above).

### Markdown Example 2 (Closed Callout)

This is a **Tip** alert using a foldable callout that is closed by default:

```markdown
> [!TIP]-
> This is a GitHub tip in a closed callout.
```

#### HTML for Example 2

```html
<details class="callout callout-foldable callout-tip" data-callout="tip">
  <summary class="callout-title">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
    class="lucide lucide-lightbulb-icon lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5
    1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/>
    <path d="M10 22h4"/></svg>
    <p class="callout-title-text">Tip</p>
  </summary>
  <div class="callout-body">
    <p>This is a GitHub tip in a closed callout.</p>
  </div>
</details>
```

### Markdown for Example 3 (Open Callout)

This is an **Info** alert using a foldable callout that is open by default:

```markdown
> [!INFO]+
> This is an info alert in a foldable callout (open by default).
```

#### HTML for Example 3

```html
<details class="callout callout-foldable callout-info" data-callout="info" open>
  <summary class="callout-title">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
    class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/>
    <path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
    <p class="callout-title-text">Info</p>
  </summary>
  <div class="callout-body">
    <p>This is an info alert in a foldable callout (open by default).</p>
  </div>
</details>
```

## More Detailed Example Code

A more detailed code example is located in the `examples` folder. If you are on Windows you can run
the `run-ghalerts.ps1` script which will generate the HTML output of sample GitHub Alerts markdown
text. This will write the output to `example.html` and then start the default web browser to view
it.

If you are on MacOS or Linux, you should be able to run the `run-showalerts.sh` bash script. You might
need to run the command `chmod +x run-showalerts.sh` to make the script executable. This script has
**not** been tested since we aren't currently developing any of this on Linux or MacOS.

The example shows one possible way to implement a set of alert types and their icons.

## License

This project is licensed under the [MIT License](LICENSE.md)

## Project Lineage

Portions of this software are based on the work of others, used under their respective MIT
Licenses. In keeping with the requirements of the MIT License, here are the license notices for
these authors:

- [Adam Chovanec](docs/LICENSE-chovanec.md)
- [Thiago Okada](docs/LICENSE-thiagokokada.md)
