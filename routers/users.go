package routers

import (
	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
)

func UserRoute(r *mux.Router, cr *controllers.Users) {
	route := r.PathPrefix("/users").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
}
