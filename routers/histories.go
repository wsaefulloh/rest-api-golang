package routers

import (
	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
)

func HistoryRoute(r *mux.Router, cr *controllers.Histories) {
	route := r.PathPrefix("/history").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
