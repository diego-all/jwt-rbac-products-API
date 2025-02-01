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

	"github.com/golang-jwt/jwt/v4"
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

type AppClaims struct {
	UserId             string `json:"userId"`
	jwt.StandardClaims        /// deprecated
}

type JWTToken struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
	Token  string `json:"token"`
	// TokenHash []byte    `json:"-"`
	TokenHash string    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Role      string    `json:"role,omitempty"`
	SecretKey string
	jwt.RegisteredClaims
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt int64     `json:"expires_at"` // Cambiado a `int64`
	IssuedAt  int64     `json:"issued_at"`  // Cambiado a `int64`
}

//
// claims := jwt.MapClaims{
// 	"username": username,
// 	"exp":      time.Now().Add(time.Hour).Unix(),
// }

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

	fmt.Println("query token: ", &token.Token)

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

// email
// func (t *Token) GenerateJWTToken(userID int, ttl time.Duration) (*Token, error) {
// func (j *JWTToken) GenerateJWTToken(email string) (*JWTToken, error) {
// 	//func (j *JWTToken) GenerateJWTToken(email string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"email": email,
// 		"exp":   time.Now().Add(24 * time.Hour).Unix(),
// 		"role":  "trin",
// 	}

// 	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenJ := JWTToken{
// 		ID:        1,
// 		UserID:    123,
// 		Email:     "user@example.com",
// 		Token:     token,
// 		TokenHash: []byte("hashed-token-value"),
// 		Role:      "admin",
// 		SecretKey: "my-secret-key",
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    "my-app",
// 			Subject:   "user-authentication",
// 			Audience:  jwt.ClaimStrings{"my-service"},
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira en 24 horas
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 		},
// 	}

// 	//return token.SignedString([]byte(j.SecretKey))

// 	return
// 	return token, nil
// 	// TR return JWTToken Object
// }

// func NewToken() JWTToken {
// 	return JWTToken{
// 		SecretKey: "secret", // Cambiar por una clave segura
// 	}
// }

// func (j *JWTToken) GenerateJWTToken(email string) (*JWTToken, error) {

// GenerateJWTToken genera un nuevo token JWT
// func (j *JWTToken) GenerateJWTToken(email string, userID int, role string) (*JWTToken, error) {
func (j *JWTToken) GenerateJWTToken(email string, userID int) (*JWTToken, error) {

	// recordar token string
	token := &JWTToken{
		UserID: userID,
		Expiry: time.Now(),
		Email:  email,
		Role:   "trin",
		// IssuedAt:  jwt.RegisteredClaims.IssuedAt,
		// ExpiresAt: jwt.RegisteredClaims.ExpiresAt,
	}

	secretKey := "secret"

	if secretKey == "" {
		return nil, errors.New("JWT_SECRET no está definido en las variables de entorno")
	}

	// expirationTime := time.Now().Add(24 * time.Hour) // Expira en 24 horas
	// claims := JWTToken{
	// 	UserID:    332434,
	// 	Email:     email,
	// 	Role:      "trin",
	// 	SecretKey: secretKey,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Subject:   email,
	// 		ExpiresAt: jwt.NewNumericDate(expirationTime),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }

	// expirationTime := time.Now().Add(24 * time.Hour)

	// claims := JWTToken{
	// 	UserID:    332434,
	// 	Email:     email,
	// 	Role:      "trin",
	// 	SecretKey: secretKey,
	// 	ExpiresAt: expirationTime.Unix(), // ✅ Convertir a `int64`
	// 	// RegisteredClaims: jwt.RegisteredClaims{
	// 	// 	ExpiresAt: time.Now().Add(2 * time.Hour), // se debe hacer un casting, y considerar campo en la db
	// 	// },
	// 	IssuedAt: time.Now().Unix(), // ✅ Convertir a `int64`
	// }

	// claims2 := AppClaims{
	// 	UserId: user.id,
	// 	StandardClaims: jwt.StandardClaims{
	// 		// value in envvar token Duration
	// 		ExpiresAt: time.Now().Add(2 * time.Hour ),
	// 	},

	// },

	// Algoritmo de firmado
	// tokenn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// JWT is just a Base64-encoded string.
	// token.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, token)

	fmt.Println("TOKEN:", token)
	fmt.Println("tokenObj", tokenObj.Claims)

	// signedToken, err := token.SignedString([]byte(secretKey))
	tokenString, err := tokenObj.SignedString([]byte(secretKey))
	fmt.Println("SIGNEDTOKEN:", tokenString)
	if err != nil {
		// 500
		fmt.Println("Error al firmar el token:", err)
		return nil, err
	}

	// Variable intermedia para almacenar el token en string
	// tokenStringValue := tokenString

	// Convertir el token a string antes de asignarlo
	tokenObjString := tokenObj.Raw

	fmt.Println("tokenObjString", tokenObjString)

	token.Token = tokenObjString
	token.TokenHash = tokenString
	//claims.Token = signedToken
	// return &claims, nil
	return token, nil

	// Luego de tener el token se necesita obtener un string a partir de ese token (firmarlo) tokenString, signedToken
}

