package main

import (
	"log"
	"net/http"

	"github.com/CarFinder/GoRedirecter/handlers"
)

func main() {
	http.HandleFunc("/api/image", handlers.GetImage)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("An error occured while trying to serve the app: ", err)
	}
}
