package main
import (
	"github.com/DerrickKirimi/Snippets/internal/models"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//extract name like home.tmpl to name variable
		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}