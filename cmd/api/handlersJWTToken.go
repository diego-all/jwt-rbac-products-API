package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *application) LoginJWT(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("LOGINJWT")

	// look up the user by email
	// Security Advice: Change output messages to avoid enumerate users attack
	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// fmt.Println("USER:", user)

	// validate the user's password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	fmt.Println("VALID PASSWORD", validPassword)

	// we have a valid user, so generate a token
	token, err := app.models.JWTToken.GenerateJWTToken(creds.UserName, user.ID) // userID ?? DURATION FROM MAIN
	//token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	fmt.Println("TOKEN", token)

	// save it to the database
	err = app.models.JWTToken.InsertJWT(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

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

func (app *application) ValidateJWTToken(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	valid := false
	valid, _ = app.models.JWTToken.ValidJWTToken(requestPayload.Token)
	// valid, _ = app.models.Token.ValidToken(requestPayload.Token)

	fmt.Println("VALID IN VALIDATEJWTTOKEN", valid)

	payload := jsonResponse{
		Error: false,
		Data:  valid,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

	//Get user by email

}
