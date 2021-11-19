package controllers

import (
	"encoding/json"
	"net/http"

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

	data, err := us.rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (us *users) Add(w http.ResponseWriter, r *http.Request) {
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data := models.CreateUser()
	data.Name = body.Name
	data.UserName = body.UserName
	data.Email = body.Email
	data.Password = body.Password

	us.rp.Save(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil disimpan"))
}
