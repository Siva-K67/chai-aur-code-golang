package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

// Mutex = "mutual exclusion" - ensures only ONE goroutine can access
// the protected section at a time
var mu sync.Mutex

func incrementCounter() {
	defer wg.Done()

	mu.Lock() // Lock() - "I'm about to touch shared data, block anyone else trying to do the same"
	counter++
	mu.Unlock() // Unlock() - "I'm done, someone else can go now"

}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait()

	// this will now CONSISTENTLY print 1000, every single run
	fmt.Println("final counter:", counter)
}
