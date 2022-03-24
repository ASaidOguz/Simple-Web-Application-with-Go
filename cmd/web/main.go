package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ASaidOguz/Simple-Web-App/pkg/config"
	"github.com/ASaidOguz/Simple-Web-App/pkg/handlers"
	"github.com/ASaidOguz/Simple-Web-App/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORTNUMBER = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application where all things begin ^ı^
func main() {

	app.InProduction = false

	//session we r intializing alexedwards session packages here
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cant create template cache")
	}

	app.TemplateCache = tc

	//UseCache if its true its in production mode-if false its in development mode
	//which we will create template everytime
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Application starts at %s", PORTNUMBER))

	srv := &http.Server{
		Addr:    PORTNUMBER,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
