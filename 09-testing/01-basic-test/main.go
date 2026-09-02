package main

import "fmt"

// a function we want to test
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
}
