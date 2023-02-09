package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//Напишите код, который будет сериализовывать структуру в json-строку следующего вида:
//{
//  "Имя": "Aлекс",
//  "Почта": "alex@yandex.ru"
//}

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {

	a := Person{
		Name:        "Алекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}

	b, err := json.Marshal(a) // переводим в json

	if err != nil {
		fmt.Println("error", err)
	}

	os.Stdout.Write(b)
}
