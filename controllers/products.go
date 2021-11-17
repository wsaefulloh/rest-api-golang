package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := pro.Rp.Remove(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil dihapus"))
}

func (pro *Products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func (pro *Products) SearchbyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_prod := strings.Join(vars["name"], " ")
	data, err := pro.Rp.SearchProductName(name_prod)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(name_prod)
	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) SearchbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_category := strings.Join(vars["category"], " ")
	data, err := pro.Rp.SearchProductCategory(name_category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(name_category)
	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) GetbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindbyCategory()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) GetbyDateASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindbyDateASC()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) GetbyDateDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindbyDateDESC()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) GetbyPriceDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindbyPriceDESC()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *Products) GetbyPriceASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.Rp.FindbyPriceASC()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}
