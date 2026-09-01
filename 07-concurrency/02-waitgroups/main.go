package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup — think of it as a countdown counter specifically for tracking running goroutines.
func printNumbers(wg *sync.WaitGroup) {
	// Done() tells the WaitGroup "this goroutine has finished" - runs when the function returns
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("number:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	// defer wg.Done() decrements the counter by 1, called when the goroutine's function returns.
	defer wg.Done()

	for _, l := range []string{"a", "b", "c"} {
		fmt.Println("letter:", l)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// WaitGroup - a counter that tracks how many goroutines are still running
	var wg sync.WaitGroup

	// Add(1) - "I'm about to start one more goroutine, track it"
	wg.Add(1)
	go printNumbers(&wg)

	wg.Add(1)
	go printLetters(&wg)

	// Wait() blocks here until the counter drops back to 0
	// (i.e. until BOTH goroutines have called Done())
	// main() pauses right here until the counter returns to 0
	wg.Wait()

	fmt.Println("main done")
}
