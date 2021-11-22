package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wsaefulloh/rest-api-go/helpers"
	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type auth struct {
	rp repos.RepoUser
}

type userToken struct {
	Token string `json:"token"`
}

func Auth(rps repos.RepoUser) *auth {
	return &auth{rps}
}

func (at auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	validate := validator.New()
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}
	err = validate.Struct(body)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		helpers.Respone(w, validationErrors.Error(), 400, true)
		return
	}

	pass, err := at.rp.GetPass(body.UserName)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	if !helpers.CheckPassword(pass, body.Password) {
		helpers.Respone(w, "Password Salah", 401, true)
		return
	}

	email, err := at.rp.GetEmail(body.UserName)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	if email != body.Email {
		helpers.Respone(w, "Email Salah", 401, true)
		return
	}

	role, err := at.rp.GetRole(body.UserName)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	tokens := helpers.NewToken(body.UserName, role)
	fmt.Println(role)
	theTokens, err := tokens.Create()

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	helpers.Respone(w, userToken{Token: theTokens}, 200, false)
}

func (at auth) OpenToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("token_auth")
	validation, value := helpers.CheckToken(tokenString)

	if validation != true {
		helpers.Respone(w, value, 500, true)
		return
	}

	helpers.Respone(w, value, 200, false)
}
