package main

import "fmt"

// go has no inbuild remove function for slices
// so we manually do it

func main() {

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Orignal array ", numbers)

	indexToRemove := 2

	//take everything before it, plus everything after it
	numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	fmt.Println("after removing index 2 ", numbers)

}
