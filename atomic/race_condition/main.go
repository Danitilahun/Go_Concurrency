package main

import (
	"fmt"
	"sync"
)

// To test for race condition, run this example with the -race flag
// Command: go run -race main.go
func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		count += 10 // Potential race condition: Concurrently updating 'count' without synchronization.
	}()

	go func() {
		defer wg.Done()
		count -= 15 // Potential race condition: Concurrently updating 'count' without synchronization.
	}()

	go func() {
		defer wg.Done()
		count++ // Potential race condition: Concurrently updating 'count' without synchronization.
	}()

	go func() {
		defer wg.Done()
		count = 0 // Potential race condition: Concurrently updating 'count' without synchronization.
	}()

	go func() {
		defer wg.Done()
		count = 100 // Potential race condition: Concurrently updating 'count' without synchronization.
	}()

	wg.Wait()

	fmt.Println("count", count)
}
