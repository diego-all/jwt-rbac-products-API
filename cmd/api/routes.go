package main

import (
	models "jwt-rbac-products-API/internal"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/health", app.Health)

	// Product
	mux.Post("/products", app.CreateProduct)
	mux.Get("/products/get/{id}", app.GetProduct)
	mux.Put("/products/update/{id}", app.UpdateProduct)
	mux.Get("/products/all", app.AllProducts)
	mux.Delete("/products/delete/{id}", app.DeleteProduct)

	// mux.Get("/users/login", app.Login)
	// mux.Post("/users/login", app.Login)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users models.User
		all, err := users.GetAll()
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		app.writeJSON(w, http.StatusOK, all)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = models.User{
			Email:     "you@trinis.com",
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

	return mux
}
