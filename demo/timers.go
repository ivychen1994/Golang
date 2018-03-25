package main
import "time"
import "fmt"

func main() {

	//计时器将等待2秒钟。
	//<-timer1.C阻塞定时器的通道C，直到它发送一个指示定时器超时的值。
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//只是想等待，可以使用time.Sleep
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
