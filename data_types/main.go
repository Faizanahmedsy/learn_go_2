package main

import (
	"fmt"
)

func main() {
	// String
	var name string = "Faizan"
	fmt.Println("Name:", name)

	// Integer
	var roll_no int = 14
	fmt.Println("Roll No:", roll_no)

	// Float
	var percentage float64 = 92.5
	fmt.Println("Percentage:", percentage)

	// Boolean
	var isGraduated bool = true
	fmt.Println("Is Graduated:", isGraduated)

	// Complex Number
	var complexNum complex128 = complex(5, 7)
	fmt.Println("Complex Number:", complexNum)

	// Byte (alias for uint8)
	var initial byte = 'F'
	fmt.Println("Initial:", initial)

	// Rune (alias for int32, used for representing Unicode code points)
	var unicodeChar rune = '界'
	fmt.Println("Unicode Character:", unicodeChar)

	// Array
	var marks [5]int = [5]int{85, 90, 78, 92, 88}
	fmt.Println("Marks:", marks)

	// Slice
	var subjects []string = []string{"Math", "Science", "English"}
	fmt.Println("Subjects:", subjects)

	// Map
	student := make(map[string]interface{})
	student["name"] = "Faizan"
	student["roll_no"] = 14
	student["percentage"] = 92.5
	student["isGraduated"] = true
	fmt.Println("Student Map:", student)

	// Struct
	type Student struct {
		Name        string
		RollNo      int
		Percentage  float64
		IsGraduated bool
	}
	studentStruct := Student{Name: "Faizan", RollNo: 14, Percentage: 92.5, IsGraduated: true}
	fmt.Println("Student Struct:", studentStruct)

	// Pointer
	var namePtr *string = &name
	fmt.Println("Pointer to Name:", namePtr, "Value:", *namePtr)

	// Function Type
	printFunc := fmt.Println
	printFunc("This is a function type")

	// Channel
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	fmt.Println("Channel Value:", <-ch)

	// Interface (empty interface can hold values of any type)
	var anyType interface{}
	anyType = "Faizan"
	fmt.Println("Any Type:", anyType)
	anyType = 92.5
	fmt.Println("Any Type:", anyType)
}
