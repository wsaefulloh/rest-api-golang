package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/controllers"
	"github.com/wsaefulloh/rest-api-go/repos"
)

func New() *mux.Router {
	mainRutes := mux.NewRouter()

	//inisialisasi endpoint
	mainRutes.HandleFunc("/", simpleHandlers).Methods("GET")

	//inisialisasi repos
	userRep := repos.New()
	users := controllers.Users{Rp: *userRep}

	prodRep := repos.NewProduct()
	products := controllers.Products{Rp: *prodRep}

	cateRep := repos.NewCategory()
	category := controllers.Categories{Rp: *cateRep}

	UserRoute(mainRutes, &users)
	ProductRoute(mainRutes, &products)
	CategoryRoute(mainRutes, &category)
	return mainRutes
}

func simpleHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
