package main

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// adding routes here
	a.App.Routes.Get("/", a.Handlers.Home)

	a.App.Routes.Get("/jet", func(w http.ResponseWriter, r *http.Request) {
		a.App.Renter.JetPage(w, r, "testjet", nil, nil)
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
