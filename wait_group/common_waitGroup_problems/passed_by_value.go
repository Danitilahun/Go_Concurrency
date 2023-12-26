package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg) // Passing WaitGroup by value (creating a copy)

	wg.Wait() // Waits for the WaitGroup to complete, but it may not synchronize as expected
}

func work(wg sync.WaitGroup) { // Accepting WaitGroup by value
	defer wg.Done() // Problem: Done() is called on a copy of the WaitGroup
	fmt.Println("work is done")
}
