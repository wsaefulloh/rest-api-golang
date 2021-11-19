package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/rest-api-go/configs/db"
)

func New() *mux.Router {
	mainRutes := mux.NewRouter()

	//inisialisasi endpoint
	mainRutes.HandleFunc("/", simpleHandlers).Methods("GET")

	//inisialisasi dbms
	dbms, _ := db.New()

	UserRoute(mainRutes, dbms)
	ProductRoute(mainRutes, dbms)
	CategoryRoute(mainRutes, dbms)
	HistoryRoute(mainRutes, dbms)
	return mainRutes
}

func simpleHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
