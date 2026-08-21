package main

import (
	"errors"
	"fmt"
)

// custom error type - a struct that carries extra data, not just a message
type ValidationError struct {
	Field string
	Msg   string
}

// this method is what makes ValidationError satisfy Go's built-in error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

func validateAge(age float64) error {
	if age <= 0 {
		return &ValidationError{Field: "age", Msg: "cannot be 0 or negative"}
		// returning a pointer to our custom struct
	}

	return nil
}

func createUser(age float64) error {
	err := validateAge(age)
	if err != nil {
		// wrapping, same as before
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}

func main() {
	err := createUser(-5)

	// errors.As checks: "is there a *ValidationError somewhere in this error chain?"
	// if yes, it copies it into target, so you can access its fields
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("field that failed: ", validationErr.Field)
		fmt.Println("reason: ", validationErr.Msg)
	} else {
		fmt.Println("some other error")
	}

	fmt.Println("Full error message: ", err)
}
