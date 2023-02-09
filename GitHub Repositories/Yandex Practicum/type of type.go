package main

import "fmt"

type Name string
type Fruit int

var fruit Fruit
var name Name

func main() {

	fruit = 10
	name = Name(fruit) // так, после приведения типов, работает

	fmt.Println("123")

	fmt.Println(fruit)

	for _, data := range name {
		fmt.Println(data)

	}

}
