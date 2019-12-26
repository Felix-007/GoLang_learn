package main

import "fmt"

type PPerson2 struct {
	id   int
	name string
	age  int
}

type SStudent2 struct {
	PPerson2 //匿名字段
	score    int
	name     string
}

type Person3 struct {
	id   int
	name string
	age  int
}

type Student3 struct {
	*Person3 //指针类型的匿名字段
	score    int
	name     string
}

func main() {
	var s1 SStudent2
	s1.name = "花花"
	s1.PPerson2.name = "huahua"
	fmt.Println(s1)

	var s2 Student3
	s2.name = "小明"
	s2.score = 68

	fmt.Println(s2) //{<nil> 68 小明}
	s2.Person3 = new(Person3)
	s2.Person3.name = "xiaoming"
	s2.score = 123
	fmt.Println(s2.name)
	fmt.Println(*s2.Person3)

}
