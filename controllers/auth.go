package controllers

import (
	"encoding/json"
	"net/http"

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
	w.Header().Set("Content-Type", "application/json")
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	pass, err := at.rp.GetPass(body.UserName)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	if !helpers.CheckPassword(pass, body.Password) {
		// w.Write([]byte("Password Salah"))
		helpers.Respone(w, "Password Salah", 401, true)
		return
	}

	tokens := helpers.NewToken(body.UserName)
	theTokens, err := tokens.Create()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	helpers.Respone(w, userToken{Token: theTokens}, 200, false)
}