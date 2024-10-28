/*
	Задача - перевот списка (reverse)
	Вам нужно реализовать функцию,
		которая принимает list и переворачивает порядок его элементов,
		так чтобы последний элемент стал первым, предпоследний — вторым,
		и так далее.
	Писать функцию main не нужно!
*/

package main

import (
	"container/list"
	"fmt"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	// Здесь ваш код

	res := list.New()
	for temp := l.Front(); temp != nil; temp = temp.Next() {
		res.PushFront(temp.Value)
	}
	return res
}

func printList(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
	fmt.Println()
}

func main() {

	task := list.New()

	for i := 0; i < 10; i++ {
		task.PushBack(i)
	}
	printList(task)
	// 0123456789

	printList(ReverseList(task))
	// 9876543210

}
