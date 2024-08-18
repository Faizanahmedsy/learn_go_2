package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Err in getting resp", resp.Status)

		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Err", err)
		return
	}

	fmt.Println("Data", string(data))

}
