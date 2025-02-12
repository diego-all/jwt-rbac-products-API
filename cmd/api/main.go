package main

import (
	"crypto/tls"
	"fmt"
	"jwt-rbac-products-API/database"
	models "jwt-rbac-products-API/internal"
	"time"

	"log"
	"net/http"
	"os"
)

type config struct {
	port           int
	certPath       string
	keyPath        string
	jwt_secret     string
	hash_cost      int // bcrypt
	token_duration time.Duration
	jwt_PrivateKey string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   models.Models
}

// As an improvement you could do this as an environment variable. Or create your own configuration file to handle these types of values
// ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),

// const (
// 	DSN = "data.sqlite"
// )

func main() {

	var cfg config
	cfg.port = 9090

	cfg.certPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.pem"
	cfg.keyPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.key"
	cfg.jwt_PrivateKey = "/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_private.pem"
	cfg.jwt_secret = "secret"
	cfg.token_duration = 24 * time.Hour
	cfg.hash_cost = 8

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 1. Hardcoded database credentials in connection string
	dsn := "host=localhost port=54325 user=postgres password=password dbname=e_commerce sslmode=disable timezone=UTC connect_timeout=5"

	// db, err := database.ConnectSQLite(DSN)
	//db, err := database.ConnectPostgres(dsn)
	db, err := database.ConnectPostgresPQ(dsn)

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

// func (app *application) serve() error {
// 	app.infoLog.Println("API listening on port", app.config.port)

// 	srv := &http.Server{
// 		Addr:    fmt.Sprintf(":%d", app.config.port),
// 		Handler: app.routes(),
// 	}
// 	return srv.ListenAndServe()
// }

// Pending attach *.key and *.pem with Postman.

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	// Configuramos el servidor TLS para solo aceptar TLS 1.3
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13, // Versión mínima TLS 1.3
		MaxVersion: tls.VersionTLS13, // Versión máxima también TLS 1.3
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%d", app.config.port),
		Handler:   app.routes(),
		TLSConfig: tlsConfig,
	}
	//return srv.ListenAndServe()
	return srv.ListenAndServeTLS(app.config.certPath, app.config.keyPath)
}
