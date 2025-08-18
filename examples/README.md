# Alert Callouts Examples

This folder contains a demonstration program that shows how to use the
`gm-alert-callouts` extension with Goldmark to render GitHub-style alert
callouts in HTML.

## What it does

The `showalerts.go` program:

- Converts Markdown files containing alert callouts (like `> [!NOTE]`, `> [!TIP]`, etc.) into HTML
- Uses the GFM Plus icon set for enhanced visual styling
- Supports folding/collapsible callouts
- Includes embedded CSS styling for a complete rendered output (*CSS source in
  `assets/css/alertcallouts-gfmplus.css`*)
- Auto-refreshes the HTML output every 10 seconds for development convenience

## Running the example

### Windows (PowerShell)

```powershell
# Run with the default embedded sample file in `assets/markdown/sample-gfmplus.md`
.\run-showalerts.ps1

# Run with a specific markdown file
.\run-showalerts.ps1 -File "path\to\your\file.md"

# Run and automatically open in browser
.\run-showalerts.ps1 -View

# Run with file and open in browser
.\run-showalerts.ps1 -File "assets\markdown\sample-gfmplus.md" -View
```

### Linux/macOS (Bash)

```bash
# Run with the default embedded sample and open in browser
./run-showalerts.sh

# Run with a specific file (modify the script or run directly)
go run ./showalerts.go -f "assets/markdown/sample-gfmplus.md" > example.html
```

## Sample files

The `assets/markdown/` folder contains example files:

- `sample.md` - Basic alert callout examples
- `sample-gfmplus.md` - Comprehensive examples with all callout types and features showcasing the
  `GFM Plus` icon set.

## Output

The program generates `example.html` which can be opened in any web browser to see the rendered
alert callouts.
