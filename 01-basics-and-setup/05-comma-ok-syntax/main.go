package main

import "fmt"

func main() {

	//map for comma ok syntax
	scores := map[string]int{
		"go":     90,
		"python": 85,
		"c":      99,
	}

	//ok becomes true since go exists in map
	value, ok := scores["go"]
	fmt.Println(value, ok)

	//ok2 is false because c++ isnt in the map
	//value2 defaults to 0
	value2, ok2 := scores["C++"]
	fmt.Println(value2, ok2)

	//now type assertion
	// is this value the type i think it is ?

	//i is an interface, it can hold any type.
	//i.(string) means is i a string ?
	var i interface{} = "hello"
	s, ok3 := i.(string)
	fmt.Println(s, ok3)

	// is i an int ??
	// i isnt an int so its false and n defaults to 0
	n, ok4 := i.(int)
	fmt.Println(n, ok4)

}
