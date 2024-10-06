/*
	На вход подается строка.
	Нужно определить, является ли она палиндромом.
	Если строка является палиндромом - вывести Палиндром
		иначе - вывести Нет.
	Все входные строчки в нижнем регистре.
	Палиндром — буквосочетание, слово или текст,
		одинаково читающееся в обоих направлениях
		(например, "топот", "око", "заказ").
	
	Sample Input:
		топот
	Sample Output:
		Палиндром
*/
package main

import "bufio"
import "os"
import "fmt"
import "strings"

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	textRunes := []rune(text)

	isPalindrome := true;
	for i := 0; isPalindrome && i < (len(textRunes) - 1) / 2; i++{
		if (textRunes[i] != textRunes[len(textRunes) - 1 - i]){
			isPalindrome = false
		}
	}
	if (isPalindrome == true){
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
