package main

import "fmt"

/**
异常的处理 error
panic 函数终止程序的错误
*/

func test28(a, b int) (v int, err error) {
	//var v int
	err = nil
	if b == 0 {
		fmt.Println("err=", err)
	} else {
		v = a / b
	}
	return
}
func main() {
	//fmt.Println(test28(4, 2))
	result, err := test28(10, 0)
	if err == nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("hello")
	//遇到panic自动终止
	//panic("hello")
	fmt.Println("hello")

	//panic(index out of range [4] with length 3)
	test28_2(4)

}

func test28_2(i int) {
	var arr [3]int
	arr[i] = 999
	fmt.Println(arr)
}
