/*
	Напиши веб сервер (порт :3333) -
		счетчик который будет обрабатывать
		GET (/count) и POST (/count) запросы:
	GET:  возвращает счетчик
	POST: увеличивает ваш счетчик на значение (с ключом "count"),
		которое вы получаете из формы,
		но если пришло НЕ число то нужно ответить клиенту:
			"это не число" со статусом http.StatusBadRequest (400).
*/

package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

var counter int

func handlerCount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(counter)))
	} else if r.Method == "POST" {
		r.ParseForm()
		increment, err := strconv.Atoi(r.FormValue("count"))
		if err == nil {
			w.WriteHeader(http.StatusOK)
			counter += increment
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	// порт :3333

	http.HandleFunc("/count", handlerCount)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
