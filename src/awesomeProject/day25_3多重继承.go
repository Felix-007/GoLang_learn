package main

import "fmt"

type person4 struct {
	name string
	age  int
	sex  string
}
type Person4 struct {
	id   int
	addr string
}

type Student4 struct {
	Person4 //匿名字段
	person4
	score int
}

func main() {
	var stu Student4
	stu.id = 101
	stu.addr = "火星"
	stu.name = "小火"
	stu.age = 64
	stu.sex = "男"
	stu.score = 95
	fmt.Println(stu)

	//自动推导
	s := Student4{score: 12, person4: person4{sex: "男"}, Person4: Person4{id: 12}}
	fmt.Println(s)
}
