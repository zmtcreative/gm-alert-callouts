# API Reference - gm-alert-callouts

[<-back to README](../README.md)

This document provides API documentation for the `gm-alert-callouts` Goldmark extension.

## Overview

The `gm-alert-callouts` extension enables rendering of GitHub-style alerts and Obsidian-style
callouts in Markdown documents.

> [!NOTE]
>
> GitHub uses the term `Alerts` while Obsidian uses the term `Callouts` to refer to markdown
> created using the `> [!NOTE]` syntax. Throughout the documentation (*and even within the code*)
> you will see references to `Alerts` and `Callouts` (and `alertcallouts`) -- these terms should be
> considered interchangeable for our purposes.

## Core Types and Interfaces

### Extension Interface

The extension implements `goldmark.Extender`:

```go
type Extender interface {
    Extend(goldmark.Markdown)
}
```

## Initialization

### Pre-configured Extension

#### `AlertCallouts`

```go
var AlertCallouts = NewAlertCallouts(
    UseGFMStrictIcons(),
    WithFolding(true),
)
```

A ready-to-use extension instance with sensible defaults:

- GFM + Aliases (GitHub Flavored Markdown with some additional aliases) icon set
- Folding functionality enabled

**Example:**

```go
md := goldmark.New(
    goldmark.WithExtensions(alertcallouts.AlertCallouts),
)
```

### Custom Extension Constructor

#### `NewAlertCallouts(options ...Option) *alertCalloutsOptions`

Creates a new extension instance with functional options for full customization control.

**Parameters:**

- `options ...Option`: Variadic functional options

**Returns:**

- `*alertCalloutsOptions`: Configured extension that implements `goldmark.Extender`

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMPlusIcons(),
    alertcallouts.WithFolding(true),
    alertcallouts.WithIcon("custom", "<svg>...</svg>"),
)

md := goldmark.New(goldmark.WithExtensions(extension))
```

## Configuration Options

All configuration uses functional options that can be passed to `NewAlertCallouts()`.

### Icon Configuration

-----

#### `UseGFMStrictIcons() Option`

Configures the extension with GitHub Flavored Markdown standard icons. Folding and Custom Alerts are disabled by default.

**Included Alert Types:**

- `note`
- `tip`
- `important`
- `warning`
- `caution`

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMStrictIcons(),
)
```

-----

#### `UseGFMWithAliasesIcons() Option`

Configures the extension with GitHub Flavored Markdown standard icons with some aliases. Folding and Custom Alerts are disabled by default.

**Included Alert Types:**

- `note` (aliases: `info`, `information`, `notes`)
- `tip` (aliases: `hint`, `hints`, `tips`)
- `important`
- `warning` (aliases: `warn`, `warnings`)
- `caution` (aliases: `danger`, `error`, `cautions`)

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMWithAliasesIcons(),
)
```

-----

#### `UseGFMPlusIcons() Option`

Configures the extension with an extended icon set that combines GitHub and Obsidian-style icons for maximum compatibility. This will use the expected GitHub-style icons for the standard five GitHub Alerts, but will use an extended set of icons for other callouts. **If you require strict adherence to Obsidian-style icons, use the `UseObsidianIcons()` option instead.**

This option also enables Folding, Custom Alerts and allows the use of the `[!NOICON]` alert type to render a callout without an icon. If you use a title after the `[!NOICON]` that is a recognized alert type, the styling for that type will be used.

For example:

```markdown
> [!NOICON] Warning
> This is a warning!
```

Would be styled as a `Warning` callout type, but would not have an icon.

**Included Alert Types:**

- `note` (aliases: `notes`, `info`, `information`)
- `tip` (aliases: `tips`, `hint`, `hints`)
- `important`
- `warning` (aliases: `warn`, `warnings`, `attention`)
- `caution` (aliases: `danger`, `error`, `errors`)
- `bug`
- `example`
- `failure` (aliases: `fail`, `missing`)
- `question` (aliases: `questions`, `faq`, `faqs`, `help`)
- `quote` (aliases: `quotes`, `cite`, `citation`, `citations`)
- `scroll` (aliases: `history`, `tldr`)
- `success` (aliases: `check`, `done`)
- `summary` (aliases: `abstract`, `abstracts`, `overview`, `overviews`)
- `todo` (aliases: `todos`, `todolist`, `task`, `tasks`, `tasklist`, `checklist`, `punchlist`, `outline`, `outlines`)

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMPlusIcons(),
)
```

