package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//curl -POST -d '{"name": "Igor", "age": 50}' http://localhost:8080/create
//curl -X POST -d "{\"name\": \"Alex\", \"age\": 20}" http://localhost:8080/create  создает нового пользователя
//{ это мы будем ожидать такой json
//	"name" :"some name",
//	"age":20,
//}

type User struct { //созданим структуру User
	Name string `json:"name"`
	Age  int    `json:"age"` //аннотация для того, чтобы json понят какие поля соответтвуют каким тэгам
}

func (u *User) toString() string { //создаем метод
	return fmt.Sprintf("name is %s nad age is %d \n", u.Name, u.Age)
}

type service struct { //структура для пользователя. хранит имя и возраст
	store map[string]*User
}

func main() {
	mux := http.NewServeMux() //создаем http сервер
	//функ которая будет обрабатывать наш запрос по опрю паттерну
	//	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	//		rw.Write([]byte("i'm alive"))
	//	})
	srv := service{make(map[string]*User)} //создаем сервис, мапу
	mux.HandleFunc("/get", srv.GetAll)
	mux.HandleFunc("/create", srv.Create)
	http.ListenAndServe("localhost:8080", mux) //типа читаем с порта 8080 и передаем туда
}

// создаем метод create
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //если метод POST
		content, err := ioutil.ReadAll(r.Body) //считываем тело
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //обработка ошибки
			w.Write([]byte(err.Error()))                  //пишем текст ошибки
			return
		}
		defer r.Body.Close() // закрываем body

		var u User                                          //создаем переменную типа User
		if err := json.Unmarshal(content, &u); err != nil { //
			w.WriteHeader(http.StatusInternalServerError) //обработка ошибки
			w.Write([]byte(err.Error()))                  //пишем текст ошибки
			return
		}
		//	splittedContent := strings.Split(string(content), " ") //сплитуем контент по пробелу
		s.store[u.Name] = &u //записываем в мапу контент

		w.WriteHeader(http.StatusCreated)             //пишем статус 201
		w.Write([]byte("User was created " + u.Name)) //
		return
	}
	w.WriteHeader(http.StatusBadGateway) //если не POST, то пишем что такой страницы не существует
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //если get
		response := ""                 //создаем пустую строку
		for _, user := range s.store { //читаем мапу
			response += user.toString() //заполняем строку значениями из мапы
		}
		w.WriteHeader(http.StatusOK) //пишем статус 200
		w.Write([]byte(response))    // пишем переменную
		return

	}
	w.WriteHeader(http.StatusBadRequest) //если метод не GET пишем статус 400

}
