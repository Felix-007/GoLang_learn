package main

import "fmt"

/**
结构体指针 *person
结构体指针作为函数参数
*/
func main23_2() {
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

//结构体数组指针
func main_23_4() {
	arr := [3]man{{101, "钢铁侠", 34},
		{102, "绿巨人", 45},
		{103, "队长", 193}}

	var p *[3]man
	p = &arr

	p[0].age = 56

	fmt.Println(p)
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

	for i, v := range p {
		fmt.Println(i, "=", v)
	}

}

//  数组指针map
func main() {

	m := make(map[int]*[3]man)
	p := &m
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	m[101] = new([3]man)
	fmt.Println(m[101])
	m[101] = &[3]man{{101, "钢铁侠", 34},
		{102, "绿巨人", 45},
		{103, "队长", 193}}
	m[102] = &[3]man{{104, "蝙蝠侠", 34},
		{105, "闪电侠", 45},
		{106, "海王", 193}}

	//指针操作 修改年纪
	m[101][0].age = 101

	for key, v := range m {
		fmt.Println(key, v)
	}

}
