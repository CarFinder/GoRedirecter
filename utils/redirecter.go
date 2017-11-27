package utils

import (
	"log"
	"net/http"
)

func sendRequest(url string, referHeader string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("An error occured while trying to generate a request")
	}
	request.Host = referHeader
	client := http.Client{}
	response, err := client.Do(request)
	return response, err
}
