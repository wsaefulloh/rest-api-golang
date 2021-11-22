package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wsaefulloh/rest-api-go/helpers"
	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type users struct {
	rp repos.RepoUser
}

func New(rps repos.RepoUser) *users {
	return &users{rps}
}

func (us *users) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	data, err := us.rp.FindAll()

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	helpers.Respone(w, &data, 200, false)
}

func (us *users) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateUser()
	data.Name = body.Name
	data.UserName = body.UserName
	data.Email = body.Email
	newPass, errPass := helpers.HashPassword(body.Password)
	if errPass != nil {
		// fmt.Println("gagal hashing")
		helpers.Respone(w, err.Error(), 500, true)
		return
	}
	data.Password = newPass

	us.rp.Save(data)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	helpers.Respone(w, "Berhasil Save", 201, false)
}
