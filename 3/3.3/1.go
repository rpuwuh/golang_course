/*
	Используем анонимные функции на практике.
	Внутри функции main (объявлять ее не нужно)
		вы должны объявить функцию вида func(uint) uint,
		которую внутри функции main
		в дальнейшем можно будет вызвать по имени fn.
	Функция на вход получает целое положительное число (uint),
		возвращает число того же типа в котором отсутствуют
		нечетные цифры и цифра 0.
	Если же получившееся число равно 0, то возвращается число 100.
	Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
	
	Пакет main объявлять не нужно!
	Вводить и выводить что-либо не нужно!
	Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

	Sample Input:
		727178
	Sample Output:
		28
*/

package main

import "fmt"

func main(){
	fn := func(input uint) uint {
		var res uint = 0
		
		var j uint = 1
		for i := 1; input > uint(i); i *= 10 {
			if ((input / uint(i) ) % 10) % 2 == 0 && ((input / uint(i) ) % 10) != 0 {
				res = ((input / uint(i) ) % uint(10)) * uint(j) + res
				j *= 10
			}
		}
		if (res == 0) {
			res = 100
		}
		return res
	}

	fmt.Println(fn(1234567890))
	fmt.Println(fn(13579))

}