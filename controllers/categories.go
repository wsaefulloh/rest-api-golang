package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type Categories struct {
	Rp repos.InitRepoCategory
}

func NewCategory(rps repos.InitRepoCategory) *Categories {
	return &Categories{rps}
}

func (cate *Categories) GetAll(w http.ResponseWriter, r *http.Request) {

	data, err := cate.Rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (cate *Categories) Add(w http.ResponseWriter, r *http.Request) {
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data := models.CreateCategory()
	data.Name = body.Name

	cate.Rp.Save(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil disimpan"))
}

func (cate *Categories) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	err := cate.Rp.Remove(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil dihapus"))
}

func (cate *Categories) Update(w http.ResponseWriter, r *http.Request) {
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	vars := mux.Vars(r)

	data := models.CreateCategory()
	data.Name = body.Name

	cate.Rp.Edit(data, vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil diedit"))
}
