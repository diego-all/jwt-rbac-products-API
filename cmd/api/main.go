package main

import (
	"crypto/tls"
	"fmt"
	"jwt-rbac-products-API/database"
	models "jwt-rbac-products-API/internal"
	"jwt-rbac-products-API/internal/config"
	"jwt-rbac-products-API/internal/logger"

	"log"
	"net/http"
)

// type config struct {
// 	port           int
// 	certPath       string
// 	keyPath        string
// 	jwt_secret     string
// 	hash_cost      int // bcrypt
// 	token_duration time.Duration
// 	jwt_PrivateKey string
// }

type application struct {
	config   *config.Config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   *models.Models
	// environment string // Load by env var or docker-compose
}

// // Métodos que estarán disponibles en `models`
// func (app *application) Info(msg string) {
// 	app.infoLog.Println("INFO:", msg)
// }

// func (app *application) Error(msg string) {
// 	app.errorLog.Println("ERROR:", msg)
// }

// As an improvement you could do this as an environment variable. Or create your own configuration file to handle these types of values
// ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),

// const (
// 	DSN = "data.sqlite"
// )

func main() {

	// var cfg config

	// cfg.port = 9090
	// cfg.certPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.pem"
	// cfg.keyPath = "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.key"
	// // GenerateAsimJWTToken (internal)
	// cfg.jwt_PrivateKey = "/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_private.pem"
	// cfg.jwt_secret = "secret"
	// cfg.token_duration = 24 * time.Hour
	// cfg.hash_cost = 8

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error al cargar la configuración:", err)
	}

	logger.Init()
	logger.InfoLog.Println("Servidor iniciado")

	// 1. Hardcoded database credentials in connection string
	// dsn := "host=localhost port=54325 user=postgres password=password dbname=e_commerce sslmode=disable timezone=UTC connect_timeout=5"

	// db, err := database.ConnectSQLite(DSN)
	//db, err := database.ConnectPostgres(dsn)
	db, err := database.ConnectPostgresPQ(cfg.DatabaseDSN)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  logger.InfoLog,
		errorLog: logger.ErrorLog,
		models:   models.New(db.SQL, cfg),
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
	app.infoLog.Println("API listening on port", app.config.Port)

	// Configuramos el servidor TLS para solo aceptar TLS 1.3
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13, // Versión mínima TLS 1.3
		MaxVersion: tls.VersionTLS13, // Versión máxima también TLS 1.3
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%d", app.config.Port),
		Handler:   app.routes(),
		TLSConfig: tlsConfig,
	}
	//return srv.ListenAndServe()
	return srv.ListenAndServeTLS(app.config.CertPath, app.config.KeyPath)
}
