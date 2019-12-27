package main

import "fmt"

//定义接口
//一般 er 结尾  根据接口实现功能
type Humaner interface {
	//方法
	sayhi()
	//接口的方法定义一定要实现
	//result()
}

type student8 struct {
	name  string
	age   int
	score int
}

type teacher8 struct {
	name    string
	age     int
	subject string
}

func (s *student8) sayhi() {
	fmt.Println("hi,我是", s.name, "今年", s.age, "岁,成绩:", s.score, "分!")
}

func (t *teacher8) sayhi() {
	fmt.Println("hi,我是", t.name, "今年", t.age, "岁,主讲科目:", t.subject, "!")

}

func main() {
	//接口是一种数据类型 可以接受满足对象的信息
	//接口是虚的 方法是实的
	var h Humaner

	s := student8{"小皮", 23, 89}
	t := teacher8{"大佬", 43, "科学"}
	s.sayhi()
	t.sayhi()

	h = &s
	h.sayhi()
	h = &t
	h.sayhi()

}
