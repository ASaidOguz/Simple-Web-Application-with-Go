package main

import (
	"fmt"
	"net/http"

	"github.com/ASaidOguz/Simple-Web-App/pkg/handlers"
)

const PORTNUMBER = ":8080"

// main is the main application where all things begin ^ı^
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Application starts at :%s", PORTNUMBER))
	_ = http.ListenAndServe(PORTNUMBER, nil)
}
