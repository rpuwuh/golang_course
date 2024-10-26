/*
	Необходимо написать функцию
	func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

	Описание ее работы:
		n раз сделать следующее

	прочитать по одному числу из каждого из двух каналов in1 и in2,
		назовем их x1 и x2.
	вычислить f(x1) + f(x2)
	записать полученное значение в out
	Функция merge2Channels должна быть неблокирующей,
		сразу возвращая управление.
	Функция fn может работать долгое время,
		ожидая чего-либо или производя вычисления.

	Формат ввода:
		количество итераций передается через аргумент n.
		целые числа подаются через аргументы-каналы in1 и in2.
		функция для обработки чисел перед сложением
			передается через аргумент fn.
	Формат вывода:
		канал для вывода результатов передается через аргумент out.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Millisecond)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	} else {
		fmt.Println("порядок не нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Milliseconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

// пакет используется для проверки выполнения условия задачи, не удаляйте его

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var results []int = make([]int, n)
		var mutexForResults []sync.Mutex = make([]sync.Mutex, n)
		var mutexForInputCounter [2]sync.Mutex
		var InputCounter [2]int

		wg := new(sync.WaitGroup)
		for i := 0; i < 2*n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				//var channelIsNotClosed bool
				var x int
				{
					select {
					case x, _ = <-in1:
						var index, value int
						mutexForInputCounter[0].Lock()
						index = InputCounter[0]
						InputCounter[0]++
						mutexForInputCounter[0].Unlock()

						value = fn(x)

						mutexForResults[index].Lock()
						results[index] += value
						mutexForResults[index].Unlock()
						return
					case x, _ = <-in2:
						var index, value int
						mutexForInputCounter[1].Lock()
						index = InputCounter[1]
						InputCounter[1]++
						mutexForInputCounter[1].Unlock()

						value = fn(x)

						mutexForResults[index].Lock()
						results[index] += value
						mutexForResults[index].Unlock()
						return
					}
				}
			}()
		}
		wg.Wait()
		//time.Sleep(time.Duration(1 * time.Second))
		for i := 0; i < n; i++ {
			out <- results[i]
		}
	}()
}
