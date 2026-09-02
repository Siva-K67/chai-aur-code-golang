package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// simple in-memory version, just for this test example - no real DB needed
var courses = []Course{
	{ID: 1, Name: "Go Basics", Price: 499},
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(courses)

	case http.MethodPost:
		var newCourse Course
		if err := json.NewDecoder(r.Body).Decode(&newCourse); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		courses = append(courses, newCourse)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "course created"})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/courses", coursesHandler)
	fmt.Println("server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
