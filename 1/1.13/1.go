/*
	Дано трехзначное число. Найдите сумму его цифр. 
	
	Формат входных данных
		На вход дается трехзначное число.
	Формат выходных данных
		Выведите одно целое число - сумму цифр введенного числа.
	
	Sample Input:
		745
	Sample Output:
		16
*/

package main

import "fmt"

func main () {
	var sum, x int

	fmt.Scan(&x)
	
	for x != 0 {
		sum += x % 10
		x /= 10
	}

	fmt.Println(sum)
}
