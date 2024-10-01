/*
	Из натурального числа удалить заданную цифру.
	
	Входные данные
		Вводятся натуральное число и цифра, которую нужно удалить.
	Выходные данные
		Вывести число без заданных цифр.
	
	Sample Input:
		38012732
		3
	Sample Output:
		801272
*/

package main

import "fmt"

func main(){
	var num1, exponent, digit int = 0, 1, 0

	fmt.Scan(&num1, &digit)

	for exponent < num1	{
		exponent *= 10
	}
	exponent /= 10

	for exponent != 0 {
		if ((num1 / exponent ) % 10 != digit) {
			fmt.Printf("%d", (num1 / exponent ) % 10)
        }
		exponent /= 10
	}
   
	fmt.Println()
}
