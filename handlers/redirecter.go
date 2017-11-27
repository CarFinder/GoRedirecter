package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CarFinder/GoRedirecter/utils"
)

// GetImage receives an image from another server and returns it as a response
func GetImage(w http.ResponseWriter, r *http.Request) {
	imagePath := r.URL.Query().Get("url")
	hostName := r.URL.Query().Get("host")

	if imagePath == "" || hostName == "" {
		http.Error(w, "Incorrect parameters: image url or host name expected", http.StatusBadRequest)
	}

	resp, err := utils.SendRequest(imagePath, hostName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		http.Error(w, readErr.Error(), http.StatusUnprocessableEntity)
	}

	resp.Body.Close()

	w.Header().Set("Content-Length", fmt.Sprint(resp.ContentLength))
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Write(body)
}
