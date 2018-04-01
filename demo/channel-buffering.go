package main

import "fmt"

func main() {
//默认情况下，通道是未缓冲的，意味着如果有相应的接收(<- chan)准备好接收发送的值，它们将只接受发送(chan <- )。
// 缓冲通道接受有限数量的值，而没有用于这些值的相应接收器。
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

