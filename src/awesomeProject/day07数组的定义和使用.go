package main

import "fmt"

func main01() {
	//数组的定义
	//var 数组名 [长度] 类型
	var a [10]int
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}
	fmt.Println(a)

	for jj, ggg := range a {
		fmt.Printf("i%d=%d\t", jj, ggg)
	}

}

func main02() {
	//数组的初始化 长度 类型 {元素}
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	//自动推导
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	//部分初始化
	c := [5]int{1, 2, 3}
	fmt.Println(c)
	//指定位置初始化  2:10 下标为2的赋值10
	d := [5]int{2: 10, 4: 20}
	fmt.Println(d)
	//自动推导数组长度  ...
	f := [...]int{1, 2, 3, 4}
	fmt.Println(len(f))
}

func main() {
	//常见问题
	//长度常量
	//数组下标越界 -1不能用
	a := [...]int{1, 23, 5}
	fmt.Println(len(a))
	//数组直接的赋值 赋值
	b := a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

}
