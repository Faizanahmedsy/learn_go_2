package main

import (
	"fmt"
	"time"
)

func main() {

	current_time := time.Now()
	fmt.Println("Current Time:", current_time)
	fmt.Printf("Type of current_time %T\n", current_time) // Time is of datatype time.Time

	// in go - yyyy/mm/dd  day is represented as 20006/01/02 Monday because this is when go lang was announced

	formatted_time := current_time.Format("2006/01/02, Monday")

	fmt.Println("Formatted time:", formatted_time)

	date_str := "2023-11-25"

	//The below function converts a string into a time.Time data type
	formatted_date_str, _ := time.Parse("2006-01-02", date_str)
	fmt.Printf("type of formatted_date_str %T\n", formatted_date_str)
	fmt.Println(formatted_date_str)

}
