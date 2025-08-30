package utilities

import (
	"regexp"
	"strings"
)

// IsNoIconKind returns true if the kind string indicates that no icon should be rendered.
func IsNoIconKind(kind string) bool {
	kind = strings.ToLower(strings.TrimSpace(kind))
	switch kind {
	case "noicon", "no_icon", "none", "nil", "null":
		return true
	default:
		return false
	}
}

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
	// update the wordRegex to allow full Unicode support
	wordRegex := regexp.MustCompile(`^\p{L}[\p{L}\p{N}_-]*$`)

	// Parse the embedded alert callouts data
	lines := strings.Split(icondata, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Check if it's an alias definition (contains ->)
		if strings.Contains(line, "->") {
			parts := strings.SplitN(line, "->", 2)
			if len(parts) == 2 {
				alias := strings.TrimSpace(parts[0])
				primary := strings.TrimSpace(parts[1])

				// Validate alias and primary values match [\w]+ pattern
				if !wordRegex.MatchString(alias) || !wordRegex.MatchString(primary) {
					continue // Skip invalid alias entries
				}

				// Set alias to reference the primary icon (will be set after core icons are loaded)
				if svg, exists := iconmap[primary]; exists {
					iconmap[alias] = svg
				} else {
					// Store for later processing if primary doesn't exist yet
					// This handles the case where aliases are defined before their primary keys
					defer func(alias, primary string) {
						if svg, exists := iconmap[primary]; exists {
							iconmap[alias] = svg
						}
					}(alias, primary)
				}
			}
			continue
		}

		// Parse core icon definition (key|svg)
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			svg := strings.TrimSpace(parts[1])

			// Validate key matches [\w]+ pattern
			if !wordRegex.MatchString(key) {
				continue // Skip invalid key entries
			}

			iconmap[key] = svg
		}
	}

	// Second pass to handle any aliases that couldn't be resolved in first pass
	lines = strings.Split(icondata, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "->") {
			parts := strings.SplitN(line, "->", 2)
			if len(parts) == 2 {
				alias := strings.TrimSpace(parts[0])
				primary := strings.TrimSpace(parts[1])

				// Validate alias and primary values match [\w]+ pattern
				if !wordRegex.MatchString(alias) || !wordRegex.MatchString(primary) {
					continue // Skip invalid alias entries
				}

				if svg, exists := iconmap[primary]; exists {
					iconmap[alias] = svg
				}
			}
		}
	}

	return iconmap
}
