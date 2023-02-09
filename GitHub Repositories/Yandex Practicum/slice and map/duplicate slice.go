package main

import (
	"fmt"
)

//Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:

func RemoveDuplicates(input []string) []string {

	result := make([]string, 0)
	resultTotal := make([]string, 0)

	for index, data := range input {

		for i := index + 1; i < len(input); i++ {

			if data == input[i] {
		
				result = append(input[:i], input[i+1:]...)
			}

		}

	}
	// пришлось удалить последний эл, в цикле не хочет удалять
	resultTotal = append(resultTotal, result[:len(result)-1]...)

	return resultTotal

}

func main() {

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	result := RemoveDuplicates(input)

	fmt.Println(result)

}
