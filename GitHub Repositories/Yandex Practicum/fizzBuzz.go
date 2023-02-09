package main

import "fmt"

//Напишите программу, которая выводит на экран числа от 1 до 100.
//При этом вместо чисел, кратных трём, программа должна выводить слово Fizz,
//а вместо кратных пяти — Buzz. Если число кратно и трём, и пяти, программа должна выводить слово FizzBuzz.

func main() {

	for i := 1; i < 100; i++ {

		switch {

		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")

		case i%3 == 0:
			fmt.Println("Fizz")

		case i%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(i)

		}

	}

}
