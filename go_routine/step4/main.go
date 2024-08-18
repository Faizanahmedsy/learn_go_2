package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a WaitGroup instance to synchronize the completion of goroutines
	var wg sync.WaitGroup
	fmt.Println("Party started")

	// Increment the WaitGroup counter by 2
	wg.Add(2)

	// Launch the track_faizan goroutine and pass the WaitGroup pointer to it -  means that we are going to wait for faizan to come and we will not finish the party till then
	go track_faizan(&wg)
	go track_aaftab(&wg)

	// Launch the other goroutines without involving them in the WaitGroup
	go tracK_aaqib()
	go track_saad()

	// Wait for all goroutines that are part of the WaitGroup to finish
	wg.Wait()

	// This will print "Party over" after the WaitGroup counter reaches zero (i.e., all tracked goroutines have completed)
	defer fmt.Println("Party over")
}

func track_faizan(wg *sync.WaitGroup) {
	fmt.Println("Faizan	 is coming")
	time.Sleep(5 * time.Second)
	fmt.Println("Faizan came after 5 second")

	// Decrement the WaitGroup counter by 1, indicating that this goroutine is done
	defer wg.Done()
}

func track_aaftab(wg *sync.WaitGroup) {
	fmt.Println("Aaftab is coming")
	time.Sleep(10 * time.Second)
	fmt.Println("Aaftab came after 10 seconds")
	defer wg.Done()

}

func tracK_aaqib() {
	fmt.Println("Aaqib is coming")
	time.Sleep(12 * time.Second)
	fmt.Println("Aaqib came after 12 seconds")
}

func track_saad() {
	fmt.Println("Saad is coming")
	time.Sleep(14 * time.Second)
	fmt.Println("Saad came after 14 seconds")
}
