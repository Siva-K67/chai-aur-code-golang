package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Mouse struct {
	Name string
}

func (d Dog) Sound() string {
	return "Scooby Dooby Doo!"
}

func (c Cat) Sound() string {
	return "I hate Mondays. Meow."
}

func (m Mouse) Sound() string {
	return "Squeak! (Jerry says hi)"
}

func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	d := Dog{Name: "Scooby-Doo"}
	c := Cat{Name: "Garfield"}
	m := Mouse{Name: "Jerry"}

	// a variable of type Animal (the interface) can hold ANY type that satisfies it
	var a Animal

	//So the short answer: a = d works because Go checked Dog's method set against
	// what Animal requires, found a match, and allowed it — not because Dog and
	// Animal are secretly the same type.
	a = d
	makeSound(a)

	a = c
	makeSound(a)

	a = m
	makeSound(a)

	// a SLICE of the interface type - holds a mix of different concrete types together
	animals := []Animal{d, c, m}

	fmt.Println("-- looping through mixed types --")
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}
