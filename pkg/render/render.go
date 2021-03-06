package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ASaidOguz/Simple-Web-App/pkg/config"
	"github.com/ASaidOguz/Simple-Web-App/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from App config .
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[html]
	if !ok {
		log.Fatal("couldnt get template from template cache!")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

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
