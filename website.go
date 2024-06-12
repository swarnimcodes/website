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

	// Assets
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	r.Get("/", templ.Handler(views.Home()).ServeHTTP)
	r.Get("/hello", handlers.HandleHello)

	http.ListenAndServe(":6969", r)
}
