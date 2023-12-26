# Go_Concurrency

wmic cpu get NumberOfLogicalProcessors : This command will provide you with the number of logical processors (cores) available on your Windows system.
wmic cpu get NumberOfCores : The number of physical cores (NumberOfCores) should give you the count of physical cores on your processor.

Concurrency is the ability of an application to handle multiple tasks or processes at the same time, seemingly simultaneously, by interleaving their execution.

Parallelism involves the simultaneous execution of multiple tasks or processes, where they genuinely run at the same time, utilizing multiple processors or CPU cores.

In Go (Golang), "fork-join" refers to a programming pattern where a task is divided into smaller sub-tasks (forking), which are executed independently, and then the results are combined (joined) to produce the final result.

Fork: Involves breaking down a task into smaller units of work that can be executed concurrently or in parallel.
Join: The process of merging or aggregating the results obtained from the individual tasks back together to form the final result.

Concurrent programs face several common challenges and problems due to the nature of concurrency and the interactions between concurrently executing tasks. Some of the common issues include:

1. Race Conditions: When multiple threads or goroutines access shared resources (like variables or data structures) concurrently without proper synchronization, race conditions can occur. Race conditions lead to unpredictable behavior and incorrect results due to the unexpected interleaving of operations.

2. Deadlocks: Deadlocks occur when two or more tasks are waiting indefinitely for each other to release resources that they hold. This situation halts the progress of all involved tasks, leading to a complete standstill.

3. Livelocks: Similar to deadlocks, livelocks involve tasks being unable to progress due to continuously changing states without reaching a resolution. Tasks keep responding to each other's actions, but no progress is made.

4. Resource Starvation: This occurs when a task or thread is unable to access necessary resources (like CPU time, memory, or I/O) because other tasks monopolize or block those resources.

5. Synchronization Overhead: Introducing synchronization mechanisms (like locks, mutexes, or channels) to manage concurrent access to shared resources can lead to overhead. Excessive use of synchronization can impact performance, scalability, and increase complexity.
