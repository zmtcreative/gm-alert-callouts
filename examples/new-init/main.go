package main

import (
	"bytes"
	"fmt"
	"os"

	alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
	"github.com/yuin/goldmark"
)

func main() {
	// Example using the new NewAlertCallouts() initialization method
	extension := alertcallouts.NewAlertCallouts(
		// Add icons one by one
		alertcallouts.WithIcon("note", `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>`),
		alertcallouts.WithIcon("warning", `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-alert-triangle"><path d="M21.73 18 12 2 2.27 18l9.73 2Z"/><path d="M12 9v4"/><path d="M12 17h.01"/></svg>`),
		alertcallouts.WithIcon("tip", `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-lightbulb"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/></svg>`),

		// Enable folding functionality (this is the default)
		alertcallouts.WithFolding(true),
	)

	// Create Goldmark instance with the extension
	markdown := goldmark.New(goldmark.WithExtensions(extension))

	// Example markdown content with various alert types
	mdSource := `# New Initialization Method Example

## Basic Alerts

> [!NOTE]
> This is a note alert using the new initialization method.

> [!WARNING]
> This is a warning alert with an icon.

> [!TIP]
> This is a helpful tip for users.

## Foldable Alerts

> [!NOTE]-
> This is a closed note alert (click to expand).
>
> Hidden content here!

> [!WARNING]+
> This is an open warning alert (click to collapse).
>
> - Important item 1
> - Important item 2
> - Important item 3

> [!TIP]- Pro Tips
> This foldable tip has a custom title.
>
> 1. Always test your code
> 2. Write comprehensive documentation
> 3. Use meaningful commit messages

## Complex Content

> [!NOTE] Advanced Features
> This alert demonstrates complex content:
>
> ### Code Example
> ` + "```go" + `
> func main() {
>     fmt.Println("Hello, World!")
> }
> ` + "```" + `
>
> ### Lists and Links
> - Visit [Goldmark](https://github.com/yuin/goldmark)
> - Check out the [documentation](https://pkg.go.dev/github.com/yuin/goldmark)
>
> *Markdown* **formatting** works perfectly inside alerts!
`

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(mdSource), &buf); err != nil {
		fmt.Printf("Error converting markdown: %v\n", err)
		os.Exit(1)
	}

	// Create complete HTML document
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>New AlertCallouts Initialization Example</title>
    <link rel="stylesheet" href="../css/alertcallouts.css">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            line-height: 1.6;
        }
        pre {
            background: #f6f8fa;
            padding: 16px;
            border-radius: 6px;
            overflow-x: auto;
        }
        code {
            background: #f6f8fa;
            padding: 2px 4px;
            border-radius: 3px;
            font-size: 0.9em;
        }
        pre code {
            background: none;
            padding: 0;
        }
    </style>
</head>
<body>
%s
<hr>
<p><em>This example was generated using the new <code>NewAlertCallouts()</code> initialization method with functional options.</em></p>
</body>
</html>`, buf.String())

	// Write to file
	if err := os.WriteFile("new-init-example.html", []byte(html), 0644); err != nil {
		fmt.Printf("Error writing HTML file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated new-init-example.html successfully!")
	fmt.Println("Open the file in your browser to see the rendered alerts.")
}
