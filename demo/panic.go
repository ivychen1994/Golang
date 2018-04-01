package main

import (
	"os"
	"fmt"
)
//在程序中可使用panic来检查意外错误。
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	//创建一个文件，写入内容，然后在完成之后关闭
	f := createFile("defer-test.txt")
	defer closeFile(f)
	writeFile(f)

}


func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
