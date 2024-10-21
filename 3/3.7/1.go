/*
	На стандартный ввод подается строковое представление даты и времени
		в следующем формате:
		1986-04-16T05:20:00+06:00
	Ваша задача конвертировать эту строку в Time,
		а затем вывести в формате UnixDate:
		Wed Apr 16 05:20:00 +0600 1986
	Для более эффективной работы рекомендуется ознакомиться
		с документацией о содержащихся в модуле time константах.

	Sample Input:
		1986-04-16T05:20:00+06:00
	Sample Output:
		Wed Apr 16 05:20:00 +0600 1986
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var bufStr string

	fmt.Scan(&bufStr)

	// func Parse(layout, value string) (Time, error)
	// парсит дату и время в строковом представлении
	time, err := time.Parse("2006-01-02T15:04:05Z07:00", bufStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Format("Mon Jan _2 15:04:05 MST 2006")) // UnixDate
}
