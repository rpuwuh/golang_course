/*
	Данная задача ориентирована на реальную работу с данными в формате JSON.
	Для решения мы будем использовать справочник ОКВЭД
		(Общероссийский классификатор видов экономической деятельности),
		который был размещен на web-портале data.gov.ru.
	Необходимая вам информация о структуре данных содержится в файле
		structure-20190514T0000.json, а сами данные, которые вам потребуется
		декодировать, содержатся в файле data-20190514T0100.json.
	Файлы размещены в нашем репозитории на github.com.
	Для того, чтобы показать, что вы действительно смогли декодировать документ
		вам необходимо в качестве ответа записать сумму полей global_id
		всех элементов, закодированных в файле.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type OKVED2 []struct {
	GlobalID       int    `json:"global_id"`
	SystemObjectID string `json:"system_object_id"`
	SignatureDate  string `json:"signature_date"`
	Razdel         string `json:"Razdel"`
	Kod            string `json:"Kod"`
	Name           string `json:"Name"`
	Idx            string `json:"Idx"`
}

func readFile() (jsonString []byte, errRes error) {
	jsonString = make([]byte, 0, 1200000)

	{
		file, err := os.Open("task.data")
		if err != nil {
			fmt.Println("Unable to open file:", err)
			return nil, err
		}
		//fmt.Println("File: ok")
		defer file.Close()

		reader := bufio.NewReader(file)
		i := 0
		for {
			line, err := reader.ReadString('}')
			if err != nil {
				if err == io.EOF {
					jsonString = append(jsonString, []byte(line)...)
					break
				} else {
					fmt.Println(err)
					return
				}
			}
			jsonString = append(jsonString, []byte(line)...)
			i++
		}
		//fmt.Println("Read strings:", i)
	}
	return jsonString, errRes
}

func main() {
	jsonString, err := readFile()
	if err != nil {
		os.Exit(1)
	}

	var sumIDs int
	var parsedStruct OKVED2

	if err = json.Unmarshal(jsonString, &parsedStruct); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(parsedStruct); i++ {
		sumIDs += parsedStruct[i].GlobalID
	}
	fmt.Println(sumIDs)
}
