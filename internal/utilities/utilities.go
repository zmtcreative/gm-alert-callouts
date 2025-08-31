package utilities

import (
	"regexp"
	"strings"
)

func FindNamedMatches(regex *regexp.Regexp, str string) map[string]string {
    match := regex.FindStringSubmatch(str)

    results := map[string]string{}
    if match == nil {
        return results
    }

    for i, name := range match {
        results[regex.SubexpNames()[i]] = name
    }
    return results
}

// CreateIconsMap creates a map of icon names to their SVG data from the given icon data string.
func CreateIconsMap(icondata string) map[string]string {
	iconmap := make(map[string]string)
	wordRegex := regexp.MustCompile(`^\p{L}[\p{L}\p{N}_-]*$`) // update the wordRegex to allow full Unicode support

	// Parse the embedded alert callouts data

	// First pass: load all primary icons.
	lines := strings.Split(icondata, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines, comments, and alias definitions.
		if line == "" || strings.HasPrefix(line, "#") || strings.Contains(line, "->") {
			continue
		}

		// Parse core icon definition (key|svg).
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 {
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			svg := strings.TrimSpace(parts[1])

			// Reserved prefixes, and invalid key patterns are skipped.
			if strings.HasPrefix(key, "noicon-") || strings.HasPrefix(key, "noicon_") || !wordRegex.MatchString(key) {
				continue
			}
			iconmap[key] = svg
		}
	}

	// Second pass: process all aliases.
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines, comments, and anything that is not an alias.
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "->") {
			continue
		}

		parts := strings.SplitN(line, "->", 2)
		if len(parts) == 2 {
			alias := strings.ToLower(strings.TrimSpace(parts[0]))
			primary := strings.ToLower(strings.TrimSpace(parts[1]))

			// Reserved prefixes and invalid patterns are skipped.
			if strings.HasPrefix(alias, "noicon-") || strings.HasPrefix(alias, "noicon_") || !wordRegex.MatchString(alias) || !wordRegex.MatchString(primary) {
				continue
			}

			// If the primary key exists, create the alias.
			if svg, exists := iconmap[primary]; exists {
				iconmap[alias] = svg
			}
		}
	}

	return iconmap
}
