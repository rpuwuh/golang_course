/*
	Поэтапный поиск данных
	Данная задача в основном ориентирована на изучение типа bufio.Reader,
		поскольку этот тип позволяет считывать данные постепенно.
	
	В тестовом файле, который вы можете скачать из нашего репозитория
		на github.com, содержится длинный ряд чисел, разделенных символом ";".
	Требуется найти, на какой позиции находится число 0
		и указать её в качестве ответа.
	Требуется вывести именно позицию числа,
		а не индекс (то-есть порядковый номер, нумерация с 1).
	
	Например:  12;234;6;0;78 :
	Правильный ответ будет порядковый номер числа: 4
*/

package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)

func main(){
    file, err := os.Open("task.data") 
    if err != nil { 
         fmt.Println("Unable to open file:", err) 
         return
    } 
    defer file.Close() 
  
    reader := bufio.NewReader(file) 
	i := 1
    for { 
        line, err := reader.ReadString(';') 
        if err != nil { 
            if err == io.EOF { 
                break
            } else { 
                fmt.Println(err) 
                return
            } 
        }
		if line == "0;" {
			fmt.Println(i) 
		} 
		i++
    }
	//fmt.Println("Last i:", i) 
}
