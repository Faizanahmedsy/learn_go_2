package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := sum(22, 55)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("resp", result)
	}
}

func sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Negative number not allowed")
	}
	return a + b, nil
}
