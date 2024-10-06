/*
	На вход дается строка,
		из нее нужно сделать другую строку,
		оставив только нечетные символы (считая с нуля)
	
	Sample Input:
		ihgewlqlkot
	Sample Output:
		hello
*/

package main

import "fmt"

func main() {
    text1, text2 := "", ""

    fmt.Scan(&text1)
    
	for i := 1; i < len(text1); i += 2{
		text2 += string(text1[i])
	}

	fmt.Println(text2)
}
