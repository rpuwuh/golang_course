/*
	Напишите программу, 
	которая в последовательности чисел
	находит сумму двузначных чисел, кратных 8.
	Программа в первой строке получает на вход
	число n - количество чисел в последовательности,
	во второй строке -- n чисел, входящих в данную последовательность.

	Sample Input:
		5
		38 24 800 8 16
	Sample Output:
		40
*/

package main

import "fmt"

func main() {
	var n, sum, tempInput int
    fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&tempInput)
		if (tempInput > 9 && tempInput < 100 && tempInput % 8 == 0) {
			sum += tempInput
		}
	}
	fmt.Println(sum)
}
