package models

import (
	"database/sql"
	"jwt-rbac-products-API/internal/config"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB, cfg *config.Config) Models {
	db = dbPool
	return Models{
		Product:  Product{},
		User:     User{},
		Token:    Token{},
		JWTToken: JWTToken{},
		Config:   cfg,
	}
}

type Models struct {
	Product  Product
	User     User
	Token    Token
	JWTToken JWTToken
	Config   *config.Config
}

// var cfg *config.Config
