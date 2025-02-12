package models

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Eliptic curves ES256

// Generar un JWT con ES256
func (j *JWTToken) GenerateAsimJWTToken(email string, userID int) (*JWTToken, error) {
	// Cargar la clave privada EC desde el archivo
	privateKey, err := j.readPrivateKey("ec_private.pem")
	if err != nil {
		return nil, err
	}

	// Crear la estructura del token
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

// Validar y decodificar el token
func (j *JWTToken) ValidateAsimJWTToken(tokenString string) (*JWTToken, error) {
	// Cargar la clave pública EC
	publicKey, err := j.readPublicKey("ec_public.pem")
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
