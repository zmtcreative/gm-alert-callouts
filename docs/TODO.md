# TODO/Roadmap

<!-- markdownlint-disable MD036 MD051 -->

This is a list of things we may implement or improve as development continues.

-----

## Extension Setup and General Usage

**Planned**

- [ ] Add option(s) to set custom classes (and/or attributes) for HTML elements during rendering
- [ ] Create example CSS style files for all three built-in Icon Sets in the `examples/assets/css` folder
  - [ ] GFM (Github Flavored Markdown) -- for `UseGFMIcons()` and `assets/alertcallouts-gfm.icons`
  - [x] GFM Plus -- for `UseGFMPlusIcons()` and `assets/alertcallouts-gfmplus.icons`
  - [ ] Obsidian -- for `UseObsidianIcons()` and `assets/alertcallouts-obsidian.icons`

**Completed**

- [x] Create **Obsidian-style Folding** (*on by default*) and create `DisableFolding` option to
      allow disabling the feature.
- [x] Streamline calling syntax:
  - [x] Use plain `alertcallouts.AlertCallout` call to enable with defaults (*GFM Icons, Folding
        enabled*)
  - [x] **REMOVED** -- Use `alertcallouts.AlertCalloutOptions{}` syntax to customize options -- Currently:
    - `Icons:` pass a `map[string]string` with a key/value set of markers and icons values
    - `DisableFolding:` pass `true` or `false`
- [x] Add `NewAlertCallouts()` calling function to provide more complex setup options moving
      forward
- [x] Add `CreateIconMap()` helper function to allow users to pass custom icon configuration data
      using a specially-formatted custom list of markers and icon definitions (*see [Custom Icon Definition File](#custom-icon-definition-file) later in this document*)
  - [x] Add `assets` folder to hold icon file maps for embedding

-----

## Parsing and Rendering

**Planned/In-Development**

- [ ] Insert custom classes into specific HTML elements as defined during the extension initialization process.

**Completed**

- [x] **Cleanup** -- Remove legacy `gh-alerts` and related classes from HTML element output
