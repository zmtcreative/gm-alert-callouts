# Features and Public API

This document describes the public interface for the `gm-alert-callouts` Goldmark extension, which
provides support for GitHub-style alerts and Obsidian-style callouts.

> [!NOTE]
>
> The project and this document are still under heavy development. It is possible that information
> in this document has not been updated with recent changes to the API for this extension.

## Overview

The `gm-alert-callouts` extension enables rendering of GitHub alerts and Obsidian callouts in Markdown. It supports:

- GitHub-style alerts (`[!NOTE]`, `[!TIP]`, `[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`)
- Obsidian-style callouts with folding functionality (`+` for open, `-` for closed)
- Customizable icon sets (GFM, GFM Plus, Obsidian styles)
- Custom icons via user-defined icon maps
- Nested content support within alert blocks

## Quick Start

### Using the Pre-configured Extension

For basic usage with GFM icons and folding enabled:

```go
import (
    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

md := goldmark.New(
    goldmark.WithExtensions(alertcallouts.AlertCallouts),
)
```

### Using the Customizable Constructor (Recommended)

For more control over configuration:

```go
import (
    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMPlusIcons(),
    alertcallouts.WithFolding(true),
)

md := goldmark.New(
    goldmark.WithExtensions(extension),
)
```

## Public API Reference

### Pre-configured Extension Variable

#### `AlertCallouts`

A pre-configured extension instance with:

- GFM icon set enabled
- Folding functionality enabled
- Ready to use with `goldmark.WithExtensions(alertcallouts.AlertCallouts)`

**Example:**

```go
md := goldmark.New(
    goldmark.WithExtensions(alertcallouts.AlertCallouts),
)
```

### Constructor Function

#### `NewAlertCallouts(options ...Option) *alertCalloutsOptions`

Creates a new AlertCallouts extension with customizable options. This is the recommended approach
for production use as it provides full control over the extension configuration.

**Parameters:**

- `options ...Option`: Variadic functional options to configure the extension

**Returns:**

- `*alertCalloutsOptions`: Configured extension instance that implements `goldmark.Extender`

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseObsidianIcons(),
    alertcallouts.WithFolding(false),
    alertcallouts.WithIcon("custom", "<svg>...</svg>"),
)
```

### Configuration Options

All options are functional options that can be passed to `NewAlertCallouts()`.

#### Icon Configuration Options

##### `WithIcons(icons map[string]string) Option`

Sets a complete custom icon map for alert callouts.

**Parameters:**

- `icons map[string]string`: Map where keys are alert types and values are icon strings (usually SVG)

**Example:**

```go
customIcons := map[string]string{
    "note":      "<svg>...</svg>",
    "important": "<svg>...</svg>",
    "warning":   "<svg>...</svg>",
}

extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithIcons(customIcons),
)
```

##### `WithIcon(kind, icon string) Option`

Adds or overrides a single icon in the icon map.

**Parameters:**

- `kind string`: The alert type identifier
- `icon string`: The icon string (usually SVG content)

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMIcons(),
    alertcallouts.WithIcon("custom", "<svg>custom icon</svg>"),
    alertcallouts.WithIcon("note", "<svg>overridden note icon</svg>"),
)
```

##### `UseGFMIcons() Option`

Configures the extension to use GitHub Flavored Markdown (GFM) standard icons. Includes the five
core GitHub alert types: `note`, `tip`, `important`, `warning`, `caution`, plus common aliases.

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMIcons(),
)
```

**Standard GFM Alert Types:**

- `note` (aliases: `info`, `information`, `notes`)
- `tip` (aliases: `hint`, `hints`, `tips`)
- `important`
- `warning` (aliases: `warn`, `warnings`)
- `caution` (aliases: `danger`, `error`, `cautions`)

##### `UseGFMPlusIcons() Option`

Configures the extension to use an extended icon set that combines GFM and Obsidian-style icons.
This provides the most comprehensive set of built-in icons.

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMPlusIcons(),
)
```

##### `UseObsidianIcons() Option`

