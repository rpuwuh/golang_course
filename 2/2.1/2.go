/*
	Напишите функцию,
		находящую наименьшее из четырех введённых
		в этой же функции чисел.
	
	Входные данные
		Вводится четыре числа.
	Выходные данные
		Необходимо вернуть из функции наименьшее из 4-х данных чисел
	Sample Input:
		4 5 6 7
	Sample Output:
		4
*/

func minimumFromFour() int {
    //print your code here
    var x, min int
    fmt.Scan(&min)
    for i := 1; i < 4; i++ {
        fmt.Scan(&x)
        if x < min {
            min = x
        }
    }
    return min
}
