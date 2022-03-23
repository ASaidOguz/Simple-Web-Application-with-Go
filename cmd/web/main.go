package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ASaidOguz/Simple-Web-App/pkg/config"
	"github.com/ASaidOguz/Simple-Web-App/pkg/handlers"
	"github.com/ASaidOguz/Simple-Web-App/pkg/render"
)

const PORTNUMBER = ":8080"

// main is the main application where all things begin ^ı^
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cant create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Application starts at :%s", PORTNUMBER))
	_ = http.ListenAndServe(PORTNUMBER, nil)
}
