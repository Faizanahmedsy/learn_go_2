package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go routine")

	sayHello() //main func will wait for this program to finish and then the sayName will be executed

	sayName()
}

func sayHello() {
	fmt.Println("Hello program start")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("program end")
}

func sayName() {
	fmt.Println("Faizan")
}
