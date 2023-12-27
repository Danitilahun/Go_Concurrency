package main

import "sync/atomic"

type waitGroup struct {
	counter int64
}

// atomic.AddInt64 is a function provided by Go's sync/atomic package that performs atomic addition on a 64-bit integer variable.
// This function is used to safely update an int64 variable in a concurrent environment without causing data races or conflicts
// between multiple goroutines accessing the same variable simultaneously.In a concurrent program where multiple goroutines are
// manipulating the same variable concurrently, traditional operations like x = x + 1 might not be safe due to potential race conditions.
// A race condition occurs when multiple operations on shared data interfere with each other, leading to unpredictable or erroneous behavior.
// The atomic.AddInt64 function ensures that the addition operation on the int64 variable is performed atomically without the possibility
// of interference from other concurrent operations. It takes care of synchronization and ensures that the addition operation is completed
//  without interruption.
func (wg *waitGroup) Add(n int64) {
	atomic.AddInt64(&wg.counter, n)
}
