package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	alertcallouts "github.com/ZMT-Creative/gm-alert-callouts"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed assets/css/alertcallouts-gfmplus.css
var cssData []byte

//go:embed assets/markdown/sample-gfmplus.md
var sample string

// //go:embed assets/iconsets/alertcallouts-simple.icons
// var iconSet string

func main() {
	md := CreateGoldmarkInstance(createOptions{
		useAlertCallouts: true,
		enableGFM:   true,
	})

	// Example markdown with GitHub alerts (read from embedded sample)
	mdSource := sample
	// Convert CRLF line endings to LF for consistency in processing markdown source
	// (some plugins perform better with LF line endings -- not sure why, but this has been our experience)
	mdSource = strings.ReplaceAll(mdSource, "\r\n", "\n")
	var buf bytes.Buffer
	if err := md.Convert([]byte(mdSource), &buf); err != nil {
		panic(err)
	}

	fmt.Printf(`<html><head><style type="text/css">%s</style></head><body>%s</body></html>`, cssData, buf.String())
}

type createOptions struct {
	useAlertCallouts 	bool
	enableGFM       bool
}

// CreateGoldmarkInstance creates and configures a new Goldmark instance.
func CreateGoldmarkInstance(opt createOptions) goldmark.Markdown {
	// Default initialization options -- basic Goldmark instance
	options := []goldmark.Option{
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Automatically generate IDs for headings
			parser.WithAttribute(),     // Enable attributes for nodes
		),
		goldmark.WithExtensions(),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	}

	// Add GFM-related extensions and PHP Markdown Extensions if enabled
	if opt.enableGFM {
		options = append(options,
			goldmark.WithExtensions(
				extension.GFM,
				extension.DefinitionList,
				extension.Footnote,
			),
		)
	}

	// Add GitHub Alert Callouts extension if enabled
	if opt.useAlertCallouts {
		// myIcons := InitAlertCalloutsIcons() // Initialize alert icons
		alertCalloutsOpts := alertcallouts.NewAlertCallouts(
			// alertcallouts.WithIcons(myIcons),
			alertcallouts.UseGFMPlusIcons(),
			// alertcallouts.WithIcons(alertcallouts.CreateIconsMap(iconSet)),
			alertcallouts.WithFolding(true),
		)
		options = append(options,
			goldmark.WithExtensions(alertCalloutsOpts),
		)
	}

	return goldmark.New(options...)
}

