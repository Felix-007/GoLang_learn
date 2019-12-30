package main

import "fmt"

//设置 recover()
//在defer调用的匿名函数中使用recover()
/**
defer func(){
	recover()
}()
*/
func main() {
	test30_1()
	test30_2(3) //异常中断 panic

	test30_3()

}

func test30_1() {
	fmt.Println("test30_1")
}

func test30_2(x int) {
	//fmt.Println("test30_1")
	defer func() {
		//recover()	nil
		//使用判断
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [3]int
	a[x] = 999

}

func test30_3() {
	fmt.Println("test30_3")
}
