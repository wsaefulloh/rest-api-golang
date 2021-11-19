package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func CategoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewCategory(db)
	cr := controllers.NewCategory(repo)
	route := r.PathPrefix("/category").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
