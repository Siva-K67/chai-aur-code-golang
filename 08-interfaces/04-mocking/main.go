package main

import "fmt"

// Greeter is an interface - anything with a Greet() method satisfies it
type Greeter interface {
	Greet() string
}

// RealGreeter - the "production" version
type RealGreeter struct{}

func (r RealGreeter) Greet() string {
	return "Hello from the real greeter!"
}

// FakeGreeter - the "test" version - returns something fixed, predictable
type FakeGreeter struct{}

func (f FakeGreeter) Greet() string {
	return "fake greeting"
}

// sayHello doesn't care WHICH Greeter it gets - real or fake
func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	real := RealGreeter{}
	fake := FakeGreeter{}

	sayHello(real)
	sayHello(fake)
}
