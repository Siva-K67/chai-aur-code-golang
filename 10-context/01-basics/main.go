/*
Let's start with the problem it solves. Imagine a handler that
does something slow — like a slow database query. If the client
gives up and closes the connection, your server should ideally
stop working on it too, rather than wasting time on a request
nobody's waiting for anymore.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

// a function that does slow work, checking ctx periodically to see if it should stop
func doWork(ctx context.Context) {
	for i := 1; i <= 5; i++ {
		// select case here waits on multipe channels at once and runs whichever case becomes ready first
		select {
		case <-ctx.Done():
			// ctx.Done() is a channel that gets closed when the context is cancelled
			fmt.Println("work stopped early: ", ctx.Err())
			return
		case <-time.After(1 * time.Second):
			fmt.Println("working...setp number: ", i)
		}
	}

	fmt.Println("work finished normally")
}

func main() {

	// WithTimeout creates a context that auto-cancels after the given duration
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // always call cancel, even if the timeout already fired - releases resources

	doWork(ctx)
}
