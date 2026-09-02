package main

import "fmt"

// Greeter is an interface - anything with a Greet() method satisfies it
type Greeter interface {
	Greet() string
}

// RealGreeter - used in production
type RealGreeter struct{}

func (r RealGreeter) Greet() string {
	return "Hello from the real greeter!"
}

// sayHello depends only on the interface, not a specific struct
func sayHello(g Greeter) string {
	return g.Greet()
}

func main() {
	real := RealGreeter{}
	fmt.Println(sayHello(real))
}
