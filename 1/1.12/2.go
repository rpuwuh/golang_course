/*
	Напишите программу, принимающая на вход число 𝑁(𝑁≥4),
	а затем N целых чисел-элементов среза.
	На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
	
	Sample Input:
		5
		41 -231 24 49 6
	Sample Output:
		49
*/

package main
import "fmt"

func main() {
    // здесь должен быть ваш код
    
	var capacity int
	fmt.Scan(&capacity)

	slice := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println(slice[3])
}
