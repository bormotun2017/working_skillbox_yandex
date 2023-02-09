package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Удаление последнего элемента слайса:
	s := []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[:len(s)-1]
	}
	fmt.Println(s) // [1 2]

	//Удаление первого элемента слайса:
	s = []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[1:]
	}
	fmt.Println(s) // [2 3]

	//Удаление элемента слайса с индексом i:
	s = []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s) { // защищаемся от паники
		s = append(s[:i], s[i+1:]...)
	}
	fmt.Println(s) // [1 2 4 5]

	//Сравнение двух слайсов:

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true

}
