/*
	Вклад в банке составляет x рублей.
	Ежегодно он увеличивается на p процентов,
	после чего дробная часть копеек отбрасывается.
	Каждый год сумма вклада становится больше.
	Определите, через сколько лет вклад составит не менее y рублей.
	
	Входные данные
		Программа получает на вход три натуральных числа: x, p, y.
	Выходные данные
		Программа должна вывести одно целое число.
	Sample Input:
		100
		10
		200
	Sample Output:
		8
*/

package main

import "fmt"

func main() {
	var x, p, y, years int

	fmt.Scan(&x, &p, &y)
	
	for x < y {
		x *= 100 + p
		x /= 100
		years++
	}

	fmt.Println(years)
}
