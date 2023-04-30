package main

import (
	"github.com/go-chi/chi/v5"
)

func initHandler(r *chi.Mux) {
	r.Get("/:id", readPaste)
	r.Post("/", writePaste)
}
