package main

import (
	"net/http"

	"github.com/ASaidOguz/Simple-Web-App/pkg/config"
	"github.com/ASaidOguz/Simple-Web-App/pkg/handlers"
	"github.com/bmizerany/pat"
)

//routes; we will create new mux(http handler) inside of this function
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
