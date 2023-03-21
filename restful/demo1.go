package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// jsonplaceholder sitesi
// json to go sitesi
type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	// get işlemi
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Hata: ", err)
	}
	defer response.Body.Close()

	bodybytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodybytes)

	var todo ToDo
	json.Unmarshal(bodybytes, &todo)
	fmt.Println("struct yapısı şekli: ")
	fmt.Println(todo) //bize bu hali lazım
	fmt.Println("api den dönen yapısı: ")
	fmt.Println(bodyString)

}

//-------------------------------------------------------------

func Demo2() {
	//post işlemi
	todo := ToDo{23, 2, "Çikolatali pasta yapiacak", true}
	jsontodo, err := json.Marshal(todo)

	response2, err := http.Post("http://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsontodo))
	if err != nil {
		fmt.Println("hata: ", err)
	}
	defer response2.Body.Close()
	//post işleminden dönen cevabı okuma
	bodybytes, _ := ioutil.ReadAll(response2.Body)
	bodyString := string(bodybytes)

	var toDoresp ToDo
	json.Unmarshal(bodybytes, &toDoresp)
	fmt.Println("struct yapısı şekli: ")
	fmt.Println(toDoresp) //bize bu hali lazım
	fmt.Println("api den dönen yapısı: ")
	fmt.Println(bodyString)
}
