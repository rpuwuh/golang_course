/*
	Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get
		и напечатайте ответ сервера (только тело).
*/

package main

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// здесь ваш код

	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close() // закрываем тело ответа после работы с ним

	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s\n", data) // печатаем ответ как строку
}
