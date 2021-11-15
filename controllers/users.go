package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type Users struct {
	Rp repos.InitRepo
}

func New(rps repos.InitRepo) *Users {
	return &Users{rps}
}

func (us *Users) GetAll(w http.ResponseWriter, r *http.Request) {

	data, err := us.Rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (us *Users) Add(w http.ResponseWriter, r *http.Request) {
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

	us.Rp.Save(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil disimpan"))
}
