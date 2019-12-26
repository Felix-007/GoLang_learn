package main

import "fmt"

type person6 struct {
	id   int
	name string
	age  int
}

type student6 struct {
	person6
	class int
}

func (p *person6) printInfo() {
	fmt.Printf("编号id=%d\n", p.id)
	fmt.Printf("姓名=%s\n", p.name)
	fmt.Printf("年纪=%d\n", p.age)
}

//重写父类 printInfo 方法
func (s *student6) printInfo() {
	fmt.Printf("学生班级:=%d\n", s.class)
}

func main() {
	p := person6{110, "亚索", 33}
	p.printInfo()

	//子类重写方法
	s := student6{person6{120, "岩雀", 23}, 9}
	s.printInfo()
	//父类的调用
	s.person6.printInfo()
}
