package main

import "fmt"

/*

Why use channels instead of a mutex?
Go's philosophy (there's a well-known saying):
"Don't communicate by sharing memory;
share memory by communicating."
A mutex protects shared data that multiple goroutines directly touch.
A channel instead has goroutines pass data to each other through
the pipe — no goroutine ever directly touches another's variables,
sidestepping race conditions in a different way entirely, often considered
cleaner/more idiomatic in Go for many use cases
(though mutexes are still absolutely the right tool sometimes, like your counter example).


by default, channels are unbuffered, meaning a send
blocks until someone's ready to receive, and a receive
blocks until someone sends.

*/

func main() {
	increments := make(chan int)
	done := make(chan bool)

	counter := 0

	// ONE goroutine owns `counter` completely - nobody else ever touches it directly
	go func() {
		for i := range increments {
			counter += i // only this goroutine ever modifies counter - no race possible
		}
		done <- true // signal "I've processed everything, finished"
	}()

	// 1000 goroutines send "please add 1" through the channel - they never touch counter itself
	for i := 0; i < 1000; i++ {
		increments <- 1
	}
	close(increments) // no more increments coming

	<-done // wait until the counting goroutine has processed everything

	fmt.Println("final counter:", counter)
}
