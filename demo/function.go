package main
import "fmt"

//函数intSeq()返回另一个函数，它在intSeq()函数的主体中匿名定义。返回的函数闭合变量i以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}


func main() {

	//当调用intSeq()函数，将结果(一个函数)分配给nextInt。这个函数捕获它自己的i值，每当调用nextInt时，它的i值将被更新。
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}



