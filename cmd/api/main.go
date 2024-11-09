package main

import (
	"fmt"
	"jwt-rbac-products-API/database"
	models "jwt-rbac-products-API/internal"

	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   models.Models
}

// const (
// 	DSN = "data.sqlite"
// )

func main() {

	var cfg config
	cfg.port = 9090

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 1. Hardcoded database credentials in connection string
	dsn := "host=localhost port=54325 user=postgres password=password dbname=e_commerce sslmode=disable timezone=UTC connect_timeout=5"

	// db, err := database.ConnectSQLite(DSN)
	db, err := database.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   models.New(db.SQL),
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
