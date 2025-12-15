package templates

import (
	"html/template"
	"log/slog"
	"os"
)

// NOTE: its better to do this by building a Template Cache (or a "Render Engine") and using Layout Inheritance.

func ParseTemplates(log *slog.Logger) map[string]*template.Template {

	tmplts := make(map[string]*template.Template)

	// Helper to reduce copy-paste error handling
	parse := func(name, path string) {
		t, err := template.ParseFiles(path)
		if err != nil {
			log.Error("failed to parse template", "path", path, "error", err)
			os.Exit(1)
		}
		tmplts[name] = t
	}

	parse("index", "web/templates/index.html")
	parse("404", "web/templates/404.html")

	return tmplts
}
