package main

import (
	"sync"
)

// Calling Done more times than Add panics
// Calling Done less times than Add results in deadlock
func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Adds one task to the WaitGroup

	// No call to wg.Done() to signal task completion
	// This will lead to both a panic and a deadlock
	wg.Wait() // Wait indefinitely for the task to complete
}
