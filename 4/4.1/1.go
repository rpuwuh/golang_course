/*
	Задача FIFO
	Реализуйте функции для очереди
		FIFO (First In, First Out (ПЕРВЫЙ пришел, ПЕРВЫЙ вышел))
		с помощью списков.
	Должны быть данные функции:
	Push (добавление элемента)
	Pop (удаление элемента и его возврат)
	printQueue (печать очереди в одну строку без пробелов)

	Функцию main писать не нужно!
	Писать код вне функций не нужно.
*/

package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushFront(elem)
}

func Pop(queue *list.List) interface{} {
	if queue.Back() == nil {
		return nil
	}
	res := queue.Back().Value
	queue.Remove(queue.Back())
	return res
}

func printQueue(queue *list.List) {
	for temp := queue.Back(); temp != nil; temp = temp.Prev() {
		fmt.Printf("%v", temp.Value)
	}
	fmt.Println()
}

func main() {

	queue := list.New()
	fmt.Println("Poped element: ", Pop(queue))

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123

	fmt.Println("Poped element: ", Pop(queue))

	printQueue(queue) // в ту же строку: 23

}
