package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// shape of json we expect the client sends to us the server
type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// shape of what we (the server) send back to the client
type CreateUserResponse struct {
	Message string `json:"message"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	//allow only POST in this endpoint
	//this checks which http method was used
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	//creating an obj equivalent in Golang
	var req CreateUserRequest
	//this is the reverse of encode of json. we decode the json into the struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Println("Received: ", req.Name, req.Age)

	response := CreateUserResponse{
		Message: fmt.Sprintf("User %s created", req.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/users", createUserHandler)

	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error: ", err)
	}
}

/*
output

server starting on port 8080...
Received:  Kamado Tanjiro 17

*/
