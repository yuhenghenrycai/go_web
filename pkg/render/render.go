package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/yuhenghenrycai/go_web/pkg/config"
)

var app *config.AppConfig

// set appconfig for render package
func NewRender(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("failed to get template from cache")
	}

	// write to buf first so that we can inspect the bytes sending back to client
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	n, err := buf.WriteTo(w)
	log.Printf("number of bytes sent to client is %d", n)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return templateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tp, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}
		// if no layout file in the directory, then we don't waste effort in parsing further
		layouts, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return templateCache, err
		}
		if len(layouts) > 0 {
			tp, err = tp.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[name] = tp
	}
	return templateCache, nil
}
