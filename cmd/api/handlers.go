package main

import (
	"errors"
	"fmt"
	models "jwt-rbac-products-API/internal"
	"net/http"

	"github.com/segmentio/ksuid"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *application) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// TODO
func (app *application) SignUp(w http.ResponseWriter, r *http.Request) {

	type credentials struct {
		UserName  string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// Pending, maybe will be located in model user Insert function
	id, err := ksuid.NewRandom()
	if err != nil {
		app.errorJSON(w, errors.New("Internal server error"))
		return
	}
	fmt.Println("ID GENERADO", id)

	// hash and save the password in the database
	// HASH_COST = 8

	// The password is bee hashed in model, maybe will be required here.

	var user = models.User{
		Email: creds.UserName,
		// Password: string(hashedPassword),
		Password: creds.Password,
		// ID:       id,
		// ID: int(id),
		LastName:  creds.LastName,
		FirstName: creds.FirstName,
	}

	// doubt pointer
	_, err = app.models.User.Insert(user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "logged out",
	}

	app.infoLog.Fatal("sdsdsdsd")

	_ = app.writeJSON(w, http.StatusOK, payload)
}
