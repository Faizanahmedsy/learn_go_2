package main

import "fmt"

type user struct {
	name string
	age  int
	sex  string
}

type car struct {
	make   string
	engine string
	price  int
}

func main() {

	user_1 := user{
		name: "faizan",
		age:  12,
		sex:  "male",
	}

	user_1_car := car{
		make:   "bugati",
		engine: "v8",
		price:  500000,
	}

	fmt.Println("user", user_1.name, "has a car", user_1_car.make, "of price", user_1_car.price)
}
