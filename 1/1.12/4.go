/*
	Дан массив, состоящий из целых чисел.
	Нумерация элементов начинается с 0.
	Напишите программу, которая выведет элементы массива,
		индексы которых четны (0, 2, 4...).
	
	Входные данные
	Сначала задано число N —
		количество элементов в массиве (1≤N≤100).
	Далее через пробел записаны N чисел —
		элементы массива.
	Массив состоит из целых чисел.
	
	Выходные данные
	Необходимо вывести все элементы массива с чётными индексами.
	
	Sample Input:
		6
		4 5 3 4 2 3
	Sample Output:
		4 3 2
*/

package main
import "fmt"

func main()  {
	var capacity int
	fmt.Scan(&capacity)

	slice := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Printf("%d", slice[0])
	for i := 2; i < capacity; i+=2 {
		fmt.Printf(" %d", slice[i])
	}

}
