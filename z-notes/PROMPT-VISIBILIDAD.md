Estoy escribiendo una API en golang con la libreria chi para el router.

Aca esta el programa principal que esta en el paquete main

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
	// GenerateAsimJWTToken (internal)
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



Ahora bien se tienen los modelos en un paquete por aparte (package models)  en la ruta internal/models.go

package models

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Product:  Product{},
		User:     User{},
		Token:    Token{},
		JWTToken: JWTToken{},
	}
}

type Models struct {
	Product  Product
	User     User
	Token    Token
	JWTToken JWTToken
}

Se tiene otro archivo tokens.go que tambien esta en el paquete models (package models) en la ruta  internal/tokens.go, Donde se tienen las funciones asociadas a los modelos.

func (t *Token) AuthenticateTokenII(r *http.Request) (*User, error) {

	token, err := t.ExtractToken(r)
	if err != nil {
		return nil, err
	}

	user, err := t.ValidTokenII(token)
	if err != nil {
		return nil, err
	}

	// return t.ValidTokenII(token), nil
	return user, nil
}

Se tiene un middleware en el paquete main (package main) en el archivo middleware.go

package main

import "net/http"

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.models.Token.AuthenticateTokenII(r)
		// _, err := app.models.Token.AuthenticateToken(r)
		// _, err := app.models.JWTToken.AuthenticateJWTToken(r)
		// _, err := app.models.JWTToken.AuthenticateAsimJWTToken(r)

		if err != nil {
			payload := jsonResponse{
				Error:   true,
				Message: "invalid authentication credentials",
			}

			_ = app.writeJSON(w, http.StatusUnauthorized, payload)
			return
		}
		next.ServeHTTP(w, r)
	})
}

El Objetivo es poder invocar el infoLog y el errorLog que estan declarados en el paquete main desde el paquete models. como se muestra a continuacion:
	app.infoLog.Println("API listening on port", app.config.port)

Que cambios debo generar en la API para poder invocar desde el paquete models, las funciones del paquete main?

POdrias mostrarme las modificaciones necesarias y darme la respuesta en español.








#############################################################################

Existe otra forma de hacerlo sin modificar  el type Models? este type es solo para los modelos de datos.

Solo requiero tener acceso a las funciones de main desde el paquete models.


Considerar que esta porcion de codigo debe conservarse por el tema del pool

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Product:  Product{},
		User:     User{},
		Token:    Token{},
		JWTToken: JWTToken{},
	}
}
Podrias generar una respuesta en español y modificar el codigo para tener la visbilidad de main desde el package models.