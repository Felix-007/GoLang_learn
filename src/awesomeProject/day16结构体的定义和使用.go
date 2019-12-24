package main

import "fmt"

//结构体的定义
//	   一般大写开头 别的文件引用
//type 结构体名 struct{
//
//}

type Person struct {
	id   int
	name string
	age  int
	sex  string
	addr string
}

func main() {
	//1.顺序初始话
	var p2 Person = Person{1, "小米", 12, "男", "郊区"}
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	//2.自动推导
	p3 := Person{2, "小明", 12, "女", "城区"}
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	p1 := Person{id: 1, name: "小米", sex: "女", age: 18}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	//3.
	var p4 Person
	p4.name = "小南"
	p4.age = 12
	p4.id = 5
	fmt.Println(p4)
	fmt.Printf("%T\n", p4)

	p6 := Person{id: 82, name: "大夕", age: 67}
	fmt.Println(p6)
	fmt.Printf("%p\n", &p6)
	fmt.Printf("%p\n", &p6.id)
	//结构体的比较
	p7 := Person{id: 82, name: "大夕", age: 67}
	p8 := Person{id: 82, name: "大夕", age: 67}
	p9 := Person{id: 82, name: "大1", age: 67}
	fmt.Println(p7 == p8)
	fmt.Println(p8 != p9)
	//结构体的赋值
	p8 = p9
	fmt.Println(p8 == p9)

}
