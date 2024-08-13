package main

import "fmt"

func main() {

	// arrays hold multiple var of the same type, and has a fixed length

	var arra [5]int
	arra[2] = 7
	fmt.Println(arra) // [0 0 7 0 0]

	arrb := [3]string{"faizan", "ahmed", "saiyed"}

	slice1 := []string{"faizan", "ahmed", "saiyed"}
	//we can push new elements in slice unlink arrays

	//interanlly go will create a new array with the apppended value
	slice1 = append(slice1, "is great")

	fmt.Println((arrb))

	fmt.Println(slice1)

	fmt.Println("Hello world")
}
