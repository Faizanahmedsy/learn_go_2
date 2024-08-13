package main

import (
	"fmt"
	"time"
)

func main() {
	// Record the start time
	start := time.Now()

	// Loop that prints numbers from 0 to 999
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}

	// Record the end time
	elapsed := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %s\n", elapsed)
}
