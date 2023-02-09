package main

import "fmt"

//Выведите перечень деликатесов — продуктов дороже 500 рублей.
//Заказ выдан слайсом []string{"хлеб", "буженина", "сыр", "огурцы"}. Посчитайте стоимость заказа.

func main() {
	price := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	arr := []string{"хлеб", "буженина", "сыр", "огурцы"}

	fmt.Println("прайс на продукты")

	for key, data := range price {
		fmt.Println(key, data)

	}
	fmt.Println()

	fmt.Println("продукты дороже 500 из списка ", arr, "\n")

	for key, data := range price {

		for i := 0; i < len(arr); i++ {

			if key == arr[i] && data > 500 {

				fmt.Println(key, data)

			}
		}

	}

}
