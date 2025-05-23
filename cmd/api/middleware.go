package main

import "net/http"

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// _, err := app.models.Token.AuthenticateTokenII(r)
		// _, err := app.models.Token.AuthenticateToken(r)
		_, err := app.models.JWTToken.AuthenticateJWTTokenII(r)
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
