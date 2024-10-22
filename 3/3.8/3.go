/*
	Напишите элемент конвейера (функцию),
		что запоминает предыдущее значение и отправляет значения
		на следующий этап конвейера только если оно отличается от того,
		что пришло ранее.

	Ваша функция должна принимать два канала - inputStream и outputStream,
		в первый вы будете получать строки,
		во второй вы должны отправлять значения без повторов.
	В итоге в outputStream должны остаться значения,
		которые не повторяются подряд.
	Не забудьте закрыть канал ;)

	Функция должна называться removeDuplicates()
	Выводить или вводить ничего не нужно!
*/

// Пакет и функция main уже объявлены.
// Выводить или вводить ничего не нужно!

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var str2 string
	var firstIteraion bool = true

	for str1 := range inputStream {
		if str1 != str2 || firstIteraion == true {
			outputStream <- str1
		}
		firstIteraion = false
		str2 = str1
	}
	close(outputStream)
}
