package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error in parsing template")
		return
	}
}
