package alertcallouts

import (
	"testing"
)

// Test cases for new initialization methods
func TestNewAlertCalloutsOptions(t *testing.T) {
	t.Run("Default options", func(t *testing.T) {
		ext := NewAlertCallouts()

		if ext == nil {
			t.Fatal("NewAlertCallouts() returned nil")
		}

		if ext.Icons == nil {
			t.Error("Expected Icons to be initialized as empty map")
		}

		if len(ext.Icons) != 0 {
			t.Errorf("Expected empty Icons map, got %d items", len(ext.Icons))
		}

		if ext.FoldingEnabled != true {
			t.Error("Expected FoldingEnabled to be true by default")
		}
	})

	t.Run("With single icon option", func(t *testing.T) {
		ext := NewAlertCallouts(WithIcon("note", "<svg>note icon</svg>"))

		if len(ext.Icons) != 1 {
			t.Errorf("Expected 1 icon, got %d", len(ext.Icons))
		}

		if ext.Icons["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected note icon, got %s", ext.Icons["note"])
		}
	})

	t.Run("With multiple icons option", func(t *testing.T) {
		icons := map[string]string{
			"note":    "<svg>note</svg>",
			"warning": "<svg>warning</svg>",
			"info":    "<svg>info</svg>",
		}

		ext := NewAlertCallouts(WithIcons(icons))

		if len(ext.Icons) != 3 {
			t.Errorf("Expected 3 icons, got %d", len(ext.Icons))
		}

		for kind, expected := range icons {
			if ext.Icons[kind] != expected {
				t.Errorf("Expected %s icon to be %s, got %s", kind, expected, ext.Icons[kind])
			}
		}
	})

	t.Run("Disable folding", func(t *testing.T) {
		ext := NewAlertCallouts(WithFolding(false))

		if ext.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})

	t.Run("With combined options", func(t *testing.T) {
		icons := map[string]string{"tip": "<svg>tip</svg>"}

		ext := NewAlertCallouts(
			WithIcons(icons),
			WithFolding(false),
			WithIcon("important", "<svg>important</svg>"),
		)

		if len(ext.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(ext.Icons))
		}

		if ext.Icons["tip"] != "<svg>tip</svg>" {
			t.Errorf("Expected tip icon, got %s", ext.Icons["tip"])
		}

		if ext.Icons["important"] != "<svg>important</svg>" {
			t.Errorf("Expected important icon, got %s", ext.Icons["important"])
		}

		if ext.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})
}

func TestWithIconOption(t *testing.T) {
	t.Run("Adds icon to nil map", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithIcon("test", "<svg>test</svg>")
		option(opts)

		if opts.Icons == nil {
			t.Fatal("Expected Icons map to be initialized")
		}

		if opts.Icons["test"] != "<svg>test</svg>" {
			t.Errorf("Expected test icon, got %s", opts.Icons["test"])
		}
	})

	t.Run("Adds icon to existing map", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"existing": "<svg>existing</svg>"},
		}

		option := WithIcon("new", "<svg>new</svg>")
		option(opts)

		if len(opts.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(opts.Icons))
		}

		if opts.Icons["existing"] != "<svg>existing</svg>" {
			t.Error("Existing icon should be preserved")
		}

		if opts.Icons["new"] != "<svg>new</svg>" {
			t.Error("New icon should be added")
		}
	})

	t.Run("Overwrites existing icon", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"note": "<svg>old</svg>"},
		}

		option := WithIcon("note", "<svg>new</svg>")
		option(opts)

		if opts.Icons["note"] != "<svg>new</svg>" {
			t.Errorf("Expected icon to be overwritten, got %s", opts.Icons["note"])
		}
	})
}

