package main

import (
	"fmt"
	"net/http"
)

func main() {
	connectDB()
	defer db.Close()

	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/courses/", courseByIDHandler)

	fmt.Println("server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}
