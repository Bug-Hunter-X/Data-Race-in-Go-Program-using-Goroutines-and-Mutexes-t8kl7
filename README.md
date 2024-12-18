# Data Race in Go Program using Goroutines and Mutexes

This repository contains a Go program that demonstrates a data race when using goroutines and mutexes. The program attempts to increment a counter using multiple goroutines, but due to improper mutex usage, the final counter value may be incorrect.

## Bug Description

The program uses a `sync.Mutex` to protect the `counter` variable. However, the way the mutex is used can lead to a data race. The solution provides a way to fix this using an atomic counter which eliminates the need for a mutex altogether. 

## Bug Solution

The solution fixes the data race by using the `sync/atomic` package which provides atomic operations that are safe for concurrent access. This approach makes the code simpler and avoids the potential for deadlocks or race conditions. 

## How to run the code

1. Clone the repository: `git clone <repository_url>`
2. Navigate to the repository directory: `cd <repository_directory>`
3. Run the buggy program: `go run bug.go`
4. Run the fixed program: `go run bugSolution.go`

Observe the difference in the output and understand the effect of the race condition on the result.