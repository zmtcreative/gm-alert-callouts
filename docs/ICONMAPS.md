# Icon Customization Guide

[<-back](../README.md)

This guide explains how to create custom icons and icon maps for the `gm-alert-callouts` extension.

## Overview

The `gm-alert-callouts` extension uses SVG icons to visually distinguish different alert types. You can customize these icons using:

- Built-in icon sets (GFM, GFM Plus, Obsidian)
- Custom icon maps created programmatically
- Icon definition files with the `CreateIconsMap()` function
- Individual icon overrides

> [!IMPORTANT]
>
> When adding custom icons and icon sets in `gm-alert-callouts` you are **only** inserting the icon
> code into the HTML output. You still need to create the necessary CSS styling to format the final
> alert/callout style. An example of a CSS style file for the `GFM Plus` built-in icon set can be
> found in the `examples/css/alertcallouts-gfmplus.css` file. This should provide a starting point
> for customizing the styling for your project.

## CreateIconsMap Function

### Function Signature

```go
func CreateIconsMap(iconData string) map[string]string
```

Parses icon data in a specific format and returns a map suitable for use with the `WithIcons()` option.

**Parameters:**

- `iconData string`: Icon definitions in the defined text format

**Returns:**

- `map[string]string`: Map where keys are alert types and values are SVG markup

### Icon Definition Format

The icon definition format supports:

- **Comments**: Lines starting with `#` (ignored during parsing)
- **Blank lines**: Empty lines (ignored during parsing)
- **Core definitions**: `key|svg_content` format
- **Aliases**: `alias->primary_key` format

#### Format Rules

1. **Core Definitions**: Use `key|svg_content`
   - `key`: Alert type identifier (lowercase recommended)
   - `svg_content`: Complete SVG markup

2. **Aliases**: Use `alias->primary_key`
   - `alias`: Alternative name for an alert type
   - `primary_key`: Must reference an existing core definition

3. **Comments and Whitespace**:
   - Lines starting with `#` are comments
   - Blank lines are ignored
   - Leading/trailing whitespace is trimmed

#### Example Icon Definition File

```text
# Custom Alert Icons
# Core GitHub alert types with Lucide icons

note|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
tip|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg>
important|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/><path d="M12 7v2"/><path d="M12 13h.01"/></svg>
warning|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg>
caution|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/></svg>

# Common aliases for GitHub compatibility
info->note
hint->tip
danger->caution
error->caution
```

## Usage Methods

### Method 1: Embedded Icon Files (Recommended)

Use Go's `//go:embed` directive to embed icon definition files at compile time:

**File: `icons/custom.icons`**

```text
# My Custom Icons
note|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>
warning|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>

# Aliases
info->note
```

**File: `main.go`**

```go
package main

import (
    _ "embed"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

//go:embed icons/custom.icons
var customIconData string

func main() {
    // Create icon map from embedded data
    customIcons := alertcallouts.CreateIconsMap(customIconData)

    // Use with extension
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

### Method 2: Runtime File Loading

Load icon definition files at runtime:

```go
package main

import (
    "os"

    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    // Read icon file at runtime
    iconData, err := os.ReadFile("path/to/icons.icons")
    if err != nil {
        panic(err)
    }

    // Parse icons
    customIcons := alertcallouts.CreateIconsMap(string(iconData))

    // Create extension
    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(customIcons),
        alertcallouts.WithFolding(true),
    )

    md := goldmark.New(goldmark.WithExtensions(extension))
}
```

### Method 3: Inline String Data

Define icon data directly in code:

```go
package main

import (
    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    iconData := `# Inline Icon Definitions
note|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>
warning|<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>

info->note
`

    customIcons := alertcallouts.CreateIconsMap(iconData)

    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(customIcons),
        alertcallouts.WithFolding(true),
    )

    md := goldmark.New(goldmark.WithExtensions(extension))
}
```

### Method 4: Programmatic Icon Maps

Create icon maps directly without using `CreateIconsMap()`:

```go
package main

import (
    "github.com/yuin/goldmark"
    alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
)

func main() {
    // Define icons programmatically
    customIcons := map[string]string{
        "note": `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>`,
        "tip":  `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>`,
        // Add aliases by duplicating values
        "info": `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>`, // Same as note
        "hint": `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">...</svg>`, // Same as tip
    }

    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(customIcons),
        alertcallouts.WithFolding(true),
    )

    md := goldmark.New(goldmark.WithExtensions(extension))
}
```

## SVG Icon Best Practices

### Recommended SVG Structure

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
     class="alert-icon">
  <path d="..."/>
</svg>
```

### Best Practices

1. **Consistent sizing**: Use `width="24" height="24"` for consistency
2. **Scalable design**: Include `viewBox="0 0 24 24"` for proper scaling
3. **Color inheritance**: Use `stroke="currentColor"` or `fill="currentColor"`
4. **Accessibility**: Ensure icons are recognizable at small sizes
5. **Performance**: Minimize SVG complexity for faster rendering

### Icon Design Guidelines

