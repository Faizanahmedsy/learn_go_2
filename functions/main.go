package main

import (
	"errors"
	"fmt"
)

func main() {

	resp, err := sum(-9, 4)

	// fmt.Println("resp", resp, "err", err)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("resp", resp)
	}
}

func sum(a int, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("Negative number not allowed")
	}

	return a + b, nil

}
