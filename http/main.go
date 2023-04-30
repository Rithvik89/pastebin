package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"pastebin/internal/read_service"
	"pastebin/internal/write_service"
)

type App struct {
	srv           *http.Server
	DB            *sql.DB
	write_service *write_service.WriteService
	read_service  *read_service.ReadService
}

func main() {

	app := &App{}
	initServer(app)
	initDB(app)

	fmt.Println("Initialized the services required for this web server")

	app.srv.ListenAndServe()
	fmt.Println("app running on", app.srv.Addr)

}
