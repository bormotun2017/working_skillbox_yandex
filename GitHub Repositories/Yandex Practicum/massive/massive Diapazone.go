package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr2 := make([]int, 0, 0)
	arr3 := make([]int, 0, 0)

	arr2 = append(arr2, arr[5:]...) // добавляем в массив множество
	arr2 = append(arr2, arr[:5]...) // другое множество

	arr3 = append(arr3, arr[:]...)  // добавляем весь массив
	arr3 = append(arr3, arr2[:]...) // соединяем

	slice := append(arr[0:3], 4, 4, 4) // добавляем 3 4ки после 4го эл массива

	fmt.Println(slice) //0 1 2 4 4 4
	fmt.Println(arr)   // 0 1 2 4 4 4 6 7 8 9

	fmt.Println()

	fmt.Println(arr2) // 5 6 7 8 9 0 1 2 3 4
	fmt.Println(arr3) // 0 1 2 3 4 5 6 7 8 9 5 6 7 8 9 0 1 2 3 4

}
