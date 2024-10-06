/*
	Ваша задача сделать проверку подходит ли пароль
		вводимый пользователем под заданные требования.
	Длина пароля должна быть не менее 5 символов,
		он может содержать только арабские цифры
		и буквы латинского алфавита.
	На вход подается строка-пароль.
	Если пароль соответствует требованиям - вывести "Ok",
		иначе вывести "Wrong password"
	
	Sample Input:
		fdsghdfgjsdDD1
	Sample Output:
		Ok
*/

package main

import (
    "fmt"
	"unicode"
)

func main() {
	passwordBytes := ""
    fmt.Scan(&passwordBytes)
    
	password := []rune(passwordBytes)
    
	passwordCheckFlag := true
	if len(password) < 5 {
		passwordCheckFlag = false
	}
	for i := 0; passwordCheckFlag && i < len(password); i++ {
		if (unicode.Is(unicode.Latin, (password[i])) || unicode.IsDigit(password[i])) == false {
			passwordCheckFlag = false
		}
	}

	if passwordCheckFlag {fmt.Println("Ok")} else {fmt.Println("Wrong password")}
}
