package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		// Imagine multiple goroutines spawned concurrently.
		// Each spawned goroutine simultaneously increments the WaitGroup counter by 3.
		// Then, three inner goroutines are launched within each outer goroutine,
		// each calling wg.Done() to decrement the counter by 1.

		// Issue Explanation:
		// - Multiple goroutines concurrently increment the WaitGroup counter using wg.Add(3).
		// - Inside each goroutine, three inner goroutines are spawned to decrement the counter.
		// - However, immediately after adding 3 to the counter, wg.Wait() is called,
		//   blocking until the counter becomes zero.
		// - Since the inner goroutines decrement the counter by 1 each, the total decrement
		//   per outer goroutine is 3, expecting the WaitGroup to be ready to pass wg.Wait().
		// - However, there's a potential deadlock situation:
		//   - Some inner goroutines might execute their wg.Done() calls quickly,
		//     allowing the outer goroutines to proceed past wg.Wait() faster.
		//   - But, if certain inner goroutines execute their wg.Done() calls more slowly
		//     than the outer goroutines calling wg.Wait(), a deadlock may occur.
		//   - The outer goroutines will be waiting indefinitely for the WaitGroup counter
		//     to reach zero, while some inner goroutines are still in progress,
		//     leading to a deadlock scenario where the program halts.

		go func() {
			wg.Add(3) // Concurrently increment the WaitGroup counter by 3.
			go func() {
				wg.Done() // Decrement the counter by 1.
			}()
			go func() {
				wg.Done() // Decrement the counter by 1.
			}()
			go func() {
				wg.Done() // Decrement the counter by 1.
			}()
			wg.Wait() // Potential deadlock situation due to improper synchronization.
		}()
	}
}
