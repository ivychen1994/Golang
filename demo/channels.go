package main

import "fmt"

func main() {
//通道是连接并发goroutine的管道。可以从一个goroutine向通道发送值，并在另一个goroutine中接收到这些值。
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

