package main

import "fmt"

//Дан массив целых чисел A и целое число k.
//Нужно найти и вывести индексы пары чисел, сумма которых равна k.
//Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке.

func main() {

	arr := make(map[int]int)
	arrindex := make([]int, 0, 0)

	for i := 0; i < 9; i++ {

		arr[i] = i

	}
	fmt.Println(arr)
	k := 15

	for index, data := range arr {

		for i := 0; i < len(arr); i++ {

			if data+arr[i] == k {
				arrindex = append(arrindex, index, i)
			}

		}

	}

	fmt.Println()
	fmt.Println(k, "\n")

	if len(arrindex) == 0 {
		fmt.Println("нет таких чисел")
	}

	fmt.Println(arrindex)
}
