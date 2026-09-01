package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func incrementCounter() {
	defer wg.Done()
	counter++
	//counter++ looks like a single atomic step, but under the hood it's actually three separate operations:
	// (1) read the current value of counter,
	// (2) add 1 to it, (3) write the new value back.
}

func main() {

	// launch 1000 goroutines, each incrementing counter once
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait()
	fmt.Println("Finalcounter value ", counter)
}

/*

Output
Finalcounter value  1000
Finalcounter value  1000
Finalcounter value  999
Finalcounter value  974
Finalcounter value  993
Finalcounter value  1000
Finalcounter value  973
Finalcounter value  987

*/
