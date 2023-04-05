package main

import (
	"net/http"
)

func initServer(app *App) {

	srv := http.Server{
		Addr:    "localhost:3000",
		Handler: nil,
	}

	app.srv = &srv

}
