package main

import "fmt"

func main() {
	//数组的长度是固定的..
	var arr [5]int = [5]int{1, 2, 3, 4}
	fmt.Println(arr)

	//切片 可以看做一个 复合数据类型
	//切片长度不固定可增加
	//var slice []int
	var slice []int = []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 7) //新增元素
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) //cap查看切片的容量
	//自动推导创建
	slice2 := []int{2, 34, 5}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	slice2 = append(slice2, 1)
	fmt.Println(cap(slice2))

	//make函数创建   长度 容量(可以省略)
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println("赋值操作")

	//重点 切片名就是地址 append的收 切片地址可能会改变
	//切片append的时候长度超过容量 容量自动扩容
	//一般 容量*2 如果超过1024 每次扩容 上次1/4
	//
	fmt.Println("切片的地址和扩容")
	fmt.Printf("%p\n", slice2)
	slice2 = append(slice2, 123)
	slice2 = append(slice2, 54)
	fmt.Printf("%p\n", slice2)

}
