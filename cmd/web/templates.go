package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/ctfrancia/colom/pkg/forms"
)

type templateData struct {
	CSRFToken       string
	CurrentYear     int
	Form            *forms.Form
	Flash           string
	IsAuthenticated bool
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// humanDate takes a UTC time and returns a formatted string for better human legibility
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
