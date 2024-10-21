/*
	На стандартный ввод подается строковое представление даты
		и времени определенного события в следующем формате:
		2020-05-15 08:00:00
	Если время события до обеда (13-00), то ничего менять не нужно,
		достаточно вывести дату на стандартный вывод в том же формате.
	Если же событие должно произойти после обеда,
		необходимо перенести его на то же время на следующий день,
		а затем вывести на стандартный вывод в том же формате.
	Sample Input:
		2020-05-15 08:00:00
	Sample Output:
		2020-05-15 08:00:00
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var bufStr string

	bufStr, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	bufStr = strings.TrimSpace(bufStr)

	// func Parse(layout, value string) (Time, error)
	// парсит дату и время в строковом представлении
	eventTime, err := time.Parse("2006-01-02 15:04:05", bufStr)
	if err != nil {
		panic(err)
	}
	if eventTime.Hour() >= 13 {
		eventTime = eventTime.Add(time.Hour * 24)
	}
	fmt.Println(eventTime.Format("2006-01-02 15:04:05")) //
}