func TestWithIconsOption(t *testing.T) {
	t.Run("Sets icons map", func(t *testing.T) {
		icons := map[string]string{
			"note":    "<svg>note</svg>",
			"warning": "<svg>warning</svg>",
		}

		opts := &alertCalloutsOptions{}
		option := WithIcons(icons)
		option(opts)

		if len(opts.Icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(opts.Icons))
		}

		for kind, expected := range icons {
			if opts.Icons[kind] != expected {
				t.Errorf("Expected %s icon to be %s, got %s", kind, expected, opts.Icons[kind])
			}
		}
	})

	t.Run("Replaces existing icons", func(t *testing.T) {
		opts := &alertCalloutsOptions{
			Icons: map[string]string{"old": "<svg>old</svg>"},
		}

		newIcons := map[string]string{"new": "<svg>new</svg>"}
		option := WithIcons(newIcons)
		option(opts)

		if len(opts.Icons) != 1 {
			t.Errorf("Expected 1 icon, got %d", len(opts.Icons))
		}

		if opts.Icons["new"] != "<svg>new</svg>" {
			t.Error("Expected new icon")
		}

		if _, exists := opts.Icons["old"]; exists {
			t.Error("Expected old icon to be removed")
		}
	})
}

func TestWithFoldingOption(t *testing.T) {
	t.Run("Enables folding", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithFolding(true)
		option(opts)

		if opts.FoldingEnabled != true {
			t.Error("Expected FoldingEnabled to be true")
		}
	})

	t.Run("Disables folding", func(t *testing.T) {
		opts := &alertCalloutsOptions{}
		option := WithFolding(false)
		option(opts)

		if opts.FoldingEnabled != false {
			t.Error("Expected FoldingEnabled to be false")
		}
	})
}

func TestIconSetOptions(t *testing.T) {
	t.Run("UseGFMIcons", func(t *testing.T) {
		ext := NewAlertCallouts(UseGFMIcons())

		if len(ext.Icons) == 0 {
			t.Error("Expected GFM icons to be loaded")
		}

		// Check for some common GFM alert types
		expectedTypes := []string{"note", "tip", "important", "warning", "caution"}
		for _, alertType := range expectedTypes {
			if _, exists := ext.Icons[alertType]; !exists {
				t.Errorf("Expected GFM icons to include %s", alertType)
			}
		}
	})

	t.Run("UseGFMPlusIcons", func(t *testing.T) {
		ext := NewAlertCallouts(UseGFMPlusIcons())

		if len(ext.Icons) == 0 {
			t.Error("Expected GFM Plus icons to be loaded")
		}
	})

	t.Run("UseObsidianIcons", func(t *testing.T) {
		ext := NewAlertCallouts(UseObsidianIcons())

		if len(ext.Icons) == 0 {
			t.Error("Expected Obsidian icons to be loaded")
		}
	})
}

func TestCreateIconsMapFunction(t *testing.T) {
	t.Run("Basic icon parsing", func(t *testing.T) {
		iconData := `note|<svg>note icon</svg>
warning|<svg>warning icon</svg>`

		icons := CreateIconsMap(iconData)

		if len(icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(icons))
		}

		if icons["note"] != "<svg>note icon</svg>" {
			t.Errorf("Expected note icon, got %s", icons["note"])
		}

		if icons["warning"] != "<svg>warning icon</svg>" {
			t.Errorf("Expected warning icon, got %s", icons["warning"])
		}
	})

	t.Run("Icon parsing with comments and empty lines", func(t *testing.T) {
		iconData := `# This is a comment
note|<svg>note icon</svg>

# Another comment
warning|<svg>warning icon</svg>`

		icons := CreateIconsMap(iconData)

		if len(icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(icons))
		}
	})

	t.Run("Icon aliases", func(t *testing.T) {
		iconData := `note|<svg>note icon</svg>
info -> note`

		icons := CreateIconsMap(iconData)

		if len(icons) != 2 {
			t.Errorf("Expected 2 icons, got %d", len(icons))
		}

		if icons["info"] != "<svg>note icon</svg>" {
			t.Errorf("Expected info to alias note icon, got %s", icons["info"])
		}
	})
}
