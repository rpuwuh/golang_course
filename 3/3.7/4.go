/*
	На стандартный ввод подаются данные о продолжительности периода
		(формат приведен в примере).
	Кроме того, вам дана дата в формате Unix-Time: 1589570165
		в виде константы типа int64
		(наносекунды для целей преобразования равны 0).
	Требуется считать данные о продолжительности периода,
		преобразовать их в тип Duration,
		а затем вывести (в формате UnixDate) дату и время,
		получившиеся при добавлении периода к стандартной дате.

	Небольшая подсказка:
		базовую дату необходимо явно перенести в зону UTC
		с помощью одноименного метода.

	Sample Input:
		12 мин. 13 сек.
	Sample Output:
		Fri May 15 19:28:18 UTC 2020
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

const ConstUnixTime int64 = 1589570165

func GetDate() time.Duration {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	splitedTime := strings.TrimRight(inputedTime, "\n")
	splitedTime = strings.ReplaceAll(splitedTime, "мин.", "m")
	splitedTime = strings.ReplaceAll(splitedTime, "сек.", "s")
	splitedTime = strings.ReplaceAll(splitedTime, " ", "")

	eventDuration, err := time.ParseDuration(splitedTime)
	if err != nil {
		panic(err)
	}

	return eventDuration
}

func main() {
	dur := GetDate()
	unixTime := time.Unix(ConstUnixTime, 0)
	unixTime = unixTime.Add(dur)
	fmt.Println(unixTime.Format("Mon Jan _2 15:04:05 MST 2006"))
}
