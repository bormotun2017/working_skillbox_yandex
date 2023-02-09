package main

//go mod init test //надо создавать go mod
//go test -v -count=1 .
//-v передает логи
//-count включает кэширование
import "testing"

func TestAdd(t *testing.T) { // описание функ тест. Add имя функ которую будем тестировать
	exp := 5 //задаем переменные
	x := 3
	y := 2
	res := add(x, y) //передаем их в функ add

	if res != exp { //если res не равно правильному ответу
		t.Fail() //фэйлим
	}
	exp = 5 //задаем переменные
	x = 2
	y = 2

	if res != exp { //если res не равно правильному ответу
		t.Logf("4 не равно 5")

		t.Fail() //фэйлим
	}

}
