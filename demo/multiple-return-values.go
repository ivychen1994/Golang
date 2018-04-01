package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
//使用从多个赋值的调用返回2个不同的值。如果只想要返回值的一个子集，请使用空白标识符_
	_, c := vals()
	fmt.Println(c)
}

