package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the home page of Ippo Makunochi!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page of Takamura.")
}

// reads a query parameter from the URL, e.g. /greet?name=Siva
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// r.URL holds the full URL
	//.Query() parses the ?
	// .Get("name") pulls out the value for "name" - empty string if not present
	if name == "" {
		name = "stranger"
	}

	fmt.Fprintln(w, "HEllo ", name)
}

// a struct to share our JSON response
//back ticks are aka struct tags
// renames this to "status" in the JSON output (Go name "Status" stays capitalized, JSON key doesn't)

type StatusResponse struct {
	Status  string ` json:"status" `
	Message string ` json:"message" `
}

// responds with actual JSON, not plain text - this is what real APIs do
func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := StatusResponse{
		Status:  "ok",
		Message: "service is running",
	}

	// tells the browser/client "this response is JSON"
	w.Header().Set("Content Type", "application/json")

	// converts the struct into JSON and writes it to w
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/status", statusHandler)

	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
