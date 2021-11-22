package main

import (
	"log"
	"os"

	"github.com/wsaefulloh/rest-api-go/configs/command"
)

func main() {
	err := command.Run(os.Args[1:])
	if err != nil {
		log.Fatal("Unable to run app")
	}
}
