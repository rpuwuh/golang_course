/*
	Вам необходимо написать функцию calculator следующего вида:
	func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
	В качестве аргумента эта функция получает два канала только для чтения,
		возвращает канал только для чтения.
	Через канал arguments функция получит ряд чисел, а через канал done -
		сигнал о необходимости завершить работу.
	Когда сигнал о завершении работы будет получен,
		функция должна в выходной (возвращенный) канал
		отправить сумму полученных чисел.
	Функция calculator должна быть неблокирующей, сразу возвращая управление.
	Выходной канал должен быть закрыт после
		выполнения всех оговоренных условий,
		если вы этого не сделаете, то превысите предельное время работы.
*/

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	ret := make(chan int)
	sum := 0

	go func(ret chan int) {
		defer close(ret)
		for {
			select {
			case v, ok := <-arguments:
				if ok {
					sum += v
				}
			case <-done:
				ret <- sum
				return
			}
		}
	}(ret)

	return (ret)
}