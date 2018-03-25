package main

import "fmt"

//关闭通道表示不会再发送更多值
// 它反复从j的工作接收more := <-jobs。在这种特殊的2值形式的接收中，
// 如果作业已关闭并且已经接收到通道中的所有值，
// 则 more 的值将为 false。当已经完成了所有的工作时，使用这个通知。
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

