package routers

import (
	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
)

func ProductRoute(r *mux.Router, cr *controllers.Products) {
	route := r.PathPrefix("/products").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
