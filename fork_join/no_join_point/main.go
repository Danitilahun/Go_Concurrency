package main

import (
	"fmt"
	"time"
)

func concurrentFunction() {
	time.Sleep(500 * time.Millisecond) // Simulating some work with a sleep of 100 milliseconds
	fmt.Println("concurrentFunction finished.")
}

func main() {
	go concurrentFunction() // Run concurrentFunction in a goroutine
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Main function finished waiting and Exit.")
}
