package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/rest-api-go/helpers"
	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type products struct {
	rp repos.RepoProduct
}

func NewProduct(rps repos.RepoProduct) *products {
	return &products{rps}
}

func (pro *products) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("token_auth")

	validation, value := helpers.CheckToken(tokenString)

	if validation != true {
		helpers.Respone(w, "Login dulu", 500, true)
		return
	}

	if value != "admin" {
		helpers.Respone(w, "Maaf anda bukan admin", 400, true)
		return
	}

	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	pro.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	helpers.Respone(w, "Product berhasil disimpan", 201, false)
}

func (pro *products) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := pro.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	helpers.Respone(w, "Product berhasil dihapus", 200, false)
}

func (pro *products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)
	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	pro.rp.Edit(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	helpers.Respone(w, "Data berhasil diupdate", 200, false)
}

func (pro *products) SearchbyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_prod := strings.Join(vars["name"], " ")
	data, err := pro.rp.SearchProductName(name_prod)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// fmt.Println(name_prod)
	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) SearchbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_category := strings.Join(vars["category"], " ")
	data, err := pro.rp.SearchProductCategory(name_category)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// fmt.Println(name_category)
	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) GetbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindbyCategory()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	json.NewEncoder(w).Encode(&data)
}

func (pro *products) GetbyDateASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindbyDateASC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) GetbyDateDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindbyDateDESC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) GetbyPriceDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindbyPriceDESC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (pro *products) GetbyPriceASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindbyPriceASC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}
