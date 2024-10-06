/*
	Дается строка.
	Нужно удалить все символы,
		которые встречаются более одного раза
		и вывести получившуюся строку
	
	Sample Input:
		zaabcbd
	Sample Output:
		zcd
*/

package main

import "fmt"

func main() {
    text1, text2 := "", ""
	slice1 := [256]bool{}
	slice2 := [256]bool{}
    
	fmt.Scan(&text1)
    
	for i := 0; i < len(text1); i++ {
		if (slice1[text1[i]] == false) {
			slice1[text1[i]] = true
		} else if (slice2[text1[i]] == false) {
			slice2[text1[i]] = true
		}
	}
	for i := 0; i < len(text1); i++ {
		if (slice2[text1[i]] == false) {
			text2 += string(text1[i])
		}
	}

	fmt.Println(text2)
}
