package main

import (
	models "jwt-rbac-products-API/internal"
	"net/http"

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

	return mux
}
