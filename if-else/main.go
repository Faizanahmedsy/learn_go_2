package main

import "fmt"

func main() {

	var x int

	fmt.Println("Please enter a number")
	fmt.Scan((&x))

	if x%2 == 0 {
		fmt.Println("X is a primary number")
	} else {
		fmt.Println("X is not a primary number")
	}
}
