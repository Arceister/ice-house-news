package main

import (
	"context"
	"net/http"

	"github.com/Arceister/ice-house-news/lib"
	server "github.com/Arceister/ice-house-news/server"
)

func main() {
	r := server.NewServer()

	config := lib.NewConfig()
	app := config.App
	db := config.Database

	database := lib.NewDB(db)
	database.DB.Query(context.Background(), "SELECT 1=1 AS RESULT")

	http.ListenAndServe(app.Port, r.Chi)
}
