package main

import "fmt"

//函数有一个int参数，因此参数将通过值传递给它。
func zeroval(ival int) {
	ival = 0
}

//函数体中的* iptr代码将指针从存储器地址解引用到该地址处的当前值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

