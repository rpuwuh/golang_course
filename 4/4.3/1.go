/*
	Напишите веб сервер, который по пути /get отдает текст "Hello, web!".
	Порт должен быть :8080.
*/

package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	// здесь ваш код

	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/get", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
