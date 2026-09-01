package main

import "fmt"

func main() {
	// unbuffered channel - a send blocks until someone receives, and vice versa
	messages := make(chan string)

	// sending directly in main(), with NO goroutine to receive it
	messages <- "hello"

	fmt.Println("this line never runs")
}

/*
Recall from stage 5: an unbuffered channel send blocks until
something is ready to receive it.
Here, messages <- "hello" blocks... and waits... and waits — but there's no
other goroutine anywhere trying to receive from messages. Nobody's ever going to show up.
*/
