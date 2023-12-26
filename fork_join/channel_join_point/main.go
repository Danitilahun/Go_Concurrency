package main

import (
	"fmt"
	"time"
)

func concurrentFunction(done chan bool) {
	time.Sleep(500 * time.Millisecond) // Simulating some work with a sleep of 500 milliseconds
	fmt.Println("concurrentFunction finished.")
	done <- true // Send a signal on the 'done' channel to indicate completion
}

func main() {
	now := time.Now()
	done := make(chan bool) // Create a channel to signal completion

	go concurrentFunction(done) // Run concurrentFunction in a goroutine

	<-done // Wait to receive a signal from the 'done' channel
	fmt.Println("Main function finished waiting and Exit.")
	fmt.Println("Time elapsed : ", time.Since(now))
}
