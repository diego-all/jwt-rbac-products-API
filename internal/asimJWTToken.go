package models

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// func (t *JWTToken) GetByAsimJWTToken(plainText string) (*JWTToken, error) {
// }

// Generar un JWT con ES256 (Eliptic curves ES256)
func (j *JWTToken) GenerateAsimJWTToken(email string, userID int) (*JWTToken, error) {
	// Cargar la clave privada EC desde el archivo
	privateKey, err := j.readPrivateKey("/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_private.pem")
	if err != nil {
		return nil, err
	}

	token := &JWTToken{
		UserID: userID,
		Email:  email,
		Role:   "trin",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira en 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()), // Desde cuándo es válido
			Issuer:    "g3notype",
			Audience:  []string{"mis-usuarios"},
		},
	}

	// Crear el objeto JWT con ES256
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodES256, token)

	// Firmar el token con la clave privada
	tokenString, err := tokenObj.SignedString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("error al firmar el token: %w", err)
	}

	// Asignar el token firmado
	token.Token = tokenString
	token.TokenHash = tokenString

	return token, nil
}

// Called by middleware
func (j *JWTToken) AuthenticateAsimJWTToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no valid authorization header received")

	}

	tokenString := headerParts[1]

	token, err := j.GetByJWTToken(tokenString)
	// fmt.Println("OELO", token.Token, token.Email, token.Expiry) //'QUE PASA CON TOKEN HASH?
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	// YA HAY UNA FUNCION PARA VALIDAR EL TOKEN ValidateAsimJWTToken(), ESTA SIENDO LLAMADA DESDE LOS HANDLERS EN ROUTE.go

	// Validación de la firma del JWT
	// Aquí deberías proporcionar la clave pública o el secreto que usas para firmar el token
	// Usando el método Parse para verificar la firma
	// parsedToken, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	// Reemplaza esto con la clave secreta o clave pública adecuada
	// 	// Aquí estoy utilizando un ejemplo con una clave secreta
	// 	// Si estás usando RS256 o algún otro algoritmo, necesitarás la clave pública adecuada
	// 	return []byte("secret"), nil
	// })

	// parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	// Verificar que el algoritmo sea ES256
	// 	if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
	// 		return nil, errors.New("método de firma no válido")
	// 	}
	// 	return publicKey, nil
	// })

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Verificar si el token es válido
	// if !parsedToken.Valid {
	// 	return nil, errors.New("invalid token")
	// }

	fmt.Println("ACA VOY")

	// fmt.Println("TOKEN ID", token.ID)
	// fmt.Println("USER ID", token.UserID)       // trae 0
	// fmt.Println("EMAIL", token.Email)          // trae 0
	// fmt.Println("TOKEN", token.Token)          // trae 0
	// fmt.Println("TOKEN HASH", token.TokenHash) // trae 0
	// fmt.Println("EXPIRY", token.Expiry)        // trae 0
	// fmt.Println("CREATED AT", token.CreatedAt) // trae 0
	// fmt.Println("UPDATED AT", token.UpdatedAt) // trae 0
	// EL TOKEN Y TOKENHASH DESDE EL MODELO VIEN IGUALES (VALIDAR)

	// fmt.Println("token.Expiry", token.Expiry)
	// fmt.Println("IssuedAt", token.RegisteredClaims.IssuedAt)   // trae 0
	// fmt.Println("ExpiresAt", token.RegisteredClaims.ExpiresAt) // trae 0
	// fmt.Println("TOKEN COMPLETO", token)
	// fmt.Println("Registered Claims", token.RegisteredClaims)
	// fmt.Println("Role", token.Role)

	//PILAS CON LA ZONA HORARIA RECORDAR EN LA DB EXPIRY WITH TIME ZONE (timestamptz)
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

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	// En token sencillo utilizan una funcion (ValidToken) del receiver Token
	// Internamente validan la fecha de expiracion del token.
	// valid, err := app.models.Token.ValidToken(tokenToValidate)
	// VALIDAR
	if err := token.Valid(); err != nil {
		return nil, errors.New("token inválido")
	}

	if err != nil {
		return nil, err
	}

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	email := claims["email"].(string)
	// 	user, _ := (&User{}).FindByEmail(email)
	// 	return user, nil
	// }

	fmt.Println("USER", user)

	return user, nil
	// return nil, errors.New("invalid token")
}

// func (j *JWTToken) ValidateJWTToken(tokenString string, secretKey string) (*JWTToken, error) {
func (j *JWTToken) ValidateAsimJWTToken(tokenString string) (*JWTToken, error) {

	// La lectura de la llave publica deberia ser en package main
	publicKey, err := j.readPublicKey("/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_public.pem")
	if err != nil {
		return nil, err
	}

	// Parsear y validar el token
	token, err := jwt.ParseWithClaims(tokenString, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea ES256
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("método de firma no válido")
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	// Convertir el token a la estructura personalizada
	claims, ok := token.Claims.(*JWTToken)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

// Leer la clave pública desde un archivo PEM
func (j *JWTToken) readPublicKey(filepath string) (*ecdsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo clave pública: %w", err)
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, fmt.Errorf("error parseando clave pública: %w", err)
	}

	return publicKey, nil
}

// Función para leer la clave privada desde un archivo PEM
func (j *JWTToken) readPrivateKey(filepath string) (*ecdsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo clave privada: %w", err)
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, fmt.Errorf("error parseando clave privada: %w", err)
	}

	return privateKey, nil
}

// Convertir el token a la estructura personalizada
// claims, ok := token.Claims.(*JWTToken)
// if !ok || !token.Valid {
// 	return nil, errors.New("token inválido")
// }
