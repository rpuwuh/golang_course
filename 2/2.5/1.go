/*
	На вход подается строка.
	Нужно определить, является ли она правильной или нет.
	Правильная строка начинается с заглавной буквы
		и заканчивается точкой.
	Если строка правильная - вывести Right иначе - вывести Wrong

	Маленькая подсказка:
		fmt.Scan() считывает строку до первого пробела,
		вы можете считать полностью строку за раз с помощью bufio:
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	Sample Input:
		Быть или не быть.
	Sample Output:
		Right
*/

package main

import "bufio"
import "os"
import "fmt"
import "unicode"
import "strings"

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if (err != nil){
		fmt.Print("Wrong\n")
	} else {
		text = strings.TrimSuffix(text, "\n")
		textRunes := []rune(text)

		if (unicode.IsUpper(textRunes[0]) && textRunes[len(textRunes) - 1] == '.'){
			fmt.Print("Right\n")
		} else {
			fmt.Print("Wrong\n")
		}
	}
}
