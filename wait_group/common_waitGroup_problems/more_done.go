package main

// import "sync"

// func main() {
// 	var wg sync.WaitGroup
// 	// Problem Description:
// 	// Calling Done() without a prior call to Add() causes a panic in sync.WaitGroup.

// 	wg.Done() // Calls Done() without adding any tasks using Add()

// 	// The program panics with a message indicating "sync: negative WaitGroup counter"
// 	// The panic occurs because Done() was called without adding any tasks using Add(),
// 	// leading to an incorrect decrement of the WaitGroup counter.
// }
