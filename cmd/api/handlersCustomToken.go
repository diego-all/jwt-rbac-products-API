package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// "0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z" in response from Insert

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
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

	// TODO authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// look up the user by email
	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// validate the user's password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// make sure user is active
	// if user.Active == 0 {
	// 	app.errorJSON(w, errors.New("user is not active"))
	// 	return
	// }

	// we have a valid user, so generate a token
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// It's necessary retrieve the data that isn't in the struct from the database. (token.CreatedAt, token.UpdatedAt)
	// Maybe modify the Insert function

	// Pointers (GenerateToken returns *models.token)
	// cannot use *token (variable of type models.Token) as *models.Token value in argument to app.models.Token.InsertReturningcompilerIncomp

	// save it to the database
	// En Go, no puedes asignar directamente un valor a un puntero sin hacer una conversión explícita.

	tokenReturning, err := app.models.Token.InsertReturning(*token, *user)
	// err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	fmt.Println("tokenReturning tokenReturning tokenReturning ...", tokenReturning)

	fmt.Println("createdAt", tokenReturning.CreatedAt)
	fmt.Println("updatedAt", tokenReturning.UpdatedAt)
	fmt.Println(tokenReturning.Email)
	fmt.Println(tokenReturning.ID)
	fmt.Println(tokenReturning.TokenHash)
	fmt.Println(tokenReturning.UserID)
	fmt.Println(tokenReturning.Expiry)
	fmt.Println(tokenReturning.Token)

	// time.Sleep(5 * time.Second)
	// fmt.Println("token.Token", token.Token)

	// tkn, err := app.models.Token.GetDataForUpdateHandlerToken(token.Token)
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }
	// fmt.Println("TOKENNNN AFTER RETRIEVED", tkn)

	// Option I (bad)
	// token.CreatedAt = tkn.CreatedAt
	// token.UpdatedAt = tkn.UpdatedAt

	token.CreatedAt = tokenReturning.CreatedAt
	token.UpdatedAt = tokenReturning.UpdatedAt

	// Option II (bad)
	// Overwrite the generated token with a new empty instance
	// undefined: modelscompilerUndeclaredName  (Interfaces required)
	// token = &models.Token{
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	//fmt.Println("TOKEN TOKEN TOKEN TOKEN TOKEN:", token.CreatedAt)

	// send back a response
	payload = jsonResponse{
		Error:   false,
		Message: "logged in",
		Data:    envelope{"token": token, "user": user},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}

// Parce ser de prueba
func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	valid := false
	valid, _ = app.models.Token.ValidToken(requestPayload.Token)

	payload := jsonResponse{
		Error: false,
		Data:  valid,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

	//Get user by email

}
