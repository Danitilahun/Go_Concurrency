package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // Create a boolean channel 'done' for synchronization
	now := time.Now()       // Record the current time

	// Launch goroutines for tasks
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	<-done // Wait to receive a signal from the 'done' channel (task1 completed)
	<-done // Wait to receive a signal from the 'done' channel (task2 completed)
	<-done // Wait to receive a signal from the 'done' channel (task3 completed)
	<-done // Wait to receive a signal from the 'done' channel (task4 completed)

	// Print the elapsed time since the start
	fmt.Println("elapsed:", time.Since(now))
}

// Function for task1
func task1(done chan bool) {
	time.Sleep(100 * time.Millisecond) // Simulate work with a sleep of 100 milliseconds
	fmt.Println("task1")               // Print task1 completion message
	done <- true                       // Send a signal on the 'done' channel to indicate completion
}

// Function for task2
func task2(done chan bool) {
	time.Sleep(200 * time.Millisecond) // Simulate work with a sleep of 200 milliseconds
	fmt.Println("task2")               // Print task2 completion message
	done <- true                       // Send a signal on the 'done' channel to indicate completion
}

// Function for task3
func task3(done chan bool) {
	fmt.Println("task3") // Print task3 completion message
	done <- true         // Send a signal on the 'done' channel to indicate completion
}

// Function for task4
func task4(done chan bool) {
	time.Sleep(100 * time.Millisecond) // Simulate work with a sleep of 100 milliseconds
	fmt.Println("task4")               // Print task4 completion message
	done <- true                       // Send a signal on the 'done' channel to indicate completion
}
