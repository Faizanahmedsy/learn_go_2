package main

import (
	"fmt"
	"time"
)

func main() {
	// Record the start time
	start := time.Now()

	// for Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while loop

	j := 0

	for j < 10 {
		fmt.Println("hey", j)
		j++
	}

	// Record the end time
	elapsed := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %s\n", elapsed)
}
