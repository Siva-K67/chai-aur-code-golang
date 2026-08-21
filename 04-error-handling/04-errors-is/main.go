package main

import (
	"errors"
	"fmt"
)

// a known, reusable error - a "sentinel error"
// declared once at package level, so it can be compared against later
var ErrNotFound = errors.New("record not found")

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
		// returning our known sentinel error directly
	}

	return nil
}

func loadProfile(id int) error {
	err := findUser(id)
	if err != nil {
		return fmt.Errorf("loading profile...%w", err)
		// wrapping it, same as before - adds context but keeps the original error inside
	}

	return nil
}

func main() {

	err := loadProfile(0)

	//even though err was wrapped, error.Is can still detect ErrorNotFound indide it
	if errors.Is(err, ErrNotFound) {
		fmt.Println("This is specifically a not found error")
	} else {
		fmt.Println("Some other error occured")
	}

	fmt.Println("HEre is the full error message: ", err)

}

/*

output

This is specifically a not found error
HEre is the full error message:  loading profile...record not found

*/
