package main

import "fmt"

//支持匿名函数，可以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))
}

//递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

