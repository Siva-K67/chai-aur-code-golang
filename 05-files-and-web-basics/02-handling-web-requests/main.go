package main

import (
	"fmt"
	"net/http"
)

// this exact function signature is called a handler. Every function that
// responds to a web request in Go has this shape: a ResponseWriter
// (to send data back) and a *Request (info about what came in).
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the home page of Ippo Makunochi!")
	// fmt.Fprintln(w, ...) instead of printing to your terminal, it writes to w —
	//  which sends it back to the browser as the response.
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page of Hajime no ippo!")
}

func main() {

	//this is routing: "when someone visits /, run homeHandler."
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("server starting on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
