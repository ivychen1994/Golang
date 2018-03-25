package main

import "time"
import "fmt"

func main() {
//选择(select)可等待多个通道操作。将goroutine和channel与select结合是Go语言的一个强大功能。
//每个通道将在一段时间后开始接收值，以模拟阻塞在并发goroutines中执行的RPC操作。
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

