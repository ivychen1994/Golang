package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//使用互斥体来安全地访问多个goroutine中的数据。\
//使用显式锁定互斥体来同步对多个goroutine的共享状态的访问。
func main() {

	//状态(state)是一个映射。
	var state = make(map[int]int)

	//互斥将同步访问状态。
	var mutex = &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0

	//启动100个goroutine来对状态执行重复读取，每个goroutine中每毫秒读取一次。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				//Lock()互斥体以确保对状态的独占访问，读取所选键的值，
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				//Unlock()互斥体，并增加readOps计数。
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}


	time.Sleep(time.Second)


	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
