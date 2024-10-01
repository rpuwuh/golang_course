/*
	Входные данные
		Даны три натуральных числа a, b, c.
		Определите, существует ли треугольник с такими сторонами.
	Выходные данные
		Если треугольник существует, выведите строку "Существует",
			иначе выведите строку "Не существует".
		Строку выводите без кавычек.
	
	Sample Input:
		4 5 6
	Sample Output:
		Существует
*/

package main 

import "fmt"

func main() {
	var a, b, c uint64

	fmt.Scan(&a, &b, &c)

	if (a + b > c) && (a + c > b) && (b + c > a) {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}