package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swarnimcodes/website/handlers"
	"github.com/swarnimcodes/website/views"
)

func main() {
	fmt.Println("\nServer has started!")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", templ.Handler(views.Home()).ServeHTTP)
	r.Get("/hello", handlers.HandleHello)

	http.ListenAndServe(":6969", r)
}
