package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go routine")

	go sayHello() //this func will run seprately and the main func does not care weather it finishes exicution or not, main fun will finish without waiting for sayHello

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
