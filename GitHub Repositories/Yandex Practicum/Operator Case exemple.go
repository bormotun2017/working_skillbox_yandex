package main

import (
	"fmt"
	"time"
)

func main() {

	var dates string
	fmt.Println("введите дату рождения в формате YYYY-MM-DD")
	fmt.Scan(&dates)
	// преобразим стрингу в дату
	dt1, err := time.Parse("2006-01-02", dates)
	if err != nil {
		fmt.Println("ошибка ввода даты")
	}

	switch y := dt1.Year(); {
	case y >= 1946 && y < 1965:
		fmt.Println("Привет, бумер!")
	case y >= 1965 && y < 1981:
		fmt.Println("Привет, представитель X!")
	case y >= 1981 && y < 1997:
		fmt.Println("Привет, миллениал!")
	case y >= 1997 && y < 2013:
		fmt.Println("Привет, зумер!")
	case y >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Здравствуйте!")
	}
}
