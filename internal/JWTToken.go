package models

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// It is not recommended to define unnecessary properties or claims. They make the JWT token heavier.
type JWTToken struct {
	// ID     int    `json:"id"`
	UserID int    `json:"user_id,omitempty"` //MANDATORY*
	Email  string `json:"email,omitempty"`   //MANDATORY* ?
	Token  string `json:"token"`
	// TokenHash []byte `json:"-"`
	TokenHash string    `json:"token_hash"`
	Expiry    time.Time `json:"expiry"`         //MANDATORY*
	Role      string    `json:"role,omitempty"` //MANDATORY*
	//SecretKey string
	jwt.RegisteredClaims
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Its required because the token must be lightweight
type SaveJWTToken struct {
	// ID     int    `json:"id"`
	UserID int    `json:"user_id,omitempty"` //MANDATORY*
	Email  string `json:"email,omitempty"`   //MANDATORY* ?
	Token  string `json:"token"`
	// TokenHash []byte    `json:"-"`
	TokenHash string `json:"-"`
	// Expiry    time.Time `json:"expiry"`         //MANDATORY*
	Role string `json:"role,omitempty"` //MANDATORY*
	//SecretKey string
	jwt.RegisteredClaims
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *JWTToken) GetByJWTToken(plainText string) (*JWTToken, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, user_id, email, token, token_hash, expiry, created_at, updated_at from tokens where token = $1`

	var token JWTToken

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

	fmt.Println("query token: ", token.Token)
	fmt.Println("FROM MODEL TOKEN HASH: ", token.TokenHash)

	// La consulta retorna igual el token y el tokenhash  !! VALIDAR!!
	// Valide el orden de los parametros
	// Quiza sea a nivel de tipo de dato por byte?

	return &token, nil
}

func (j *JWTToken) GetUserForJWTToken(token JWTToken) (*User, error) {
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

// ACLARAR EL TEMA DEL TTL
// func (j *JWTToken) GenerateJWTToken(email string) (*JWTToken, error) {
// func (j *JWTToken) GenerateJWTToken(email string, userID int, role string) (*JWTToken, error) {
func (j *JWTToken) GenerateJWTToken(email string, userID int) (*JWTToken, error) {

	token := &JWTToken{
		UserID: userID,
		Email:  email,
		Role:   "trin",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: email,
			// ExpiresAt: jwt.NewNumericDate(expirationTime), // CONFIG (MAIN-INTERNAL)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()), // Se puede usar para indicar desde cuándo es válido
			Issuer:    "g3notype",
			Audience:  []string{"mis-usuarios"},
		},
	}

	secretKey := "secret"

	if secretKey == "" {
		return nil, errors.New("JWT_SECRET no está definido en las variables de entorno")
	}

	// Algoritmos de firmado

	// SIMETRIC
	// tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS512, token) // strongest hash

	fmt.Println("TOKEN:", token)
	fmt.Println("tokenObj", tokenObj)
	fmt.Println("tokenObj", tokenObj.Claims)

	// signedToken, err := token.SignedString([]byte(secretKey))
	tokenString, err := tokenObj.SignedString([]byte(secretKey))
	fmt.Println("SIGNEDTOKEN:", tokenString)
	if err != nil {
		// 500
		fmt.Println("Error al firmar el token:", err)
		return nil, err
	}

	token.Token = tokenString
	token.TokenHash = tokenString //Realmente es necesario?
	//claims.Token = signedToken

	// return &claims, nil
	return token, nil

	// Luego de tener el token se necesita obtener un string a partir de ese token (firmarlo) tokenString, signedToken
}

// Called by middleware
func (j *JWTToken) AuthenticateJWTToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no valid authorization header received")

	}

	// PARECE SER QUE NO SE ESTA USANDO ACA ESTE SECRET VALIDAR
	secretKey := "secret"

	if secretKey == "" {
		return nil, errors.New("la clave secreta no está definida")
	}

	tokenString := headerParts[1]

	fmt.Println("TACOMAN")

	token, err := j.GetByJWTToken(tokenString)
	fmt.Println("OELO", token.Token, token.TokenHash, token.Email, token.Expiry)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	// Parsear el token y extraer los claims
	// token, err := jwt.ParseWithClaims(tokenString, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {

	// 	fmt.Println("TOKEN INSIDE", token)
	// 	// Verificar que el método de firma sea el esperado
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, errors.New("método de firma no válido")
	// 	}
	// 	return []byte(secretKey), nil
	// })

	//ZOOM a este parse
	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, errors.New("unexpected signing method")
	// 	}
	// 	return []byte(j.SecretKey), nil
	// })

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	email := claims["email"].(string)
	// 	user, _ := (&User{}).FindByEmail(email)
	// 	return user, nil
	// }

	fmt.Println("token.Expiry", token.Expiry)
	fmt.Println("ExpiresAt", token.ExpiresAt) // trae 0

	if token.Expiry.Before(time.Now()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")
	}

	if token.Expiry.UTC().Before(time.Now().UTC()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")
	}

	user, err := j.GetUserForJWTToken(*token)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	fmt.Println("USER", user)

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	// // Verificar si el token es válido
	// if !token.Valid {
	// 	return nil, errors.New("token inválido")
	// }

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (j *JWTToken) ExtractJWTToken(r *http.Request) (string, error) {

	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		return "", errors.New("no authorization header received")
		// return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", ErrInvalidAuthHeader
		// return nil, errors.New("no valid authorization header received")
	}

	// PARECE SER QUE NO SE ESTA USANDO ACA ESTE SECRET VALIDAR
	secretKey := "secret"
	if secretKey == "" {
		return "", errors.New("la clave secreta no está definida")
	}

	tokenString := headerParts[1]

	return tokenString, nil
}

func (j *JWTToken) ValidJWTTokenII(plainText string) (*User, error) {

	// token, err := j.GetByJWTToken(tokenString)
	token, err := j.GetByJWTToken(plainText)
	fmt.Println("OELO", token.Token, token.TokenHash, token.Email, token.Expiry)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	fmt.Println("token.Expiry", token.Expiry)
	fmt.Println("ExpiresAt", token.ExpiresAt) // trae 0

	if token.Expiry.Before(time.Now()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")
	}

	if token.Expiry.UTC().Before(time.Now().UTC()) {
		fmt.Println("EXPIRED TOKEN")
		return nil, errors.New("expired token")
	}

	user, err := j.GetUserForJWTToken(*token)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	fmt.Println("USER", user)

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (j *JWTToken) AuthenticateJWTTokenII(r *http.Request) (*User, error) {

	token, err := j.ExtractJWTToken(r)
	if err != nil {
		return nil, err
	}

	user, err := j.ValidJWTTokenII(token)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return user, nil
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

	stmt = `insert into tokens (user_id, email, token, token_hash, expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7)`

	_, err = db.ExecContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		//[]byte(token.TokenHash), // Convertir string a bytea si es necesario
		token.RegisteredClaims.ExpiresAt.Time,
		token.RegisteredClaims.IssuedAt.Time,
		time.Now(),
		// time.Now(),
		// time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) DeleteByJWTToken(plaintText string) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from tokens where token = $1`
	_, err := db.ExecContext(ctx, stmt, plaintText)

	if err != nil {
		return err
	}
	return nil
}

// Este es el oficial
func (j *JWTToken) ValidJWTToken(plainText string) (bool, error) {

	token, err := j.GetByJWTToken(plainText)

	fmt.Println("TOKEN: ", token.Token)
	fmt.Println("DESDE MODEL-VALIDJWTTOKENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN")

	if err != nil {
		return false, errors.New("no matching token found")
	}

	// Verificar la firma del token (si es un JWT)
	// valid, err := j.VerifyJWTSignature(token.Token)
	// if err != nil || !valid {
	// 	return false, errors.New("invalid JWT token signature")
	// }

	// token, err := jwt.ParseWithClaims(tokenstring)

	// Verificar si el token está asociado a un usuario válido
	_, err = j.GetUserForJWTToken(*token)
	// _, err = t.GetUserForToken(*token)
	if err != nil {
		return false, errors.New("no matching user found")
	}

	// TR
	if token.Expiry.Before(time.Now()) {
		return false, errors.New("expired token")
	}

	return true, nil

}

// Debio haber sido escrito como handler, ya esta corregido
// Debe llevarse esta logica a valodJWTToken
// func (j *JWTToken) ValidateJWTToken(tokenString string, secretKey string) (*JWTToken, error) {
// 	// Verificar si la clave secreta está vacía
// 	if secretKey == "" {
// 		return nil, errors.New("la clave secreta no está definida")
// 	}
// 	fmt.Println("SECRETKEY", secretKey)

// 	fmt.Println("TokenSTRING", tokenString)

// 	// Parsear el token y extraer los claims
// 	token, err := jwt.ParseWithClaims(tokenString, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {

// 		fmt.Println("TOKEN INSIDE", token)
// 		// Verificar que el método de firma sea el esperado
// 		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		// 	return nil, errors.New("método de firma no válido")
// 		// }
// 		if token.Method != jwt.SigningMethodHS512 {
// 			return nil, errors.New("método de firma no válido")
// 		}
// 		return []byte(secretKey), nil
// 	})
// 	fmt.Println("TOKEN", token)

// 	if err != nil {
// 		fmt.Println("ERROR AL VALIDAR EL TOKEN")
// 		return nil, fmt.Errorf("error al validar el token: %w", err)
// 	}

// 	// Verificar si el token es válido
// 	if !token.Valid {
// 		return nil, errors.New("token inválido")
// 	}

// 	// Extraer los claims
// 	claims, ok := token.Claims.(*JWTToken)
// 	if !ok {
// 		return nil, errors.New("no se pudieron extraer los claims")
// 		fmt.Println("ERROR AL EXTRAER CLAIMS")
// 	}

// 	fmt.Println(claims.Role)

// 	return claims, nil
// }
