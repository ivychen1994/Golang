package main

import (
	"fmt"
	"time"
)

func main() {
	//通道是连接并发goroutine的管道.可以从一个goroutine向通道发送值，并在另一个goroutine中接收到这些值。

	/*1.
	//创建一个新通道，通道由输入的值传入。
	messages := make(chan string)

	//使用通道 <- 语法将值发送到通道。 这里从一个新的goroutine发送“ping”到在上面的消息通道
	go func() { messages <- "ping" }()

	//“ping”消息通过通道成功地从一个goroutine传递到另一个goroutine。
	msg := <-messages
	fmt.Println(msg)*/

	/*//2.默认情况下，通道是未缓冲的.这里使一个字符串的通道缓冲多达2个值
	messages := make(chan string, 2)


	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)*/

	/*//3.通道在goroutine上同步执行程序
	done := make(chan bool, 1)
	go worker(done)
	//如果从此程序中删除 <-done 行，程序将在工作程序(worker)启动之前退出。
	<-done*/

	//4.当使用通道作为函数参数时，可以指定通道是否仅用于发送或接收值。

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

//ping功能只接受用于发送值的通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
