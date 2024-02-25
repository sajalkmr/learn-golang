package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
}

func main() {
	// Measure time taken without goroutine
	start := time.Now()

	for i := 0; i < 5; i++ {
		fmt.Println("Main Routine:", i)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}

	elapsedWithoutGoroutine := time.Since(start)

	// Measure time taken with goroutine
	start = time.Now()

	go printNumbers()

	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine:", i)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}

	elapsedWithGoroutine := time.Since(start)

	// Output the results
	fmt.Println("Time taken without goroutine:", elapsedWithoutGoroutine)
	fmt.Println("Time taken with goroutine:", elapsedWithGoroutine)
}
