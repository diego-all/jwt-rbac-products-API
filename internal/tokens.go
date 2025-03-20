package models

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	ErrNoAuthHeader      = errors.New("no authorization header received")
	ErrInvalidAuthHeader = errors.New("invalid authorization header format")
	ErrTokenSizeMismatch = errors.New("token wrong size")
	ErrTokenNotFound     = errors.New("no matching user found")
	ErrTokenExpired      = errors.New("expired token")
)

func (t *Token) GetByToken(plainText string) (*Token, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, user_id, email, token, token_hash, expiry, created_at, updated_at from tokens where token = $1`

	var token Token

	row := db.QueryRowContext(ctx, query, plainText)

	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Email,
		&token.Token,
		&token.TokenHash,
		&token.Expiry,
		&token.CreatedAt,
		&token.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Artesanal usada antes en el handler de Login()
func (t *Token) GetDataForUpdateHandlerToken(plainText string) (*Token, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, user_id, email, token, token_hash, expiry, created_at, updated_at from tokens where token = $1`

	var token Token

	row := db.QueryRowContext(ctx, query, plainText)

	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Email,
		&token.Token,
		&token.TokenHash,
		&token.Expiry,
		&token.CreatedAt,
		&token.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (t *Token) GetUserForToken(token Token) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, created_at, updated_at from users where id = $1`

	var user User
	row := db.QueryRowContext(ctx, query, token.UserID)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (t *Token) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token))
	token.TokenHash = hash[:]

	return token, nil
}

// Called by middleware
// AuthenticateToken takes the full http request, extracts the authorization header,
// takes the plain text token from that header and looks up the associated token entry
// in the database, and then finds the user associated with that token. If the token
// is valid and a user is found, the user is returned; otherwise, it returns an error.
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {

	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no valid authorization header received")
	}

	token := headerParts[1]

	if len(token) != 26 {
		return nil, errors.New("token wrong size")
	}

	tkn, err := t.GetByToken(token)
	if err != nil {
		fmt.Println("no matching user found")
		return nil, errors.New("no matching user found")
	}

	if tkn.Expiry.Before(time.Now()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")
	}

	// if tkn.Expiry.UTC().Before(time.Now().UTC()) {
	// 	fmt.Println("EXPIRED TOKEN")
	// 	return nil, errors.New("expired token")
	// }

	user, err := t.GetUserForToken(*tkn)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	return user, nil
}

// Principio de responsabilidad única (SRP - Single Responsibility Principle)
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

func (t *Token) Insert(token Token, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// delete any existing tokens
	stmt := `delete from tokens where user_id = $1`
	_, err := db.ExecContext(ctx, stmt, token.UserID)
	if err != nil {
		return err
	}

	token.Email = u.Email

	stmt = `insert into tokens (user_id, email, token, token_hash, expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7)`

	_, err = db.ExecContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		token.Expiry,
		time.Now(),
		time.Now(),
	)

	fmt.Println("DESDE INSERT TOKEN", token)
	if err != nil {
		return err
	}
	return nil
}

// stmt = `INSERT INTO tokens (user_id, email, token, token_hash, created_at, updated_at, expiry)
// VALUES ($1, $2, $3, $4, $5, $6, $7)
// RETURNING id, created_at, updated_at`

// Recordar tema de la modificacion de punteros ...

// func (t *Token) GetDataForUpdateHandlerToken(plainText string) (*Token, error) {
// (*Token, error)
// func (t *Token) InsertReturning(token Token, u User) error {
func (t *Token) InsertReturning(token Token, u User) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// delete any existing tokens
	stmt := `delete from tokens where user_id = $1`
	_, err := db.ExecContext(ctx, stmt, token.UserID)
	if err != nil {
		return nil, err
	}

	token.Email = u.Email

	// stmt = `insert into tokens (user_id, email, token, token_hash, expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7)`
	stmt = `insert into tokens (user_id, email, token, token_hash, expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7) returning id, created_at, updated_at`

	// _, err = db.ExecContext(ctx, stmt,
	// 	token.UserID,
	// 	token.Email,
	// 	token.Token,
	// 	token.TokenHash,
	// 	token.Expiry,
	// 	time.Now(),
	// 	time.Now(),
	// )

	err = db.QueryRowContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		token.Expiry,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&token.ID, &token.CreatedAt, &token.UpdatedAt)

	if err != nil {
		fmt.Println("Error ejecutando INSERT:", err)
		return &token, err
	}
	return &token, nil
}

func (t *Token) DeleteByToken(plaintText string) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from tokens where token = $1`
	_, err := db.ExecContext(ctx, stmt, plaintText)

	if err != nil {
		return err
	}
	return nil
}

// Validar si es exclusivamente para test
func (t *Token) ValidToken(plainText string) (bool, error) {
	token, err := t.GetByToken(plainText)

	fmt.Println("PLAINTEXT: ", plainText)

	fmt.Println("TOKEN: ", token.Token)
	fmt.Println("DESDE MODEL-VALID-TOKENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN")

	if err != nil {
		return false, errors.New("no matching token found")
	}

	_, err = t.GetUserForToken(*token)
	if err != nil {
		return false, errors.New("no matching user found")
	}

	// TR
	if token.Expiry.Before(time.Now()) {
		return false, errors.New("expired token")
	}

	return true, nil

}

// Extracts the Bearer token from the Authorization header of an HTTP request.
// If the header is missing or improperly formatted, it returns an error.
func (t *Token) ExtractToken(r *http.Request) (string, error) {

	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		return "", ErrNoAuthHeader
		// return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", ErrInvalidAuthHeader
		// return nil, errors.New("no valid authorization header received")
	}

	token := headerParts[1]

	return token, nil
}

// Validate a token and return the associated user.
func (t *Token) ValidTokenII(plainText string) (*User, error) {

	// Se debe traer la logica para validar el token.
	// Luego renombrar ValidTokenII() por ValidToken() del andamio
	if len(plainText) != 26 {
		return nil, ErrTokenSizeMismatch
		// return nil, errors.New("token wrong size")
	}

	tkn, err := t.GetByToken(plainText)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	if tkn.Expiry.Before(time.Now()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")

	}

	user, err := t.GetUserForToken(*tkn)
	if err != nil {
		return nil, ErrTokenNotFound
		// return nil, errors.New("no matching user found")
	}

	// app.infoLog.Println("API listening on port", app.config.port)

	// return user, nil
	return user, nil
}
