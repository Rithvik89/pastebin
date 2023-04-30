package main

import (
	"database/sql"
	"net/http"
	"pastebin/internal/read_service"
	"pastebin/internal/write_service"
	"pastebin/sqlc_db"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func initServer(app *App) {

	r := chi.NewRouter()

	initHandler(r)

	srv := http.Server{
		Addr:    "localhost:5001",
		Handler: r,
	}

	app.srv = &srv

}

func initDB(app *App) {
	db, err := sql.Open("mysql", "root:mypassword@/mysql_pb")

	if err != nil {
		panic(err)
	}

	app.DB = db
}

func initServices(app *App) {
	querier := sqlc_db.New(app.DB)
	app.write_service = write_service.New(querier)
	app.read_service = read_service.New(querier)
}
