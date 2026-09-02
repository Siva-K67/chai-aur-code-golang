package main

import "testing"

// FakeGreeter - only exists for tests, returns a fixed, predictable value
type FakeGreeter struct{}

func (f FakeGreeter) Greet() string {
	return "fake greeting"
}

func TestSayHello(t *testing.T) {
	fake := FakeGreeter{}

	// sayHello accepts ANY Greeter - here we hand it the fake, not the real one
	result := sayHello(fake)
	expected := "fake greeting"

	if result != expected {
		t.Errorf("sayHello(fake) = %q; expected %q", result, expected)
	}
}
