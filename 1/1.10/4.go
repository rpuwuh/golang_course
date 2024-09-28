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
	var n, maxValue, tempInput int

	fmt.Scan(&tempInput)
	if (tempInput > maxValue) {
		maxValue = tempInput
		n = 1
	} else if (tempInput == maxValue) {
		n++
	}
	for tempInput != 0 {
		fmt.Scan(&tempInput)
		if (tempInput > maxValue) {
			maxValue = tempInput
			n = 1
		} else if (tempInput == maxValue) {
			n++
		}
	}
	fmt.Println(n)
}
