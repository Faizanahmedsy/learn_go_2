package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "Faizan"
	m["surname"] = "Saiyed"

	for key, value := range m {
		fmt.Println("key", key, "value", value)

	}
}
