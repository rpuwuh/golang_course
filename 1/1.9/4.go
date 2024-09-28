/*
	Определите является ли билет счастливым.
	Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних. 

    Формат входных данных
    На вход подается номер билета - одно шестизначное  число.

    Формат выходных данных
    Выведите "YES", если билет счастливый, в противном случае - "NO".
*/

package main

import "fmt"

func main() {
	var a, b int
    fmt.Scan(&a)
    
    b = a / 100 % 10 + a / 10 % 10 + a % 10
    a = a / 100000 + a / 10000 % 10 + a / 1000 % 10
    
    if a == b {
		fmt.Println("YES")
    } else {
		fmt.Println("NO")
    }
}
