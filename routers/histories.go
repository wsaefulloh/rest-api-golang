package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/middleware"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func HistoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewHistory(db)
	cr := controllers.NewHistory(repo)
	route := r.PathPrefix("/history").Subrouter()
	route.HandleFunc("/", middleware.Validate(cr.GetAll)).Methods(http.MethodGet)
	route.HandleFunc("", middleware.Validate(cr.GetAll)).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
}
