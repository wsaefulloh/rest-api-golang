package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func ProductRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewProduct(db)
	cr := controllers.NewProduct(repo)
	route := r.PathPrefix("/products").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods("GET")
	route.HandleFunc("", cr.GetAll).Methods("GET")
	route.HandleFunc("/date/desc", cr.GetbyDateDESC).Methods("GET")
	route.HandleFunc("/date/asc", cr.GetbyDateASC).Methods("GET")
	route.HandleFunc("/price/desc", cr.GetbyPriceDESC).Methods("GET")
	route.HandleFunc("/price/asc", cr.GetbyPriceASC).Methods("GET")
	route.HandleFunc("/category", cr.GetbyCategory).Methods("GET")
	route.HandleFunc("/search/name", cr.SearchbyName).Methods("GET")
	route.HandleFunc("/search/category", cr.SearchbyCategory).Methods("GET")
	route.HandleFunc("/", cr.Add).Methods("POST")
	route.HandleFunc("", cr.Add).Methods("POST")
	route.HandleFunc("/{id}", cr.Delete).Methods("DELETE")
	route.HandleFunc("/{id}", cr.Update).Methods("PUT")
}
