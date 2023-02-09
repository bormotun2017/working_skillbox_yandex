package main

import "reflect"

//К сожалению, ни Swagger-описания, ни статьи с API-ответом в любимом сервисе заметок — нет.
//Опишите данный объект в виде структуры на Go, в учебных целях отбросив «так делать нельзя» и «это не дело».
//На входе есть строка с сырыми данными, требуется написать функцию её десериализации:

type T struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	// код десериализации
}

func main() {


	rawResp := T{Header: struct {
		Code    int
		Message string
	}{Code:0 , Message: ""}
	Data: []struct {

	}{},
	}


}
