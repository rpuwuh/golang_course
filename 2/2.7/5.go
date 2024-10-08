/*
	Внимательно прочитайте ВСЕ условия и подсказки чтобы правильно решить задачу!
	Требуется вычислить период колебаний (t) математического маятника
		(мы округлили некоторые значения для удобства проверки),
		для этого нужно найти циклическую частоту колебания пружинного маятника (w),
		в формуле w встречается масса которую также нужно найти, все нужные формулы приведены ниже:
	Напишите три функции, каждая из которых будет выполнять конкретную формулу.
	Название функций обязательно должны соответствовать букве формулы: T(), W() и M().
	Для того чтобы найти t - необходимо сначала найти w, и т.д.
	Так что используйте результат функции W() в формуле функции T() -
		то-есть вызывайте функцию W() в T().
	Аналогично и с W(), M(). 
	𝑡 = 6 / 𝑤 , 𝑤 = (𝑘 / 𝑚) ^ 1/2, 𝑚 = 𝑝 ∗ 𝑣,
	ВАЖНО! Считайте, что пакет main уже объявлен,
		а также функция main() с вызовом ВАШЕЙ будущей функции T() уже есть.
	Несмотря на то, что тестирование будет через ввод-вывод,
		вам НЕ требуется вводить и выводить что-либо.
	Для подсчета используйте УЖЕ ВВЕДЕННЫЕ ГЛОБАЛЬНЫЕ переменные k,p,v ТИПА float64 !!!
	
	Пакет math уже импортирован!
	Напоминаю: корень (sqrt) можно найти с помощью пакета "math", например:
	​fmt.Println(math.Sqrt(9))
	// результат: 3
*/

func M() float64 {
    return (p * v)
}

func W() float64 {
    return (math.Sqrt(k / M()))
}

func T() float64 {
    return (6 / W())
}
