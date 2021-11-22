package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func AuthRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewUsers(db)
	cr := controllers.Auth(repo)

	route := r.PathPrefix("/auth").Subrouter()
	route.HandleFunc("/", cr.Login).Methods(http.MethodPost)
	route.HandleFunc("/chek", cr.OpenToken).Methods(http.MethodGet)
}
