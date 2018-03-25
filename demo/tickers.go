package main
import "time"
import "fmt"

func main() {
	//代码机使用与计时器的机制类似：发送值到通道。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	//当代码停止后，它不会在其通道上接收任何更多的值。我们将在1600ms后停止。
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
