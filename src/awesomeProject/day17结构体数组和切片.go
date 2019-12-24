package main

import "fmt"

//结构体数组  值传递
//结构体切片  地址传递
type Student struct {
	id   int
	name string
	age  int
	sex  string
}

func main() {
	//结构体数组  值传递  不影响修改
	arry1 := [3]Student{
		{id: 1, name: "小大", age: 2},
		{id: 2, name: "小中", age: 4},
		{id: 3, name: "小小", age: 3}}
	fmt.Println(arry1)
	fmt.Printf("%p\n", &arry1)

	for v, k := range arry1 {
		fmt.Println(v, "=", k)
	}

	//结构体切片  地址传递 影响原值
	s1 := []Student{{id: 4, name: "小大", age: 5},
		{id: 5, name: "小中", age: 10},
		{id: 6, name: "小小", age: 8}}
	s1 = append(s1, Student{id: 8, name: "id8", age: 128})
	fmt.Println(s1)
	fmt.Println("切片排序")
	sort2(s1)
	fmt.Println("排序后", s1)
}

//结构体数组数据 作为函数参数
func sort1(a [3]Student) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j].age > a[j+1].age {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("数组a=", a)
}

//结构体切片数据 作为函数参数
func sort2(a []Student) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j].age > a[j+1].age {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	a = append(a, Student{id: 8, name: "id8", age: 999})
	fmt.Println("切片a=", a)
}
