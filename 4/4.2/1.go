/*
	Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP,
		считайте от сервера 3 сообщения,
		и выведите их в верхнем регистре.

	Рекомендуем использовать буфер в 1024 байта.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	message := make([]byte, 1024) // создадим буфер
	for i := 0; i < 3; i++ {
		n, err := conn.Read(message)
		if err != nil {
			log.Println(err)
		}
		str := string(message[:n]) // напечатаем сообщение
		str = strings.ToUpper(str)
		fmt.Print(str)
	}
}

func Server() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)

	}
	conn, err := listener.Accept()
	if err != nil {

		log.Println(err)

	}
	defer conn.Close()
	conn.Write([]byte("message"))
	time.Sleep(10)
	conn.Write([]byte("MesSaGe"))
	time.Sleep(10)
	conn.Write([]byte("MESSAGE"))
	conn.Close()
}

func main() {
	// ваш код
	go Server()

	time.Sleep(100 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	handleConnection(conn)
}
