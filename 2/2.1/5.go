/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */

 func sumInt (a ...int) (n int, sum int) {
    for i := 0; i < len(a); i++ {
		n++
        sum += a[i]
	}
    return
}
