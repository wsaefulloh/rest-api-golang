package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/rest-api-go/helpers"
	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type categories struct {
	rp repos.RepoCategory
}

func NewCategory(rps repos.RepoCategory) *categories {
	return &categories{rps}
}

func (cate *categories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := cate.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (cate *categories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateCategory()
	data.Name = body.Name

	cate.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	helpers.Respone(w, "Categori berhasil ditambahkan", 201, false)
}

func (cate *categories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := cate.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	helpers.Respone(w, "Categori berhasil dihapus", 200, false)
}

func (cate *categories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)

	data := models.CreateCategory()
	data.Name = body.Name

	cate.rp.Edit(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	helpers.Respone(w, "Categori berhasil diupdate", 200, false)
}

func (cate *categories) GetCategory(id int) bool {
	value, _ := cate.rp.GetCategory(id)

	if value == true {
		return true
	}

	return false
}
