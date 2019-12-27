package main

import "fmt"

type Persioner interface {
	Humaner9
	sing(string)
}
type Humaner9 interface {
	sayhi()
}

type student9 struct {
	name  string
	age   int
	score int
}

type teacher9 struct {
	name    string
	age     int
	subject string
}

func (s *student9) sayhi() {
	fmt.Println("hi,我是", s.name, "今年", s.age, "岁,成绩:", s.score, "分!")
}

func (s *student9) sing(name string) {
	fmt.Println("我是", s.name, "唱"+name)
}
func (t *teacher9) sayhi() {
	fmt.Println("hi,我是", t.name, "今年", t.age, "岁,主讲科目:", t.subject, "!")

}
func main27() {
	stu := student9{"小米", 23, 34}
	var h Humaner9
	h = &stu
	h.sayhi()

	var p Persioner
	p = &stu
	p.sayhi()
	p.sing("ABC")
}

func main() {
	var h Humaner9
	var p Persioner
	stu := student9{"小米", 23, 34}

	//h=&stu
	//p=h	//err

	p = &stu
	h = p
	h.sayhi()
}
