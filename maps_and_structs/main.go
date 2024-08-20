package main

import "fmt"

type Stream struct {
	Program string
	Branch  string
}

type Student struct {
	Name      string
	RollNo    string
	Grade     string
	CGPA      float64
	Graduated bool
	Stream    Stream
}

func main() {

	Student := Student{
		Name:      "faizan",
		RollNo:    "14",
		Grade:     "A",
		CGPA:      9.7,
		Graduated: false,
		Stream: Stream{
			Program: "Btech",
			Branch:  "IT",
		},
	}

	faculty := []string{"faizan", "aaftab", "saad", "aaqib"}

	ranker := Student.CGPA

	fmt.Println("Student:", Student)
	fmt.Println("Faculty:", faculty)
	fmt.Println("Ranker CGPA:", ranker)

}
