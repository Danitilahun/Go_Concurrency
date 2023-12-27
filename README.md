# Go Concurrency

## CPU Information Retrieval Commands on Windows

- `wmic cpu get NumberOfLogicalProcessors`: Provides the count of logical processors (cores) available on the Windows system.
- `wmic cpu get NumberOfCores`: Offers the count of physical cores available on the processor.

## Definitions

- **Concurrency**: The capability of an application to handle multiple tasks or processes simultaneously by interleaving their execution.
- **Parallelism**: Simultaneous execution of multiple tasks genuinely running at the same time using multiple CPU cores.

## Fork-Join in Go (Golang)

"Fork-join" in Go refers to dividing a task into smaller sub-tasks (forking), executing them independently, and then merging the results (joining) to produce the final output.

- **Fork**: Breaking a task into smaller units for concurrent or parallel execution.
- **Join**: Merging the results from individual tasks to form the final result.

## Common Challenges in Concurrent Programming

1. **Race Conditions**: Unpredictable behavior due to multiple threads accessing shared resources without synchronization.
2. **Deadlocks**: Tasks waiting indefinitely for each other, halting progress.
3. **Livelocks**: Continuous state changes without resolution, preventing progress.
4. **Resource Starvation**: Inability to access necessary resources due to monopolization by other tasks.
5. **Synchronization Overhead**: Introduction of synchronization mechanisms impacting performance and complexity.

## Example Code

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Wait()
    fmt.Println("Wait() executed immediately")
}
```
