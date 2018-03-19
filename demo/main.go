//当前程序的包名
package main

import (
	"fmt"
	//"reflect"
	//"sort"
	//"time"
	"runtime"
)

//常量的定义
const PI = 3.14

//全局变量的声明与赋值
var name = "gopher"

//结构的声明
type gopher struct {
}

//j接口的声明
type golang interface {
}

type (
	byte int8
	文本 string
)

var (
	aaa      = "hello"
	sss, bbb = 1, 2
)

//向mian函数作为程序入口电启动
func main() {
	//var b int=1
	//var b 文本
	//b="文本啊啊啊"
	/*
	b:=1//变量声明与赋值的最简写法
	var e,f,g,h int=5,6,7,8//多个变量声明同时赋值
	var i,j,k,l=9,10,11,12//省略变量类型，由系统推断
	fmt.Println(b+e+f+g+h+i+j+k+l)
	*/
	/*
	//在相互兼容的两种类型之间转换
	var a float32=1.1
	aa :=int(a)
	fmt.Println(aa)
	var a int=65
	aa:=string(a)
	fmt.Println(aa)
	*/

	/*
		//指针
		c:=1
		c++
		var p *int=&c
		fmt.Println(*p)
	*/

	/*
	a:=10
	if a:=1;a>0{
		fmt.Println(a)
	}*/

	//for使用
	//a:=1
	/*for{
		a++
		if a>3 {
			break
		}
		fmt.Println(a)
	}
	for a<=3{
		a++
		fmt.Println(a)
	}
	fmt.Println("Over")
	*/

	/*
	//switch语句
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("none")
	}
	*/

	//数组
	//var a[2]string
	//a:=[20]int{19:1}
	//a:=[...]int{1,2,3,4,5}
	//a := [...]int{99: 1}
	//var p *[100]int = &a
	//fmt.Println(*p)
	//x, y := 1, 2
	//b := [...]*int{&x, &y}
	//fmt.Println(b)
	/*	a:=[2]int{1,2}
		b:=[2]int{1,3}
		fmt.Println(a==b)
		//指针
		p:=new([10]int)
		p[1]=2
		fmt.Println(p)
		//多维数组
		q:=[2][3]int{
			{1,1,1},
			{2,2,2}}
		fmt.Println(q)
		c:=[...]int{5,2,6,3,9}
		num:=len(c)
		for i:=0;i<num ;i++  {
			for j:=i+1;j<num ;j++  {
				if c[i]<c[j] {
					temp:=c[i]
					c[i]=c[j]
					c[j]=temp
				}
			}
		}
		fmt.Println(c)*/

	//切片Slice
	//	var sl []int
	//fmt.Println(sl)
	//a:=[10]int{1,2,3,4,5,6,7,8,9}
	//fmt.Println(a)
	//sl:=a[5:]//取出后五个元素
	//s2:=a[:5]//取出前五个元素
	//fmt.Println(s2)

	//slice
	/*s1:=make([]int,3,10)
	fmt.Println(len(s1),cap(s1))//长度，容量

	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a[3:5]
	fmt.Println(len(sa),cap(sa))
	fmt.Println(sa)//类型转换*/

	//apend
	/*s1:=make([]int,3,6)
	fmt.Print("%p\n",s1)
	s1=append(s1,1,2,3)
	fmt.Print("%v %p\n",s1,s1)
	s1=append(s1,1,2,3)
	fmt.Print("%v %p\n",s1,s1)*/

	/*a:=[]int{1,2,3,4,5}
	s1:=a[2:5]
	s2:=a[1:3]
	fmt.Println(s1,s2)
	s2=append(s2,1,2,1,1,1,1,1,1,1,1)//容量超过，重新分配
	s1[0]=9
	fmt.Println(s1,s2)//发现s1中的3变成了9，s2中的3也成了9*/

	//Copy
	/*s1:=[]int{1,2,3,4,5,6}
	s2:=[]int{7,8,9,10,1,1,1,1,1,}
	copy(s2[2:4],s1[2:3])
	fmt.Println(s2)*/

	//Map
	/*var m map[int]string
	m = make(map[int]string)
	//m:=make(map[int]string)
	m[1] = "ok"
	//delete(m,1)
	//a:=m[1]
	fmt.Println(m)*/

	/*var m map[int]map[int]string
	m = make(map[int]map[int]string)
	//赋值操作
	m[1]=make(map[int]string)
	m[1][1]="ok"
	a:=m[1][1]
	a, ok := m[2][1]
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	a, ok = m[2][1]
	a, ok = m[2][1]
	fmt.Println(m, a,ok)*/

	//迭代操作
	/*sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)*/

	/*m:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	s:=make([]int,len(m))
	i:=0
	for k,_:=range m {
		s[i]=k
		i++
	}
	sort.Ints(s)
	fmt.Println(m)*/

	//将Map中的键值和value互换
	/*m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)*/

	//fmt.Println(math.MaxInt32)

	//当传递类型为Int,string时，只是值拷贝
	/*a := 1
	B(a)
	fmt.Println(a)*/

	//使用指针,拷贝地址
	/*a := 1
	B(&a)
	fmt.Println(a)*/

	//对slice的内存地址进行了拷贝
	/*s := []int{1, 2, 3, 4}
	C(s)
	fmt.Println(s)*/

	//匿名函数
	/*a := func() {
		fmt.Println("Func A")
	}
	a()*/

	//闭包中使用匿名函数
	/*f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))*/

	//defer
	/*fmt.Println("a")
	defer fmt.Println("b")
	defer  fmt.Println("c")*/

	/*for i:=0;i<3 ;i++  {
		defer func() {
			fmt.Println(i)
		}()
	}*/

	//A()
	//B()
	//C()

	/*a:=&person{
		Name:"joe",
		Age:19,
	}
	//a.Name="joe"
	//a.Age=19
	a.Name="ok"//不需要取地址
	fmt.Println(a)
	A(a)
	fmt.Println(a)*/
	/*a:=&struct {
		Name string
		Age int
	}{
		Name:"joe",
		Age:19,
	}*/
	//a:=person{human:human{Sex:9},Name:"joe",Age:19,}
	//a.Sex=13
	//fmt.Println(a)

	/*a:=A{}
	a.Print()
	b:=B{}
	b.Print()
*/
	/*var a USB
	a = PhoneConnector{"PhoneConnexter"}
	a.Connect()
	Disconnect(a)

	pc := PhoneConnector{"PhoneConnexter"}
	var c Connecter
	c = Connecter(pc)
	c.Connect()*/

	/*var b interface{}
	fmt.Println(b == nil)
	var p *int = nil
	b = p
	fmt.Println(b == nil)*/

	/*m:=Manager{User:User{1,"OK",12},title:"123"}
	t:=reflect.TypeOf(m)
	//fmt.Printf("%#v\n",t.Field(0))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0,1}))*/

	/*
	u := User{1, "ok",12}
	Info(u)*/

	/*x:=123
	v:=reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)*/

	/*u:=User{1,"ok",12}
	Set(&u)
	fmt.Println(u)*/

	/*u:=User{1,"OK",12}
	v:=reflect.ValueOf(u)
	mv:=v.MethodByName("Hello")

	args:=[]reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
*/

	//go Go()
	//time.Sleep(2 * time.Second) //两秒钟
	//
	//c := make(chan bool) //创建一个channel
	//go func() {
	//	fmt.Println("Go Go Go!!!")
	//	c <- true
	//	close(c)
	//}()
	////<-c
	//for v:=range c{
	//	fmt.Println(v)
	//}
	runtime.GOMAXPROCS(runtime.NumCPU())
	c:=make(chan bool,10)
	for i:=0;i<10 ;i++  {
		go Go(c,i)
	}
	for i:=0;i<10 ;i++  {
		<-c
	}


}

