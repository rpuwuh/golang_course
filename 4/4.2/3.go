/*
	Считайте с консоли две переменные,
		сначала name, затем age.

	Сделайте HTTP запрос на http://127.0.0.1:8080/hello ,
		передав name и age в query параметрах,
		затем напечатайте ответ сервера (только тело).
*/

package main

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type paramsStruct struct {
	name string
	age  int
}

func main() {
	// здесь ваш код

	paramsRead := paramsStruct{"", 0}
	fmt.Scan(&paramsRead.name, &paramsRead.age)

	// Создаем URL с параметрами
	baseURL := "http://127.0.0.1:8080/hello"
	paramsURL := url.Values{}
	paramsURL.Add("name", paramsRead.name)
	paramsURL.Add("age", strconv.Itoa(paramsRead.age))

	fullURL := baseURL + "?" + paramsURL.Encode()

	// Отправляем GET-запрос
	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Ошибка при отправке GET-запроса:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s\n", data) // печатаем ответ как строку
}
