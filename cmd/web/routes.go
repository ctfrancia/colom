package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", app.home)
	r.Post("/login", app.login)
	r.Post("/submit-acta", app.submitActa)

	return r
}
