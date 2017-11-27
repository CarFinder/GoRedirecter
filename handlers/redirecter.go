package handlers

import (
	"fmt"
	"net/http"
)

// GetImage receives an image from another server and returns it as a response
func GetImage(w http.ResponseWriter, r *http.Request) {
	imagePath := r.URL.Query().Get("url")

	if imagePath != "" {
		fmt.Print(imagePath)
	}
}
