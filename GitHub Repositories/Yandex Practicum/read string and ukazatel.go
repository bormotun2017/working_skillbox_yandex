package main

import (
	"bufio"
	"fmt"
	"os"
)

//программа выводит то что ввел пользователь

func f(cnt *int) { // передача указателя
	*cnt++

}

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		read, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if read == "exit\r\n" {
			break
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)

	}
}
