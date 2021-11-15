package routers

import (
	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
)

func CategoryRoute(r *mux.Router, cr *controllers.Categories) {
	route := r.PathPrefix("/category").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
