# Para el middleware implementar el patron familia (Strategy)



1. Definir la interfaz común
go
Copiar
Editar
type Authenticator interface {
    Authenticate(r *http.Request) (*User, error)
}
2. Implementar cada tipo de autenticación como una estrategia
Token personalizado
go
Copiar
Editar
type TokenAuthenticator struct {
    TokenModel *Token
}

func (t *TokenAuthenticator) Authenticate(r *http.Request) (*User, error) {
    return t.TokenModel.AuthenticateToken(r)
}
JWT clásico
go
Copiar
Editar
type JWTAuthenticator struct {
    JWTModel *JWTToken
}

func (j *JWTAuthenticator) Authenticate(r *http.Request) (*User, error) {
    return j.JWTModel.AuthenticateJWTToken(r)
}
JWT asimétrico
go
Copiar
Editar
type JWTAsimAuthenticator struct {
    JWTModel *JWTToken
}

func (j *JWTAsimAuthenticator) Authenticate(r *http.Request) (*User, error) {
    return j.JWTModel.AuthenticateAsimJWTToken(r)
}


func (app *application) AuthTokenMiddleware(authenticator Authenticator) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            _, err := authenticator.Authenticate(r)
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
}


🔶 ¿Qué es el Patrón Estrategia?
El Patrón Estrategia te permite definir una familia de algoritmos (en tu caso, distintas formas de autenticar un token), encapsular cada uno y hacerlos intercambiables. Este patrón permite que el algoritmo varíe independientemente de los clientes que lo usan (como tu middleware).


    ⚙️ Beneficios de este enfoque:
    ✅ Evitas duplicación de código.
    ✅ Puedes agregar nuevos tipos de autenticación fácilmente sin modificar el middleware.
    ✅ Separas las responsabilidades: el middleware solo se preocupa por autenticar, no cómo.
    ✅ Fácil de mantener, probar y escalar.
