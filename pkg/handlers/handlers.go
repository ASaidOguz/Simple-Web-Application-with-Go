package handlers

import (
	"net/http"

	"github.com/ASaidOguz/Simple-Web-App/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.html")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.html")
}
