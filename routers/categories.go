package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func CategoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewCategory(db)
	cr := controllers.NewCategory(repo)
	route := r.PathPrefix("/category").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
	route.HandleFunc("", cr.Add).Methods(http.MethodPost)
	route.HandleFunc("/{id}", cr.Delete).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", cr.Update).Methods(http.MethodPut)
}
