package main

import (
	"html/template"
	"path/filepath"

	"github.com/austinthieu/snippetbox/internal/models"
)

// templateData type acts as a holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// filepath.Glob() gets a slice of all file paths that match the pattern
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.tmpl') from the full file path
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
