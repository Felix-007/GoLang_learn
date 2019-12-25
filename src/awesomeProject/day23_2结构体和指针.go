package main

import "fmt"

/**
结构体指针 *person
结构体指针作为函数参数
*/
func main() {
	h := man{101, "小明", 34}
	fmt.Println(h)
	fmt.Printf("%p\n", &h)
	p := &h
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)
	var p1 *man = &h
	fmt.Printf("%p\n", p1)

	//指针修改
	(*p).age = 45
	(*p1).name = "大鱼"
	//直接操作
	p.id = 201
	fmt.Println(h)

	changeM(p)
	fmt.Println(h)
}

func changeM(m *man) {
	m.age = 99
}

type man struct {
	id   int
	name string
	age  int
}
