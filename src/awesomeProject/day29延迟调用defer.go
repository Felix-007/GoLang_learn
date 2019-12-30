package main

import "fmt"

func main() {

	//defer 只能出现在函数内部
	//defer fmt.Println("hello")
	//fmt.Println("老王")

	//一般应用在 文件关闭

	////多个defer 采用后进先出的顺序执行
	//defer fmt.Println("hello2")
	//defer fmt.Println("老王")
	//defer fmt.Println("你好")

	////即使函数或者某个延时调用发送错误 这些调用也就会被执行
	//defer fmt.Println("hello3")
	//defer fmt.Println("老王3")
	//test29(0) //最后出现异常
	//defer fmt.Println("你好3")

	//defer匿名函数使用
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("匿名函数1", a)
		fmt.Println("匿名函数1", b)
	}(a, b)
	defer func() {
		fmt.Println("匿名函数2", a)
		fmt.Println("匿名函数2", b)
	}()
	a = 100
	b = 200
	fmt.Println("main函数", a)
	fmt.Println("main函数", b)
}

func test29(x int) {
	v := 100 / x
	fmt.Println(v)
}
