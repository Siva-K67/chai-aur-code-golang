package main

import "fmt"

// naming the empty interface - "Any" is just a label for interface{}
// interface{} means "any type is accepted here" - no restrictions at all
type Any interface {
}

func printAnything(val Any) {
	fmt.Println(val)
}

func main() {
	printAnything(777)
	printAnything("WWII")
	printAnything(true)
}
