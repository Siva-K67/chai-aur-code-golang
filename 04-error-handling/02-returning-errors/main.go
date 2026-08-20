package main

import (
	"errors"
	"fmt"
)

// this function can fail, so its return type includes the error
// convention: return the real result AND an error, error goes LAST
func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cant div by 0")
	}
	return a / b, nil
}

func main() {

	//case 1
	result, err := divide(50, 5)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("Answer: ", result)
	}

	//case 2
	result2, err2 := divide(50, 0)
	if err2 != nil {
		fmt.Println("error: ", err2)
	} else {
		fmt.Println("Answer: ", result2)
	}

}
