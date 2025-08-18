# Custom Icon Maps for gm-alert-callouts

This document explains how to create custom icon files and use them with the
`gm-alert-callouts` Goldmark extension.

> [!NOTE]
>
> The project and this document are still under heavy development. It is possible that information
> in this document has not been updated with recent changes to the API for this extension.

## Overview

The `gm-alert-callouts` extension supports custom icon sets through `.icons` files and the
`CreateIconsMap()` function. This allows you to define your own SVG icons for alert callouts or
modify existing icon sets to match your design requirements.

## Icon File Format

Icon files use a simple text format with the following structure:

```text
# Comments start with # and are ignored
# Blank lines are also ignored

# Core icon definitions use: key|svg_content
note|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>
tip|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>

# Alias definitions use: alias->primary_key
info->note
hint->tip
```

### Format Rules

1. **Comments**: Lines starting with `#` are treated as comments and ignored
2. **Blank lines**: Empty lines are ignored
3. **Core definitions**: Use the format `key|svg_content` where:
   - `key` is the alert type name (e.g., `note`, `tip`, `warning`)
   - `svg_content` is the complete SVG markup for the icon
4. **Aliases**: Use the format `alias->primary_key` where:
   - `alias` is an alternative name that maps to an existing key
   - `primary_key` must reference a key that has been defined with SVG content

### Icon Key Guidelines

- Keys are case-sensitive
- Keys should contain only lowercase letters, numbers, and hyphens
- Standard GitHub alert types are: `note`, `tip`, `important`, `warning`, `caution`
- You can define additional custom keys as needed

## Using Custom Icon Files

### Method 1: Embedding with go:embed

This is the recommended approach for packaging icons with your application:

```go
package main

import (
    _ "embed"

    "github.com/yuin/goldmark"
    "github.com/ZMT-Creative/gm-alert-callouts"
)

//go:embed path/to/your/custom.icons
var customIconData string

func main() {
    // Create the icon map from your embedded data
    customIcons := alertcallouts.CreateIconsMap(customIconData)

    // Use with the extension
    md := goldmark.New(
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.WithIcons(customIcons),
                alertcallouts.WithFolding(true),
            ),
        ),
    )

    // Use your markdown processor...
}
```

### Method 2: Loading from File at Runtime

For dynamic icon loading:

```go
package main

import (
    "os"

    "github.com/yuin/goldmark"
    "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    // Read icon file at runtime
    iconData, err := os.ReadFile("path/to/your/custom.icons")
    if err != nil {
        // Handle error
        return
    }

    // Create the icon map
    customIcons := alertcallouts.CreateIconsMap(string(iconData))

    // Use with the extension
    md := goldmark.New(
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.WithIcons(customIcons),
            ),
        ),
    )
}
```

### Method 3: Programmatic Icon Definition

For simple customizations without files:

```go
package main

import (
    "github.com/yuin/goldmark"
    "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    md := goldmark.New(
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.WithIcon("note", `<svg>...</svg>`),
                alertcallouts.WithIcon("tip", `<svg>...</svg>`),
                alertcallouts.UseGFMIcons(), // Load defaults first
            ),
        ),
    )
}
```

## Creating SVG Icons

### Best Practices

1. **Consistent sizing**: Use `width="24" height="24"` for consistency with built-in icons
2. **Scalable design**: Use `viewBox="0 0 24 24"` to ensure proper scaling
3. **Styling**: Include appropriate classes and use `stroke="currentColor"` to inherit text color
4. **Accessibility**: Ensure icons are recognizable at small sizes

### Example SVG Structure

```svg
<svg xmlns="http://www.w3.org/2000/svg"
     width="24"
     height="24"
     viewBox="0 0 24 24"
     fill="none"
     stroke="currentColor"
     stroke-width="2"
     stroke-linecap="round"
     stroke-linejoin="round"
     class="custom-icon">
  <circle cx="12" cy="12" r="10"/>
  <path d="M12 16v-4"/>
  <path d="M12 8h.01"/>
</svg>
```

## Complete Example

Here's a complete example creating a custom icon file:

**custom-icons.icons**:

```text
# My Custom Alert Icons
# Using Material Design icons

note|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/></svg>
warning|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path d="M1 21h22L12 2 1 21zm12-3h-2v-2h2v2zm0-4h-2v-4h2v4z"/></svg>
error|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/></svg>

# Aliases for compatibility
important->note
caution->warning
danger->error
info->note
tip->note
```

**main.go**:

```go
package main

import (
    _ "embed"

    "github.com/yuin/goldmark"
    "github.com/ZMT-Creative/gm-alert-callouts"
)

//go:embed custom-icons.icons
var customIconData string

func main() {
    customIcons := alertcallouts.CreateIconsMap(customIconData)

    md := goldmark.New(
        goldmark.WithExtensions(
            alertcallouts.NewAlertCallouts(
                alertcallouts.WithIcons(customIcons),
                alertcallouts.WithFolding(true),
            ),
        ),
    )

    // Process your markdown...
}
```

## Built-in Icon Sets

The extension includes three built-in icon sets you can reference or extend:

- **GFM Icons** (`UseGFMIcons()`): GitHub Flavored Markdown standard icons
- **GFM Plus Icons** (`UseGFMPlusIcons()`): Extended set with additional icon types
- **Obsidian Icons** (`UseObsidianIcons()`): Obsidian-style icons

You can examine the source files in the `assets/` directory to see their definitions and use them
as templates for your custom icons.

## Troubleshooting

### Common Issues

1. **Icons not displaying**: Check that your SVG markup is valid and properly escaped
2. **Aliases not working**: Ensure the primary key exists before defining aliases
3. **Missing icons**: Verify that all referenced alert types have corresponding icon definitions

### Debugging Tips

- Use `fmt.Printf("%+v\n", customIcons)` to inspect the generated icon map
- Test SVG markup in a browser before adding to icon files
- Check for typos in key names and ensure consistent casing

## API Reference

### CreateIconsMap Function

```go
func CreateIconsMap(iconData string) map[string]string
```

Parses icon data in the defined format and returns a map of icon keys to SVG content.

**Parameters:**

- `iconData` (string): The content of an icon file as a string

**Returns:**

- `map[string]string`: A map where keys are alert types and values are SVG markup

**Usage:**

This function is typically used with Go's `//go:embed` directive to embed icon files at compile
time, but can also be used with dynamically loaded content.
