package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// simulates a slow database call that respects context cancellation
func slowDBCall(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		// pretend this is a real DB query that took 3 seconds
		return "course details", nil
	case <-ctx.Done():
		// the request was cancelled/timed out before the "query" finished
		return "", ctx.Err()
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	// r.Context()
	// every single *http.Request automatically carries its own
	// context.Context, created internally by Go's net/http package.
	// You don't create it yourself with WithTimeout this time — Go
	// already did that for you, tied to the actual lifetime of the HTTP connection.
	ctx := r.Context()

	result, err := slowDBCall(ctx)
	if err != nil {
		fmt.Println("request cancelled or timeout: ", err)
		http.Error(w, "request timed out", http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, "got results: ", result)
}

func main() {
	http.HandleFunc("/slow", slowHandler)
	fmt.Println("server starting at pot 8080...")
	http.ListenAndServe(":8080", nil)
}
