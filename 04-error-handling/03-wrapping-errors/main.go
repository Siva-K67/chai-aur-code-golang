package main

import (
	"errors"
	"fmt"
)

// a 'lower' level function that can fail
func findUser(id int) error {
	return errors.New("User not found")
}

// a "higher level" func that calls another func findUser, and adds context if it fails
func loadProfile(id int) error {
	err := findUser(id)
	if err != nil {
		// %w wraps the original error, adding a message on top
		return fmt.Errorf("Loading profile...%w", err)
	}

	return nil
}

func main() {
	err := loadProfile(5)
	fmt.Println(err)
}
