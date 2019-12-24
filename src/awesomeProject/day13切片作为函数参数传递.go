package main

import "fmt"

//切片作为函数的参数传递  是地址传递
//切片名就是个地址
//形参的改变也是实参的改变
/*
注意:append使用修改的时候 形参地址会发生地址的改变
	不会指向实参的 所以实参就不会改变
*/

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	s1 := demo(s)
	fmt.Println(s)
	fmt.Println(s1)
}

func demo(s []int) []int {
	//fmt.Println(s)
	//s[0] = 123
	//fmt.Println(s)
	s = append(s, 7, 8)
	return s
}
