/*
	На стандартный ввод подается строковое представление двух дат,
		разделенных запятой (формат данных смотрите в примере).
	Необходимо преобразовать полученные данные в тип Time,
		а затем вывести продолжительность периода между меньшей
		и большей датами.

	Sample Input:
		13.03.2018 14:00:15,12.03.2018 14:00:15
	Sample Output:
		24h0m0s
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func GetDate() []time.Time {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	splitedTime := strings.Split(inputedTime, ",") //splitedTime[0] && splitedTime[1] соответственно
	splitedTime[1] = strings.TrimRight(splitedTime[1], "\n")

	// func Parse(layout, value string) (Time, error)
	// парсит дату и время в строковом представлении
	times := []time.Time{}
	eventTime, err := time.Parse("02.01.2006 15:04:05", splitedTime[0])
	if err != nil {
		panic(err)
	}
	times = append(times, eventTime)
	eventTime, err = time.Parse("02.01.2006 15:04:05", splitedTime[1])
	if err != nil {
		panic(err)
	}
	times = append(times, eventTime)

	return times
}

func main() {
	times := GetDate()

	if len(times) < 2 {
		panic("critical error")
	}
	if times[0].Compare(times[1]) > 0 {
		fmt.Println(time.Since(times[0]).Round(time.Second) - time.Since(times[1]).Round(time.Second))
	} else {
		fmt.Println(time.Since(times[1]).Round(time.Second) - time.Since(times[0]).Round(time.Second))
	}
}
