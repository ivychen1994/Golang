package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//1.for 循环的 range 格式可以对 slice 或者 map 进行迭代循环
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//2.可以通过赋值给 _ 来忽略序号和值
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	//3.示例
	/* create a slice *//*
	numbers := []int{0,1,2,3,4,5,6,7,8}

	*//* print the numbers *//*
	for i:= range numbers {
		fmt.Println("Slice item",i,"is",numbers[i])
	}

	*//* create a map*//*
	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}

	*//* print map using keys*//*
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	*//* print map using key-value*//*
	for country,capital := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",capital)
	}
*/
	//4.使用此语法对从通道接收的值进行迭代。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
