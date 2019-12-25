package main

import "fmt"

//指针作为函数参数  地址传递
func main() {
	a := 10
	b := 20
	swap1(&a, &b)
	fmt.Println(a, b)
}

//地址传递
func swap1(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
}
