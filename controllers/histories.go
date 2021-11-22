package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/rest-api-go/helpers"
	"github.com/wsaefulloh/rest-api-go/models"
	"github.com/wsaefulloh/rest-api-go/repos"
)

type histories struct {
	rp repos.RepoHistory
}

func NewHistory(rps repos.RepoHistory) *histories {
	return &histories{rps}
}

func (histo *histories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := histo.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	helpers.Respone(w, &data, 200, false)
}

func (histo *histories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateHistory()
	data.Invoice = body.Invoice
	data.Cashier = body.Cashier
	data.Order = body.Order
	data.Count = body.Count
	data.Total = 0

	histo.rp.Save(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	helpers.Respone(w, "History berhasil ditambahkan", 201, false)
}

func (histo *histories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := histo.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	helpers.Respone(w, "History berhasil dihapus", 200, false)
}

func (histo *histories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)

	data := models.CreateHistory()
	data.Invoice = body.Invoice
	data.Cashier = body.Cashier
	data.Order = body.Order
	data.Count = body.Count
	data.Total = 0

	histo.rp.Edit(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	helpers.Respone(w, "History berhasil diupdate", 200, false)
}