| Guideline | Recommendation | Reason |
|-----------|---------------|---------|
| **Size** | 24x24 pixels | Consistent with common icon libraries |
| **Style** | Outline or solid | Choose one style for consistency |
| **Stroke Width** | 1.5-2 pixels | Good visibility at small sizes |
| **Colors** | `currentColor` | Inherits text color for theming |
| **Complexity** | Simple shapes | Better performance and scalability |

## Built-in Icon Sets Reference

The extension includes three built-in icon sets that you can examine for reference:

### GFM Icons (`UseGFMIcons()`)

GitHub standard alert types with aliases:

- **Core types**: `note`, `tip`, `important`, `warning`, `caution`
- **Aliases**: `info->note`, `hint->tip`, `danger->caution`, `error->caution`

### GFM Plus Icons (`UseGFMPlusIcons()`)

Extended set with additional Obsidian-compatible types:

- All GFM icons plus additional callout types
- Better compatibility with Obsidian workflows

### Obsidian Icons (`UseObsidianIcons()`)

Obsidian-style icons optimized for callouts:

- Matches Obsidian's default callout appearance
- Ideal for users migrating from Obsidian

### Examining Built-in Icons

You can find the source icon definitions in the project's `assets/` folder:

- `assets/alertcallouts-gfm.icons`
- `assets/alertcallouts-gfmplus.icons`
- `assets/alertcallouts-obsidian.icons`

These files serve as examples and templates for creating custom icon sets.

## Advanced Usage Patterns

### Combining Built-in and Custom Icons

Start with a built-in set and override specific icons:

```go
extension := alertcallouts.NewAlertCallouts(
    alertcallouts.UseGFMIcons(),           // Start with GFM icons
    alertcallouts.WithIcon("note", customNoteSVG),    // Override note icon
    alertcallouts.WithIcon("custom", customSVG),      // Add custom type
    alertcallouts.WithFolding(true),
)
```

### Dynamic Icon Loading

Load different icon sets based on configuration:

```go
func createExtensionWithIcons(iconSet string) goldmark.Extender {
    var iconData string
    var err error

    switch iconSet {
    case "custom":
        iconData, err = os.ReadFile("icons/custom.icons")
    case "minimal":
        iconData, err = os.ReadFile("icons/minimal.icons")
    default:
        // Use built-in GFM icons
        return alertcallouts.NewAlertCallouts(
            alertcallouts.UseGFMIcons(),
            alertcallouts.WithFolding(true),
        )
    }

    if err != nil {
        // Fallback to built-in icons
        return alertcallouts.NewAlertCallouts(
            alertcallouts.UseGFMIcons(),
            alertcallouts.WithFolding(true),
        )
    }

    customIcons := alertcallouts.CreateIconsMap(iconData)
    return alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(customIcons),
        alertcallouts.WithFolding(true),
    )
}
```

### Icon Set Inheritance

Extend existing icon sets with additional icons:

```go
//go:embed icons/base.icons
var baseIconData string

//go:embed icons/extensions.icons
var extensionIconData string

func main() {
    // Parse base icons
    baseIcons := alertcallouts.CreateIconsMap(baseIconData)

    // Parse extension icons
    extensionIcons := alertcallouts.CreateIconsMap(extensionIconData)

    // Merge icon sets (extensions override base)
    for key, icon := range extensionIcons {
        baseIcons[key] = icon
    }

    extension := alertcallouts.NewAlertCallouts(
        alertcallouts.WithIcons(baseIcons),
        alertcallouts.WithFolding(true),
    )

    md := goldmark.New(goldmark.WithExtensions(extension))
}
```

## Troubleshooting

### Common Issues

| Issue | Cause | Solution |
|-------|-------|----------|
| Icons not displaying | Invalid SVG markup | Validate SVG in browser |
| Aliases not working | Primary key missing | Define core icon before alias |
| Parse errors | Malformed icon file | Check format and syntax |

### Debugging Tips

1. **Inspect generated map**:

   ```go
   icons := alertcallouts.CreateIconsMap(iconData)
   fmt.Printf("Generated icons: %+v\n", icons)
   ```

2. **Validate SVG markup**:
   - Test SVG content in a browser
   - Use online SVG validators
   - Check for unclosed tags

3. **Check icon file format**:
   - Ensure proper `key|value` format
   - Verify alias syntax (`alias->primary`)
   - Look for invisible characters

### Performance Considerations

- **Icon file size**: Large icon definitions increase memory usage
- **Complex SVGs**: Simplify SVGs for better rendering performance
- **Icon count**: Minimize unused icons in production builds
- **Embedding vs runtime**: Embedded icons are faster but increase binary size

## Migration from Manual Icon Maps

If you're currently using manually created icon maps, consider migrating to icon definition files:

**Before (manual map):**

```go
icons := map[string]string{
    "note": "<svg>...</svg>",
    "info": "<svg>...</svg>", // Duplicate SVG
    "tip":  "<svg>...</svg>",
    "hint": "<svg>...</svg>", // Duplicate SVG
}
```

**After (icon definition file):**

```text
# icons.icons
note|<svg>...</svg>
tip|<svg>...</svg>

# Aliases eliminate duplication
info->note
hint->tip
```

**Benefits of migration:**

- Eliminates SVG duplication
- Easier to maintain and update
- Cleaner code separation
- Support for comments and documentation

[<-back](../README.md)
