package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Println("number:", i)
		time.Sleep(500 * time.Millisecond) // pause half a second, to slow things down so we can observe order
	}
}

func printLetters() {
	for _, l := range []string{"a", "b", "c"} {
		fmt.Println("letter:", l)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	fmt.Println("------normal seq calls------")
	// WITHOUT concurrency - these run one after another, in order
	printNumbers()
	printLetters()

	fmt.Println("------now with goroutines------")
	// go keyword - runs this function CONCURRENTLY, doesn't wait for it to finish
	go printNumbers()
	go printLetters()

	// without this, main() would exit immediately, killing the goroutines before they even run
	time.Sleep(2 * time.Second)
	fmt.Println("main over")

	/*
		Adding go before a function call tells Go: "start running this concurrently, don't wait for it to finish
		before moving to the next line." This concurrently-running unit is called a goroutine — a lightweight
		thread managed by Go's runtime (much cheaper than an OS thread, you can have thousands running at once).

	*/

	/*
			------normal seq calls------
		number: 1
		number: 2
		number: 3
		letter: a
		letter: b
		letter: c
		------now with goroutines------
		letter: a
		number: 1
		number: 2
		letter: b
		letter: c
		number: 3
		main over
	*/
}
