package main

import "fmt"

func main() {

	marks := 75

	//basic if-else
	//also note that no parenthesis around the condition
	//but curly braces mandatory for even single line body
	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	//if with a short statement before the condition
	// score only exists inside this if else block
	// if we try ro access it from outside, Golang gives an error
	if score := 87; score > 80 {
		fmt.Println("Good Score:", score)
	}

	//combining conditions
	x, y := 5, 10
	if x > 0 && y > 0 {
		fmt.Println("Both are positive")
	}
}
