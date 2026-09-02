package main

import "fmt"

// an interface - defines a REQUIRED METHOD, not any data/fields
type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// Dog now has a Sound() method - this makes it satisfy the Animal interface
func (d Dog) Sound() string {
	return "woof!"
}

// Cat also has a Sound() method - it satisfies Animal too
func (c Cat) Sound() string {
	return "Meow!"
}

// ONE function, works for ANY type that has a Sound() method
func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	d := Dog{Name: "Scooby Doo"}
	c := Cat{Name: "Garfield"}

	makeSound(d)
	makeSound(c)
}
