package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Its nice method to see the page had a hit ^^
func WritetoConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page!")
		next.ServeHTTP(w, r)
	})
}

//NoSurf adds Csrf protection to all POST request
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//SessionLoad load and saves session in every  request
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
