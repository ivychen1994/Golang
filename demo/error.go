package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64)(float64, error) {
	if(value < 0){
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	//函数通常返回错误作为最后一个返回值。 可使用errors.New来构造一个基本的错误消息
	result, err:= Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	//2errors.New使用给定的错误消息构造基本错误值。错误位置中的nil值表示没有错误。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}


func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")

	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

