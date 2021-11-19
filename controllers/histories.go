package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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
	w.Header().Set("Content-Type", "application/json")
	data, err := histo.rp.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&data)
}

func (histo *histories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data := models.CreateHistory()
	data.Invoice = body.Invoice
	data.Cashier = body.Cashier
	data.Order = body.Order
	data.Count = body.Count
	data.Total = 0

	histo.rp.Save(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil disimpan"))
}

func (histo *histories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := histo.rp.Remove(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil dihapus"))
}

func (histo *histories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Data berhasil diedit"))
}
