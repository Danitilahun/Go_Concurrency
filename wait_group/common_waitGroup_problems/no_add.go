package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup // Create a WaitGroup

// 	// Problem Explanation:
// 	// The code below starts a goroutine to simulate a task but doesn't add it to the WaitGroup before waiting.
// 	// As a result, wg.Wait() doesn't wait for any tasks to be completed since no task was added to the WaitGroup using Add().

// 	// Simulating work: This could be a function or task
// 	go func() {
// 		defer wg.Done() // Indicate task completion when finished
// 		fmt.Println("Task completed")
// 	}()

// 	wg.Wait()                          // Wait for all tasks added to the WaitGroup to complete, but there are no tasks added
// 	fmt.Println("All tasks completed") // This line might execute before the goroutine finishes, leading to unexpected output
// }