/*
func A()  {
	fmt.Println("FUNC A")
}

func B()  {
	defer func() {
		if err :=recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")//阻断C 的运行
}

func C()  {
	fmt.Println("Func C")
}
*/

//闭包中使用匿名函数
/*func closure(x int) func(int) int {
	fmt.Println(&x)
	return func(y int) int {
		fmt.Println(&x)
		return x + y
	}
}*/

/*func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}*/

//只是进行了值拷贝
/*func B(a int) {
	a = 2
	fmt.Println(a)
}*/

//使用指针,拷贝地址
/*func B(a *int) {
	*a = 2
	fmt.Println(*a)
}*/

//对slice的内存地址进行了拷贝
/*func C(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	fmt.Println(s)
}*/

//结构体
type person struct {
	human //嵌套
	Name string
	Age  int
}
type human struct {
	Sex int
}

/*func A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}*/
//method
/*type A struct {
	Name string
}
type B struct {
	Name string
}
func (a A)Print()  {
	fmt.Println("A");
}
func (b B)Print()  {
	fmt.Println("B");
}*/

//INTERFACE
/*type USB interface{
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnector struct{
	name string
}

func (pc PhoneConnector)Name() string{
	return pc.name
}

func (pc PhoneConnector)Connect() {
	fmt.Println("Connect:",pc.name)
}

func Disconnect(usb USB)  {
	if pc,ok:=usb.(PhoneConnector);ok{
		fmt.Println("Disconnected.",pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}*/

/*
func Disconnect(usb interface{})  {
	switch v:=usb.(type){
	case PhoneConnector:fmt.Println("Disconnected.",v.name)
	default:fmt.Println("Unknown decive.")
	}
}*/

//reflection
/*
type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

*/
/*func (u User) Hello() {
	fmt.Println("Hello World");
}*//*


func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%vin", f.Name, f.Type, val)
	}

	for i:=0;i<t.NumMethod();i++{
		m:=t.Method(i)
		fmt.Printf("%6s:%v\n",m.Name,m.Type)
	}
}

func Set(o interface{})  {
	v:=reflect.ValueOf(o)
	if v.Kind()==reflect.Ptr&&!v.Elem().CanSet(){
		fmt.Println("XXX")
		return
	}else {
		v=v.Elem()
	}
	f:=v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("BAD")
		return
	}
	if f.Kind()==reflect.String{
		f.SetString("Byebye")
	}
}

//通过反射调用方法
func (u User)Hello(name string){
	fmt.Println("Hello",name,"","my name is",u.Name)
}
*/

//concurrent并发
func Go(c chan bool,index int ) {
	a:=1
	for i:=0;i<10000000 ;i++  {
		a+=i
	}
	fmt.Println(index,a)
	/*if index==9{
		c<- true
	}*/
	c<- true
}
