package main

import (
	"fmt"
)

//Создайте слайс и заполните его числами от 1 до 100.
//Оставьте в слайсе первые и последние 10 элементов и разверните слайс в обратном порядке.
//Подумайте, можно ли подобным образом развернуть строку?

func main() {

	var slice []int
	var slice2 []int

	for i := 1; i <= 100; i++ {

		slice = append(slice, i)

	}

	fmt.Println("Slice")
	fmt.Println(slice, "\n")

	//	slice = append(slice[len(slice)-10:], slice[:10]...) // изящно

	slice2 = append(slice[len(slice)-10:], slice[:10]...)
	//slice2 = append(slice2, slice[90:]...)
	//slice2 = append(slice2, slice[:10]...)

	fmt.Println("Slice2")
	fmt.Println(slice2, "\n")

	// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
	// Лучше создать слайс рун из строки и развернуть его
	// Например, так:

	fmt.Println("------------------------------------------------")

	input := "The quick brown 狐 jumped over the lazy 犬"

	fmt.Println(input, "\n")

	fmt.Println()
	// Get Unicode code points.
	n := 0
	// создаём слайс рун
	runes := make([]rune, len(input))
	// добавляем руны в слайс
	for _, r := range input {
		runes[n] = r
		n++
	}
	// убираем лишние нулевые руны
	runes = runes[0:n]
	// разворачиваем
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// преобразуем в строку UTF-8.
	output := string(runes)

	fmt.Println("разворачиваем строку")
	fmt.Println(output)
}
