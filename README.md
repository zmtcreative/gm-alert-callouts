# Goldmark Alert Callouts Extension

[![Go Reference](https://pkg.go.dev/badge/github.com/ZMT-Creative/gm-alert-callouts.svg)](https://pkg.go.dev/github.com/ZMT-Creative/gm-alert-callouts)
[![Go version](https://img.shields.io/github/go-mod/go-version/ZMT-Creative/gm-alert-callouts)](https://github.com/ZMT-Creative/gm-alert-callouts)
[![License](https://img.shields.io/github/license/ZMT-Creative/gm-alert-callouts)](./LICENSE.md)

A [Goldmark](https://github.com/yuin/goldmark) extension that provides support for GitHub-style alerts and Obsidian-style callouts with customizable icons and folding functionality.

## Features

- **GitHub Alerts**: Full support for GitHub's five standard alert types (`[!NOTE]`, `[!TIP]`, `[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`)
- **Obsidian Callouts**: Compatible with Obsidian-style callouts including folding (`+` for open, `-` for closed)
- **Customizable Icons**: Built-in icon sets (GFM, GFM Plus, Obsidian) with support for custom icon maps
- **Structured HTML**: Generates semantic HTML with CSS classes for easy styling
- **Nested Content**: Supports complex content including lists, code blocks, and other Markdown elements within alerts

## Installation

```bash
go get github.com/ZMT-Creative/gm-alert-callouts
```

## Quick Start

### Basic Usage

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    // Create Goldmark with the extension using built-in GFM icons
    md := goldmark.New(
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.UseGFMIcons(),
                alertcallouts.WithFolding(true),
            ),
        ),
    )

    source := `# Example
> [!NOTE]
> This is a note with an icon.

> [!WARNING]-
> This is a closed warning callout.`

    var buf bytes.Buffer
    if err := md.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }

    fmt.Println(buf.String())
}
```

### Pre-configured Extension

For convenience, a pre-configured extension is available:

```go
import (
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

// Combining with other extensions
md := goldmark.New(
    goldmark.WithExtensions(
        extension.GFM,                 // GitHub Flavored Markdown
        extension.Footnote,            // Footnotes
        alertcallouts.AlertCallouts,   // Alert callouts
    ),
)
```

This pre-configured extension includes:

- GFM icon set enabled
- Folding functionality enabled
- Ready-to-use configuration

## Supported Markdown Syntax

### Basic Alerts

```markdown
> [!NOTE]
> This is a note alert.

> [!TIP]
> This is a helpful tip.

> [!IMPORTANT]
> This is important information.

> [!WARNING]
> This is a warning message.

> [!CAUTION]
> This is a caution message.
```

### Foldable Callouts

When folding is enabled, use `+` (default open) or `-` (default closed):

```markdown
> [!TIP]+
> This callout is open by default and can be collapsed.

> [!WARNING]-
> This callout is closed by default and can be expanded.
```

### Multi-line Content

```markdown
> [!NOTE]
> This alert contains multiple paragraphs.
>
> - List item one
> - List item two
>
> And even code blocks:
> ```go
> fmt.Println("Hello, World!")
> ```
```

## Configuration Options

The extension supports functional options for flexible configuration:

- **Icon Sets**: `UseGFMIcons()`, `UseGFMPlusIcons()`, `UseObsidianIcons()`
- **Custom Icons**: `WithIcon()`, `WithIcons()`
- **Functionality**: `WithFolding()` (enable/disable collapsible callouts)

For detailed configuration options and examples, see the [API Reference](docs/FEATURES.md#configuration-options).

## HTML Output

The extension generates semantic HTML with CSS classes for styling:

- **Basic alerts**: Use `<div>` elements with `callout` and `callout-{type}` classes
- **Foldable callouts**: Use `<details>` and `<summary>` elements for collapsible functionality
- **Structured content**: Separate containers for title, icon, and content areas

For complete HTML structure details and CSS class reference, see the [API Reference](docs/FEATURES.md#html-output-structure).

## Documentation

- **[API Reference](docs/FEATURES.md)** - Detailed API documentation and usage examples
- **[Icon Customization](docs/ICONMAPS.md)** - Guide to creating custom icons and icon maps

## Requirements

- Go 1.23.0 or later
- [Goldmark](https://github.com/yuin/goldmark) v1.4.6 or later

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
