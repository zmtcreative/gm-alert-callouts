# TODO/Roadmap

This is a list of things we may implement or improve as development continues.

## Extension Setup

- [x] Create **Obsidian-style Folding** (*on by default*) and create `DisableFolding` option to
      allow disabling the feature.
- [x] Streamline calling syntax:
  - [x] Use plain `alertcallouts.AlertCallout` call to enable with defaults (*no icons, folding
        enabled*)
  - [x] Use `alertcallouts.AlertCalloutOptions{}` syntax to customize options -- Currently:
    - `Icons:` pass a `map[string]string` with a key/value set of markers and icons values
    - `DisableFolding:` pass `true` or `false`
- [ ] Add option(s) to set custom classes (and/or attributes) for HTML elements during rendering
- [x] Add `NewAlertCallouts()` calling function to provide more complex setup options moving
      forward
- [x] Add `CreateIconMap()` helper function to allow users to pass custom icon configuration data
      using a specially-formatted custom list of markers and icon definitions (*see [Custom Icon
      Definition File](#custom-icon-definition-file) later in this document*)
  - [x] Add `assets` folder to hold icon file maps for embedding

## Rendering

- [x] **Cleanup** -- Remove legacy `gh-alerts` and related classes from HTML element output
- [ ] Insert custom classes into specific HTML elements as defined during the extension initialization process.

## Custom Icon Definition File

This is the definition of a simple icon configuration file using a key/value syntax that is quick
and simple to parse in the extension. The file can be directly read by a function or it can be
embedded using the `//go:embed` syntax.

```properties
# Alert Icons Definition File
# Format: key|svg_content
# Lines starting with # are comments and will be ignored
# Blank lines are ignored

# Core icon definitions
note|<svg xmlns="http://www.w3.org/2000/svg" ...rest of svg code here.../></svg>
abstract|<svg xmlns="http://www.w3.org/2000/svg" ...rest of svg code here.../></svg>
info|<svg xmlns="http://www.w3.org/2000/svg" ...rest of svg code here.../></svg>

# Alias definitions (format: alias->primary_key)
summary->abstract
```

The implemented `CreateIconMap()` function takes a `string` value containing the text data as shown
above and parses it into a `map[string]string` variable and returns that variable when parsing is
complete.
