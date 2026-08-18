package main

import "fmt"

func addOne(n int) {
	n = n + 1
}

func addOnePointer(n *int) {
	*n = *n + 1
}

func main() {

	//pass by val- original doesm not change
	// why ? coz go is pass by value by default
	// fun gets a copy of x. any changes to this
	// copy does not change the original
	x := 5
	addOne(x)
	fmt.Println(x)

	//pass by pointer- original changes here
	// why ? address of x is passed.
	// Now n *int is a pointer, and *n = *n + 1
	// says "go to that address and modify the
	// value stored there directly" — so the original
	// x actually changes this time.
	addOnePointer(&x)
	fmt.Println(x)

	/*
		o/p
		5
		6
	`	*/

}
