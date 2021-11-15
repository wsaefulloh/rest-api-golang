package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type Products struct {
	Rp repos.InitRepoProduct
}

func NewProduct(rps repos.InitRepoProduct) *Products {
	return &Products{rps}
}

func (pro *Products) GetAll(w http.ResponseWriter, r *http.Request) {

	data, err := pro.Rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) Add(w http.ResponseWriter, r *http.Request) {
	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	pro.Rp.Save(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil disimpan"))
}

func (pro *Products) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	err := pro.Rp.Remove(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil dihapus"))
}

func (pro *Products) Update(w http.ResponseWriter, r *http.Request) {
	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	vars := mux.Vars(r)

	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	pro.Rp.Edit(data, vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil diedit"))
}
