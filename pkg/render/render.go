package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ASaidOguz/Simple-Web-App/pkg/config"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, html string) {
	// Get the template cache from App config .
	tc := app.TemplateCache

	t, ok := tc[html]
	if !ok {
		log.Fatal("couldnt get template from template cache!")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error in writing browser")

	}

}

//CreateTemplateCache creates template cache as a map for to locate templates
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err

	}
	for _, page := range pages {
		name := filepath.Base(page)

		// ts , template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