Configures the extension to use Obsidian-style icons, which may include additional callout types
common in Obsidian.

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseObsidianIcons(),
)
```

#### Functionality Options

##### `WithFolding(enable bool) Option`

Enables or disables the folding functionality for callouts. When enabled, callouts can be made
collapsible using `+` (default open) or `-` (default closed) after the alert type.

**Parameters:**

- `enable bool`: `true` to enable folding, `false` to disable

**Example:**

```go
// Enable folding
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithFolding(true),
)

// Disable folding
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithFolding(false),
)
```

## Supported Markdown Syntax

### Basic Alert Syntax

```markdown
> [!NOTE]
> This is a note alert.

> [!TIP]
> This is a tip alert.

> [!IMPORTANT]
> This is an important alert.

> [!WARNING]
> This is a warning alert.

> [!CAUTION]
> This is a caution alert.
```

### Foldable Callouts (when folding is enabled)

```markdown
> [!TIP]+
> This callout is open by default.

> [!WARNING]-
> This callout is closed by default.
```

### Multi-line Content

```markdown
> [!NOTE]
> This is the first line of the alert.
>
> This is the second paragraph.
>
> - This is a list item
> - Another list item
```

## HTML Output Structure

The extension generates structured HTML with CSS classes for styling:

```html
<div class="alert alert-note">
  <div class="alert-header">
    <span class="alert-icon"><!-- SVG icon --></span>
    <span class="alert-title">NOTE</span>
    <!-- folding button if folding enabled -->
  </div>
  <div class="alert-body">
    <!-- Alert content -->
  </div>
</div>
```

## Integration Examples

### Basic Integration

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(alertcallouts.AlertCallouts),
    )

    source := `> [!NOTE]
> This is a note!`

    var buf bytes.Buffer
    if err := md.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }

    fmt.Println(buf.String())
}
```

### Advanced Integration with Custom Icons

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    customIcons := map[string]string{
        "note":    `<svg>...</svg>`,
        "custom":  `<svg>custom icon</svg>`,
    }

    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(customIcons),
        alertcallouts.WithFolding(true),
        alertcallouts.WithIcon("special", `<svg>special icon</svg>`),
    )

    md := goldmark.New(
        goldmark.WithExtensions(extension),
    )

    source := `> [!CUSTOM]+
> This uses a custom icon!

> [!SPECIAL]-
> This uses the special icon and is closed by default.`

    var buf bytes.Buffer
    if err := md.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }

    fmt.Println(buf.String())
}
```

### Integration with Other Goldmark Extensions

```go
package main

import (
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(
            extension.GFM,
            extension.Footnote,
            alertcallouts.NewAlertCallouts(
                alertcallouts.UseGFMPlusIcons(),
                alertcallouts.WithFolding(true),
            ),
        ),
    )

    // Use md to convert markdown...
}
```

## Best Practices

1. **Use `NewAlertCallouts()` for production**: The constructor provides better control and is more
   maintainable than the pre-configured variable.
2. **Choose appropriate icon sets**:
   - Use `UseGFMIcons()` for GitHub compatibility
   - Use `UseGFMPlusIcons()` for extended functionality using the default 5 GFM Alert styles plus
     many of the same callout markers from Obsidian
   - Use `UseObsidianIcons()` for Obsidian compatibility using the same icons defined by Obsidian
3. **Icon format**: Icons should be provided as complete SVG strings or other HTML-safe content.
   SVG is recommended for scalability and styling flexibility.
4. **CSS styling**: The generated HTML includes semantic CSS classes. Provide appropriate CSS to
   style the alerts according to your design requirements.
5. **Folding functionality**: Enable folding if you need interactive callouts, disable if you
   prefer static alerts for better performance.

## Notes

- The extension does not include default CSS styling - you must provide your own CSS
- Icons are embedded as-is in the HTML output, so ensure they are properly formatted and safe
- Alert types are case-insensitive in the markdown syntax
- The extension supports nested markdown content within alert blocks
- Custom icon maps completely replace built-in icons unless combined with built-in icon options
