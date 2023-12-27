package main

import (
	"sync"
	"time"
)

func main() {
	// Initialize a WaitGroup. It will be used to synchronize goroutines.
	var wg sync.WaitGroup

	// Add 1 to the WaitGroup counter.
	wg.Add(1)

	// Launch a goroutine to perform asynchronous tasks.
	go func() {
		// Simulate a short delay in the goroutine.
		time.Sleep(time.Millisecond)

		// Decrement the WaitGroup counter by 1.
		wg.Done()

		// Increment the WaitGroup counter by 1.
		// Note: Modifying the count while it's at zero before Wait can lead to synchronization issues.
		// According to the official documentation, adding a positive delta while the counter is zero
		// should happen before a Wait. This behavior is discouraged as it can lead to race conditions.
		wg.Add(1)
	}()

	// Wait for all goroutines to finish by blocking until the WaitGroup counter becomes zero.
	wg.Wait()
}
