package main

import (
	models "jwt-rbac-products-API/internal"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/health", app.Health)

	// // Product
	// mux.Post("/products", app.CreateProduct)
	// mux.Get("/products/get/{id}", app.GetProduct)
	// mux.Put("/products/update/{id}", app.UpdateProduct)
	// mux.Get("/products/all", app.AllProducts)
	// mux.Delete("/products/delete/{id}", app.DeleteProduct)

	// mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)
	mux.Post("/validate-token", app.ValidateToken)

	//NUEVOS
	// mux.Post("/users/signup", app.SignUpJWT)
	mux.Post("/users/loginjwt", app.LoginJWT)
	// mux.Get("/test-generate-jwt-token", app)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users models.User
		all, err := users.GetAll()
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    envelope{"users": all},
		}

		// app.writeJSON(w, http.StatusOK, all)
		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = models.User{
			Email:     "diego@diego.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}

		app.infoLog.Println("Adding user...")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("Got back id of", id)
		newUser, _ := app.models.User.GetOne(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {

		token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		// token.Email = "admin@example.com"
		token.Email = "diego@diego.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(1, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		user, err := app.models.User.GetOne(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}
		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-jwt-token", func(w http.ResponseWriter, r *http.Request) {

		token, err := app.models.JWTToken.GenerateJWTToken("diego@diego.com", 1)
		// token, err := app.models.User.Token.GenerateToken(1, 60*time.Minute)
		// token, err := app.models.User.Token.GenerateToken(1, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		user, err := app.models.User.GetOne(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.InsertJWT(*token, *user)
		// err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}
		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		var payload jsonResponse
		payload.Error = false
		payload.Data = valid
		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-jwt-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")

		valid, err := app.models.JWTToken.ValidJWTToken(tokenToValidate)
		// valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		var payload jsonResponse
		payload.Error = false
		payload.Data = valid
		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-generate-jwt-token", func(w http.ResponseWriter, r *http.Request) {

		token, err := app.models.JWTToken.GenerateJWTToken("diego@diego.com", 1)
		// token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute) //Enviaban una duracion
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		// token.Email = "admin@example.com"
		token.Email = "diego@diego.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.AuthTokenMiddleware)

		// mux.Post("/users", app.AllUsers)
		// mux.Post("/users/save", app.EditUser)
		// mux.Post("/users/get/{id}", app.GetUser)
		// mux.Post("/users/delete", app.DeleteUser)
		// mux.Post("/log-user-out/{id}", app.LogUserOutAndSetInactive)

		mux.Post("/products", app.CreateProduct)
		mux.Get("/products/get/{id}", app.GetProduct)
		mux.Put("/products/update/{id}", app.UpdateProduct)
		mux.Get("/products/all", app.AllProducts)
		mux.Delete("/products/delete/{id}", app.DeleteProduct)

		mux.Post("/foo", func(w http.ResponseWriter, r *http.Request) {
			payload := jsonResponse{
				Error:   false,
				Message: "bar",
			}
			app.writeJSON(w, http.StatusOK, payload)
		})

	})

	return mux
}
