/*
The problem generics solve:

func printInt(val int) {
	fmt.Println(val)
}

func printString(val string) {
	fmt.Println(val)
}

Same logic, duplicated per type. Before generics, this was a real annoyance in Go — you had a genuine gap between "works for one specific type" and "works for absolutely anything" (interface{}, which loses all type safety).
Generics fill that gap — "works for several specific types, safely"
*/

package main

import "fmt"

// [T any] declares a TYPE PARAMETER - T stands in for whatever type gets used
// "any" means T can be anything (same as interface{}, just newer syntax)
func printValue[T any](val T) {
	fmt.Println(val)
}

// a generic function that returns a value of type T
func first[T any](items []T) T {
	return items[0]
}

func main() {
	printValue(42)             //T becomes an int here
	printValue("Inazuma 11 !") //T becomes a string
	printValue(777.777)        // T becomes a float64

	nums := []int{10, 20, 30}
	fmt.Println(first(nums)) // T = int, returns 10

	names := []string{"Shikamaru", "Temari"}
	fmt.Println(first(names)) // T = string, returns "Siva"

}
