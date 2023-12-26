# Go_Concurrency

wmic cpu get NumberOfLogicalProcessors : This command will provide you with the number of logical processors (cores) available on your Windows system.
wmic cpu get NumberOfCores : The number of physical cores (NumberOfCores) should give you the count of physical cores on your processor.

Concurrency is the ability of an application to handle multiple tasks or processes at the same time, seemingly simultaneously, by interleaving their execution.

Parallelism involves the simultaneous execution of multiple tasks or processes, where they genuinely run at the same time, utilizing multiple processors or CPU cores.

In Go (Golang), "fork-join" refers to a programming pattern where a task is divided into smaller sub-tasks (forking), which are executed independently, and then the results are combined (joined) to produce the final result.

Fork: Involves breaking down a task into smaller units of work that can be executed concurrently or in parallel.
Join: The process of merging or aggregating the results obtained from the individual tasks back together to form the final result.
