package main

import (
	"fmt"
	"net/http"
)

type App struct {
	srv *http.Server
}

func main() {

	app := &App{}
	initServer(app)

	fmt.Println("Initialized the services required for this web server")

	app.srv.ListenAndServe()
	fmt.Println("app running on", app.srv.Addr)

}
