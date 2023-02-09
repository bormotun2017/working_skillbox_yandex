package main

import (
	"fmt"
)

type panica interface {
}

func main() {
	//divideByZero()
	fmt.Println(divide(1, 0))
	fmt.Println("we survived dividing by zero!")
	fmt.Println("мы обошли панику с помощью функции recover и продолжили выполнение программы")

}

func divide(a, b int) int {
	//defer и recover должны находиться в функции в которой вызвана паника, если поместить в main, то не сработает
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	//вызвали панику случайно
	//return a / b

	//вызвали панику принудительно
	var x panica
	x = "123"
	panic(x)

}
