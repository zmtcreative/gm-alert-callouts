# Alert Icons Syntax

VS Code syntax highlighting for the Alert Callouts (*a Goldmark Extension*) Icons definition files.

## Features

- Syntax highlighting for `.acicons`, `.acicon`, `.aci`, and `.icons` files
- Support for comments (lines starting with `#`)
- Highlighting of icon definitions (`key|<svg>...</svg>`)
- Support for aliases (`alias -> target`)
- SVG content highlighting within icon definitions

## File Format

The extension recognizes files with:

- Extensions: `.acicons`, `.acicon`, `.aci`, `.icons`
- Filenames: `alertcallouts-icons`, `alertcallout-icons`

### Syntax Examples

```properties
# This is a comment
warning|<svg viewBox="0 0 24 24" fill="currentColor"><path d="..."/></svg>
error -> warning
info|<svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="10"/></svg>
```

## Installation

This extension should be automatically loaded by VS Code when placed in the extensions directory.
