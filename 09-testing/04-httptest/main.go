// This lets you test a handler function
// by simulating a request, without starting
// a real server on a real port.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := StatusResponse{
		Status:  "ok",
		Message: "service is running",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/status", statusHandler)
	fmt.Println("server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
