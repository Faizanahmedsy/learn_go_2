package main

import "fmt"

func main() {
	i := 7
	inc1(i) //i is copied by value so the I inside the inc function will be 8 but because we are not returning it it woudnt affet the I in the main func and the value will be discarded

	inc2(&i) // if we pass a pointer to the i then it can look up the value at that memory refrence and modify the original version
	fmt.Println(&i)
	fmt.Println(i)

}

func inc1(x int) {
	x++
}

func inc2(x *int) { // accepts a pointer
	*x++
}