-----

#### `UseObsidianIcons() Option`

Configures the extension with Obsidian-compatible icons, ideal for users transitioning from or integrating with Obsidian.

Folding and Custom Alerts are enabled. The `[!NOICON]` option is disabled.

**Included Alert Types:**

- `note`
- `abstract` (aliases: `summary`, `tldr`)
- `info`
- `todo` (aliases: `check`, `done`)
- `tip` (aliases: `hint`, `important`)
- `success`
- `question` (aliases: `help`, `faq`)
- `warning` (aliases: `caution`, `attention`)
- `failure` (aliases: `fail`, `missing`)
- `danger` (aliases: `error`)
- `bug`
- `example`
- `quote` (aliases: `cite`)

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseObsidianIcons(),
)
```

-----

#### `WithIcons(icons map[string]string) Option`

Sets a complete custom icon map, replacing any existing icons.

**Parameters:**

- `icons map[string]string`: Map where keys are alert types and values are icon strings (typically SVG)

**Example:**

```go
customIcons := map[string]string{
    "note":      "<svg>...</svg>",
    "warning":   "<svg>...</svg>",
    "custom":    "<svg>...</svg>",
}

extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithIcons(customIcons),
)
```

-----

#### `WithIcon(kind, icon string) Option`

Adds or overrides a single icon without affecting other configured icons. Can be used multiple times and combined with other icon options. If a `kind` already exists it will be overwritten with the new value.

**Parameters:**

- `kind string`: Alert type identifier
- `icon string`: Icon content (typically SVG markup)

**Example:**

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMIcons(), // Start with GFM icons
    alertcallouts.WithIcon("note", "<svg>custom note icon</svg>"), // Override note
    alertcallouts.WithIcon("custom", "<svg>custom alert icon</svg>"), // Add new type
)
```

-----

### Functionality Options

#### `WithFolding(enable bool) Option`

Enables or disables the collapsible callout functionality.

**Parameters:**

- `enable bool`: `true` to enable folding, `false` to disable

**Folding Syntax (when enabled):**

- `> [!TYPE]+` - Creates an open-by-default collapsible callout (using `<details open>` and `<summary>` elements)
- `> [!TYPE]-` - Creates a closed-by-default collapsible callout (using `<details>` and `<summary>` elements)
- `> [!TYPE]` - Creates a non-collapsible alert (standard behavior using `<div>` elements)

**Example:**

```go
// Enable folding (default for AlertCallouts)
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithFolding(true),
)

// Disable folding for simple alerts only
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithFolding(false),
)
```

#### `WithCustomAlerts(enable bool) Option`

Enables or disables the use of custom alerts. If the custom alert doesn't have an icon in the iconset, the extension will use the `default`, `note`, `info`, `tip` or `question` icon name (depending on which is availble in the set). The selection will be made in that order, so if you add a custom icon for `default` using the `WithIcon()` function, this icon will be used for any custom alerts.

For example:

```markdown
> [!CUSTOMIZE]
> This is a custom alert callout type.
```

This would render the alert with the `default` icon (or one of the above icons) and the title `Customize` with proper capitalization.

**Parameters:**

- `enable bool`: `true` to enable folding, `false` to disable

**Example:**

```go
// Enable Custom Alerts
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithCustomAlerts(true),
)

// Disable Custom Alerts
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.WithCustomAlerts(false),
)
```

## Usage Patterns