// func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
// 	authorizationHeader := r.Header.Get("Authorization")
// 	if authorizationHeader == "" {
// 		return nil, errors.New("no authorization header received")
// 	}

// 	headerParts := strings.Split(authorizationHeader, "")
// 	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 		return nil, errors.New("no valid authorization header received")
// 	}

// 	token := headerParts[1]

// 	if len(token) != 26 {
// 		return nil, errors.New("token wrong size")
// 	}

// 	tkn, err := t.GetByToken(token)
// 	if err != nil {
// 		return nil, errors.New("no matching user found")
// 	}

// 	if tkn.Expiry.Before(time.Now()) {
// 		return nil, errors.New("expired token")
// 	}

// 	user, err := t.GetUserForToken(*tkn)
// 	if err != nil {
// 		return nil, errors.New("no matching user found")
// 	}

// 	return user, nil
// }

func (j *JWTToken) AuthenticateJWTToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no valid authorization header received")
	}

	tokenStr := headerParts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("TOKEN:", token)

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	email := claims["email"].(string)
	// 	user, _ := (&User{}).FindByEmail(email)
	// 	return user, nil
	// }

	// if len(token) != 26 {
	// 	return nil, errors.New("token wrong size")
	// }

	// tkn, err := t.GetByToken(token)
	// if err != nil {
	// 	return nil, errors.New("no matching user found")
	// }

	// if tkn.Expiry.Before(time.Now()) {
	// 	return nil, errors.New("expired token")
	// }

	// user, err := t.GetUserForToken(*tkn)
	// if err != nil {
	// 	return nil, errors.New("no matching user found")
	// }

	// return user, nil
	return nil, errors.New("invalid token")
}

func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {

	fmt.Println(" LLEGO A: AuthenticateToken 1")
	authorizationHeader := r.Header.Get("Authorization")
	fmt.Println(" authorizationHeader: ", authorizationHeader)
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	fmt.Println(" LLEGO A: AuthenticateToken 2")

	headerParts := strings.Split(authorizationHeader, " ")
	fmt.Println("headerParts", headerParts)
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		fmt.Println("ENTRO AL IF")
		fmt.Println("len(headerParts):", len(headerParts))
		fmt.Println("headerParts[0]:", headerParts[0])
		fmt.Println("headerParts[1]:", headerParts[1])
		fmt.Println("headerParts[2]:", headerParts[2])
		return nil, errors.New("no valid authorization header received")
	}

	token := headerParts[1]
	fmt.Println("token token token:", token)

	if len(token) != 26 {
		fmt.Println("ENTRO AL IF TOKEN WRONG SIZE")
		return nil, errors.New("token wrong size")
	}

	tkn, err := t.GetByToken(token)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	if tkn.Expiry.Before(time.Now()) {
		return nil, errors.New("expired token")
	}

	user, err := t.GetUserForToken(*tkn)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

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

	stmt = `insert into tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) values ($1,$2,$3,$4,$5,$6,$7)`

	_, err = db.ExecContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		time.Now(),
		time.Now(),
		token.Expiry,
	)
	if err != nil {
		return err
	}
	return nil
}

func (j *JWTToken) InsertJWT(token JWTToken, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// delete any existing tokens
	stmt := `delete from tokens where user_id = $1`
	// _, err := db.ExecContext(ctx, stmt, token.UserID)
	_, err := db.ExecContext(ctx, stmt, token.UserID)
	if err != nil {
		return err
	}

	fmt.Println("LLEGO AL INSERT")
	fmt.Println("TOKEN", token.Token)
	fmt.Println("TOKENHASH", token.TokenHash)

	token.Email = u.Email

	stmt = `insert into tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) values ($1,$2,$3,$4,$5,$6,$7)`

	_, err = db.ExecContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		token.Expiry,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
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

func (t *Token) ValidToken(plainText string) (bool, error) {
	token, err := t.GetByToken(plainText)

	fmt.Println("PLAINTEXT: ", plainText)

	fmt.Println("TOKEN: ", token.Token)

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

	// if err != nil {
	// 	return false, errors.New("expired token")
	// }

	return true, nil

}
