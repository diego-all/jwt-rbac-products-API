package config

import (
	"errors"
	"time"
)

// original
// type Config struct {
// 	port           int
// 	certPath       string
// 	keyPath        string
// 	jwt_secret     string
// 	hash_cost      int // bcrypt
// 	token_duration time.Duration
// 	jwt_PrivateKey string
// }

type Config struct {
	Port          int
	CertPath      string
	KeyPath       string
	JWTPrivateKey string
	JWTSecret     string
	HashCost      int
	TokenDuration time.Duration
	DatabaseDSN   string
}

// WC
// func GetEnv(name string) string {
// 	return os.Getenv(name)
// }

// LoadConfig carga la configuración desde variables de entorno o valores por defecto
// Posibilidad de error: si la configuración se carga desde archivos, variables de entorno o cualquier fuente externa que pueda fallar.
func LoadConfig() (*Config, error) {

	port := 9090
	certPath := "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.pem"
	keyPath := "/home/diegoall/MAESTRIA_ING/domain-model/products-API/cmd/api/server.key"
	jwtPrivateKey := "/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_private.pem"
	jwtSecret := "/home/diegoall/MAESTRIA_ING/OAuth/jwt-rbac-products-API/cmd/api/ec_private.pem"
	hashCost := 8
	tokenDuration := 24 * time.Hour
	databaseDSN := "host=localhost port=54325 user=postgres password=password dbname=e_commerce sslmode=disable timezone=UTC connect_timeout=5"

	// port := 9090
	// certPath := os.Getenv("CERT_PATH")
	// keyPath := os.Getenv("KEY_PATH")
	// jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
	// jwtSecret := os.Getenv("JWT_SECRET")
	// hashCost := 8
	// tokenDuration := 24 * time.Hour
	// databaseDSN := os.Getenv("DATABASE_DSN")

	// Validar que las variables críticas estén definidas
	if certPath == "" || keyPath == "" || jwtPrivateKey == "" || jwtSecret == "" || databaseDSN == "" {
		return nil, errors.New("faltan variables de entorno críticas en la configuración")
	}

	return &Config{
		Port:          port,
		CertPath:      certPath,
		KeyPath:       keyPath,
		JWTPrivateKey: jwtPrivateKey,
		JWTSecret:     jwtSecret,
		HashCost:      hashCost,
		TokenDuration: tokenDuration,
		DatabaseDSN:   databaseDSN,
	}, nil
}

// # WC
// type Config struct {
// 	Db      Database
// 	Server  HTTPS
// 	Redis   Redis
// 	Limiter Limiter
// 	Logger  Logger
// 	Crypto  Crypto
// 	Lockout Lockout
// 	File    File
// 	Mongo   Mongo
// 	Env     Env
// }
