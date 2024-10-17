/*
	Поиск файла в заданном формате и его обработка
	Данная задача поможет вам разобраться в пакете encoding/csv
		и path/filepath, хотя для решения может быть использован
		также пакет archive/zip
		(поскольку файл с заданием предоставляется именно в этом формате).

	В тестовом архиве, который вы можете скачать из нашего репозитория
		на github.com, содержится набор папок и файлов.
	Один из этих файлов является файлом с данными в формате CSV,
		прочие же файлы структурированных данных не содержат.
	
	Требуется найти и прочитать этот единственный файл
		со структурированными данными
		(это таблица 10х10, разделителем является запятая),
		а в качестве ответа необходимо указать число,
		находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/csv"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
			return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
		}
		
	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	fileDesc, _ := os.Open(path)
	r := csv.NewReader(fileDesc)
	record, _ := r.ReadAll()
	if len(record) == 10 {
	   fmt.Println(record[4][2])
	   fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	}
	return nil
}
	
func main(){
	const root = "./task" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}