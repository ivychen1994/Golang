package main
import "fmt"
import "time"



//使用goroutines和channel实现一个工作池。
func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}

//这些工作程序(worker)将在工作渠道上接收工作，并将结果发送相应的结果。
// 将每个工作程序(worker)睡一秒钟，用来模拟执行的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

