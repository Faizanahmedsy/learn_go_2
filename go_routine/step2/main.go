package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go routine")

	go sayHello()

	sayName()

	// manually added a delay so that we can see the sayHello execution
	time.Sleep(6000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello program start")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("program end")
}

func sayName() {
	fmt.Println("Faizan")
}
