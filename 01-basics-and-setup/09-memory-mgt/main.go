package main

import "fmt"

func createNumber() *int {
	x := 45
	return &x
}

func main() {
	//normal variables stored to stack memory
	a := 33
	fmt.Println(a)

	//pointer-holds a memory address
	b := &a
	fmt.Println(b)  // prints address
	fmt.Println(*b) // print the value held hostage inside that address

	/*The createNumber() function is the interesting part: normally, x := 42 would live on the stack
	and disappear once the function returns. But since we return &x (its address), Go's compiler
	detects that x needs to survive past the function call — this is called escape analysis.
	It automatically moves x to the heap instead of the stack, so the pointer stays valid
	after createNumber() returns.
	*/
	ptr := createNumber()
	fmt.Println(*ptr)

	// garbage collection - Go automatically frees memory nobody references anymore
	ptr = nil
	fmt.Println(ptr)

	/*
		33
		0x32a9e3e800b0
		33
		45
		<nil>
	*/

}
