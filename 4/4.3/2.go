/*
	Напишите веб-сервер который по пути /api/user приветствует пользователя:
	Принимает и парсит параметр name и делает ответ "Hello,<name>!"
	Пример: /api/user?name=Golang
	Ответ: Hello,Golang!

	порт :9000
*/

package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("name") {
		w.Write([]byte("Hello," + r.URL.Query().Get("name") + "!"))
	}
}

func main() {
	// здесь ваш код

	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/api/user", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
