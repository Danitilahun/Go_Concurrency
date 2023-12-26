package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrentFunction() {
	time.Sleep(500 * time.Millisecond) // Simulating some work with a sleep of 500 milliseconds
	fmt.Println("concurrentFunction finished.")
}

func main() {
	now := time.Now()
	var wait_Group sync.WaitGroup // Create a WaitGroup object
	wait_Group.Add(1)             // Add 1 to the WaitGroup counter

	go func() {
		defer wait_Group.Done() // Decrement the WaitGroup counter when the goroutine finishes
		concurrentFunction()    // Run concurrentFunction in a goroutine
	}()

	// to wait for the goroutine to finish
	wait_Group.Wait() // Wait for the WaitGroup counter to reach 0
	fmt.Println("Main function finished waiting and Exit.")
	fmt.Println("Time elapsed : ", time.Since(now))
}