### Basic Alert Integration

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
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.UseGFMStrictIcons(),
                alertcallouts.WithFolding(true),
            ),
        ),
    )

    source := `> [!NOTE]
> Important information here.`

    var buf bytes.Buffer
    if err := md.Convert([]byte(source), &buf); err != nil {
        panic(err)
    }

    fmt.Println(buf.String())
}
```

-----

### Advanced Configuration

```go
func createAdvancedExtension() goldmark.Extender {
    // Start with GFM+ icons for broad compatibility
    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.UseGFMPlusIcons(),
        alertcallouts.WithFolding(true),
    )

    // Add application-specific icons
    extension = alertcallouts.NewAlertCallouts(
        alertcallouts.UseGFMPlusIcons(),
        alertcallouts.WithIcon("success", successSVG),
        alertcallouts.WithIcon("error", errorSVG),
        alertcallouts.WithFolding(true),
    )

    return extension
}
```

-----

### Integration with Other Extensions

```go
import (
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func createFullFeaturedMarkdown() goldmark.Markdown {
    return goldmark.New(
        goldmark.WithExtensions(
            extension.GFM,                    // GitHub Flavored Markdown
            extension.Footnote,               // Footnotes
            alertcallouts.NewAlertCallouts(   // Alert callouts
                alertcallouts.UseGFMPlusIcons(),
                alertcallouts.WithFolding(true),
            ),
        ),
    )
}
```

## HTML Output Structure

### Basic Alert Output

Non-foldable alerts generate `<div>` structures:

```html
<div class="callout callout-note" data-callout="note">
  <div class="callout-title">
    <svg>...</svg>
    <p class="callout-title-text">Note</p>
  </div>
  <div class="callout-body">
    <p>Alert content</p>
  </div>
</div>
```

### Foldable Callout Output

Foldable callouts use `<details>` and `<summary>` elements:

```html
<details class="callout callout-foldable callout-warning" data-callout="warning" open>
  <summary class="callout-title">
    <svg>...</svg>
    <p class="callout-title-text">Warning</p>
  </summary>
  <div class="callout-body">
    <p>Foldable content</p>
  </div>
</details>
```

### CSS Classes Reference

| Class | Applied To | Purpose |
|-------|------------|---------|
| `callout` | Container element | Base callout styling |
| `callout-foldable` | Container element | Indicates this is foldable content |
| `callout-{type}` | Container element | Type-specific styling (e.g., `callout-note`) |
| `callout-title` | Header element | Title container styling |
| `callout-title-text` | Header Title text | Header Title text styling |
| `callout-body` | Content/Body wrapper | Content/Body area styling |

### Data Attributes

| Attribute | Value | Purpose |
|-----------|--------|---------|
| `data-callout` | Alert type (e.g., "note") | JavaScript targeting and CSS selectors |
| `open` | Present/absent | Default state for `<details>` elements |

## Supported Markdown Syntax

### Alert Types

The extension recognizes alerts in the format:

```markdown
> [!TYPE]
> Content here
```

Where `TYPE` can be any configured alert type (case-insensitive).

### Folding Indicators

When folding is enabled:

- `> [!TYPE]+` - Open by default
- `> [!TYPE]-` - Closed by default
- `> [!TYPE]` - Not foldable (standard `<div>` output)

### Multi-line Content

```markdown
> [!NOTE]
> First paragraph.
>
> Second paragraph with **formatting**.
>
> - List item
> - Another item
>
> ```go
> code := "block"
> ```
```

### Nested Elements

Alerts support all standard Markdown elements:

- Paragraphs and text formatting
- Lists (ordered and unordered)
- Code blocks and inline code
- Links and images
- Tables (when supported by other extensions)

## Error Handling and Edge Cases

### Missing Icons

If an alert type has no configured icon:

- If `WithCustomAlerts(true)`, the alert renders without an icon, otherwise the alert will be rendered with the original markdown text (e.g., if you use `[!CUSTOM]` and Custom Alerts are disabled, the text `[!CUSTOM]` will be rendered as the title).
- HTML structure remains consistent
- No errors are thrown

### Using NoIcon to Force Alert Without Icon

- **This is a feature unique to this extension and is not defined in GFM or Obsidian Alerts/Callouts!**
- Use `[!NOICON] Warning` to render a 'Warning' alert without an icon and will be styled like a
  normal 'Warning' alert (*assuming 'Warning' is a valid alert type*)
- Use `[!NOICON] UnknownAlertType` to render alert with title 'UnknownAlertType' and no icon
- (*run the example code in the `examples` folder for more details*)

### Invalid Alert Types

Invalid alert types:

- Alert types can only contain letters, numbers and underscores (no dashes or other punctuation)
- Fall back to standard blockquote rendering
- Goldmark's default blockquote parser handles the content

### Malformed Syntax

Malformed alert syntax gracefully degrades:

- Missing closing brackets fall back to blockquotes
- Invalid folding indicators are ignored
- Content is preserved and rendered normally

[<-back to README](../README.md)
