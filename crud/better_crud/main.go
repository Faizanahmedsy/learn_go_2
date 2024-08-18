package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// When you write json:"userId", it tells the Go JSON encoder/decoder that when converting this struct to or from JSON, the corresponding JSON key for the UserID field should be "userId"

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:id`
	Title     string `json"title`
	Completed bool   `json:completed`
}

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

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)

	// You should pass &todo instead of todo when decoding the JSON response because json.NewDecoder(resp.Body).Decode() expects a pointer to a struct. This allows the function to directly modify the struct's fields.

	if err != nil {
		fmt.Println("Err decoding", err)
	}

	fmt.Println("data", todo)

}
