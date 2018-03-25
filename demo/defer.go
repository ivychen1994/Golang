package main
import "fmt"

func main() {
	/*
	//defer 语句会延迟函数的执行直到上层函数返回。
	defer fmt.Println("world")
	fmt.Println("hello")*/

	//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}



