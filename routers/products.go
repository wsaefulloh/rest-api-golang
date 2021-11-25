package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/middleware"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func ProductRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewProduct(db)
	cr := controllers.NewProduct(repo)
	route := r.PathPrefix("/products").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/date/desc", cr.GetbyDateDESC).Methods(http.MethodGet)
	route.HandleFunc("/date/asc", cr.GetbyDateASC).Methods(http.MethodGet)
	route.HandleFunc("/price/desc", cr.GetbyPriceDESC).Methods(http.MethodGet)
	route.HandleFunc("/price/asc", cr.GetbyPriceASC).Methods(http.MethodGet)
	route.HandleFunc("/category", cr.GetbyCategory).Methods(http.MethodGet)
	route.HandleFunc("/search/name", cr.SearchbyName).Methods(http.MethodGet)
	route.HandleFunc("/search/category", cr.SearchbyCategory).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
}
