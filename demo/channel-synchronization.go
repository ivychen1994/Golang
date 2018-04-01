package main
import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

//可以使用通道在goroutine上同步执行程序。
// 这里有一个使用阻塞接收等待goroutine完成的示例。
func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done

	//使用通道作为函数参数时，可以指定通道是否仅用于发送或接收值
	//ping函数接受一个通道接收(ping)，一个接收发送(ping)。
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}


