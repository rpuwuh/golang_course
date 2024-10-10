/*
	Для решения данной задачи вам понадобится пакет strconv,
		возможно использовать пакеты strings или encoding/csv,
		или даже bufio - вы не ограничены в выборе способа решения задачи.
	В решениях мы поделимся своими способами решения этой задачи,
		предлагаем вам сделать то же самое.
	В привычных нам редакторах электронных таблиц присутствует
		удобное представление числа с разделителем разрядов в виде пробела,
		кроме того в России целая часть от дробной отделяется запятой.
	Набор таких чисел был экспортирован в формат CSV,
		где в качестве разделителя используется символ ";".
	
	На стандартный ввод вы получаете 2 таких вещественных числа,
		в качестве результата требуется вывести частное от деления первого числа
		на второе с точностью до четырех знаков после "запятой"
		(на самом деле после точки,
			результат не требуется приводить к исходному формату).
	P.S. небольшое отступление, связанное с чтением из стандартного ввода.
	Кто-то захочет использовать для этого пакет bufio.Reader.
	Это вполне допустимый вариант, но если вы
		резонно обрабатываете ошибку метода ReadString('\n'),
		то получаете ошибку EOF, а точнее (io.EOF - End Of File).
	На самом деле это не ошибка, а состояние, означающее,
		что файл (а os.Stdin является файлом) прочитан до конца.
	Чтобы ошибка была обработана правильно, вы можете поступить так:
		if err != nil && err != io.EOF {
			...
		}
	
	Sample Input:
		1 149,6088607594936;1 179,0666666666666
	Sample Output:
		0.9750
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
	"io"
)

func readStrings() (strs []string){
	strTmp, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil || err != io.EOF {
		strTmp = strings.Replace(strTmp, " ", "", -1)
		strTmp = strings.Replace(strTmp, "\n", "", -1)
		strTmp = strings.Replace(strTmp, ",", ".", -1)
		strs = strings.Split(strTmp, ";")
	} else {
		fmt.Println("An error occured")
	}

	return 
}

func getFloats(strs []string) (res1, res2 float64) {
	if len(strs) == 2 {
		res1, _ = strconv.ParseFloat(strs[0], 64)
		res2, _ = strconv.ParseFloat(strs[1], 64)
		return
	} else {
		return 0, 1
	}
}

func main(){
	float1, float2 := getFloats(readStrings())
	fmt.Printf("%.4f\n", float1 / float2)
}
