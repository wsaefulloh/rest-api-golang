package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func UserRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewUsers(db)
	cr := controllers.New(repo)
	route := r.PathPrefix("/users").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
}
