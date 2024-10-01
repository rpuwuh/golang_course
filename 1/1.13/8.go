/*
	Количество минимумов
	Найдите количество минимальных элементов в последовательности.
	
	Входные данные
		Вводится натуральное число N, а затем N целых чисел последовательности.
	Выходные данные
		Выведите количество минимальных элементов последовательности.
	Sample Input:
		3
		21
		11
		4
	Sample Output:
		1
*/

package main 

import "fmt"

func main() {
	var tmp, n, minValue, minValueCounter int

	fmt.Scan(&n)

	if (n > 0) {
		fmt.Scan(&tmp)
		minValue = tmp
		minValueCounter = 1
	}
	for i := 1; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp == minValue{
			minValueCounter++
		} else if tmp < minValue {
			minValue = tmp
			minValueCounter = 1
		}
	}

	fmt.Println(minValueCounter)
}
