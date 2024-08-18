package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Party Time")
	go track_faizan()
	go track_aaftab()

	//  defer keyword is used to schedule a function call to be executed just before the surrounding function returns, regardless of where the return happens. This is particularly useful for tasks like closing files, releasing resources, or printing final messages, ensuring that these actions happen even if the function exits early due to an error or a return statement.
	defer fmt.Println("party over")
	time.Sleep(6000 * time.Millisecond)
}

func track_faizan() {
	fmt.Println("Aaftab coming")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Aaftab entered after 5 second")
}

func track_aaftab() {
	fmt.Println("Faizan coming")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Faizan entered after 3 second")

}
