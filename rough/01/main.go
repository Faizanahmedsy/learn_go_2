package main

import "fmt"

func main() {
	var num int

	// Scans user input for the number
	fmt.Scan(&num)

	// Prints the last digit of the number
	fmt.Printf("answer: %d\n", num%10)
}
