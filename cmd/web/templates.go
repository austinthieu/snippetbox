package main

import "github.com/austinthieu/snippetbox/internal/models"

// templateData type acts as a holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet models.Snippet
}
