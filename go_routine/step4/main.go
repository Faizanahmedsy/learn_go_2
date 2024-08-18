package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Party started")

	go track_faizan()
	go track_aaftab()
	go tracK_aaqib()
	go track_saad()
}

func track_faizan() {
	fmt.Println("Faizan is coming")
	time.Sleep(1 * time.Second)
	fmt.Println("Faizan came after 1 second")
}

func track_aaftab() {
	fmt.Println("Aaftab is coming")
	time.Sleep(2 * time.Second)
	fmt.Println("Aaftab came after 2 seconds")
}

func tracK_aaqib() {
	fmt.Println("Aaqib is coming")
	time.Sleep(3 * time.Second)
	fmt.Println("Aaqib came after 3 seconds")
}

func track_saad() {
	fmt.Println("Saad is coming")
	time.Sleep(4 * time.Second)
	fmt.Println("Saad came after 4 seconds")
}
