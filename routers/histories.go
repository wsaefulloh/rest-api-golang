package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func HistoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewHistory(db)
	cr := controllers.NewHistory(repo)
	route := r.PathPrefix("/history").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("", cr.GetAll).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
