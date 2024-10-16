/*
	bufio vs fmt
	
	Ранее в рамках этого курса при решении задач требовалось прочитать что-то
		со стандартного ввода и вывести результат соответственно
		в стандартный вывод.
	При этом кто-то использовал пакет fmt, а кто-то - bufio + os.
	Все эти пакеты имеют свои особенности,
		поэтому в этой задаче мы попробуем решить знакомую нам проблему
		с помощью пакетов, которые кто-то мог до этого момента и не применять:
		bufio + os + strconv.

	Задача состоит в следующем:
	На стандартный ввод подаются целые числа в диапазоне 0-100,
		каждое число подается на стандартный ввод с новой строки
		(количество чисел не известно).
	Требуется прочитать все эти числа 
			и вывести в стандартный вывод их сумму.

	Несколько подсказок:
		для чтения вы можете использовать типы bufio.Reader и bufio.Scanner,
			а для записи - bufio.Writer.
		При чтении помните об ошибке io.EOF.
		Конвертирование данных из строки в целое число и обратно
			осуществляется функциями Atoi и Itoa
			из пакета strconv соответственно.
		Чтение производится из стандартного ввода (os.Stdin),
			а запись - в стандартный вывод (os.Stdout).
	
	Все указанные в тексте задачи пакеты (strconv, bufio, os, io), уже импортированы (другие использовать нельзя), package main объявлен.

	Sample Input:
		33
		47
		12
		79
		15
	Sample Output:
		186
*/

package main

import (
	"io"
	"os"
	"bufio"
	"strconv"
)

// package main уже объявлен. 
func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
    // Другие использовать нельзя.
	var sum int
	
	reader := bufio.NewReader(os.Stdin)
	
	strBuf, err := reader.ReadString('\n')
	for err == nil {
		intTmp := 0
		if strBuf[len(strBuf) -1] == '\n'{
			intTmp, _ = strconv.Atoi(strBuf[:len(strBuf) - 1])
		}
		sum += intTmp

		strBuf, err = reader.ReadString('\n')
	}
	if err == nil || err == io.EOF{
		intTmp := 0
    	intTmp, _ = strconv.Atoi(strBuf[:len(strBuf)])

		sum += intTmp

		writer := bufio.NewWriter(os.Stdout)
		writer.WriteString(string(strconv.Itoa(sum)) )
        writer.Flush()
	}
}
