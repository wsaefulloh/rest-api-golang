package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wsaefulloh/rest-api-go/routers"
)

func main() {
	mainRutes := routers.New()

	fmt.Println("Server running on port 9000")
	err := http.ListenAndServe(":9000", mainRutes)
	if err != nil {
		log.Fatal("Error API")
	}
}
