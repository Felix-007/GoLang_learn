package main

import "fmt"

/**
面向对象
匿名字段 来实现继承
*/

//结构体嵌套  类的继承
type PPerson struct {
	id   int
	name string
	age  int
}

type SStudent struct {
	PPerson //匿名字段
	score   int
}

func main() {
	//学生对象的创建
	stu := SStudent{PPerson{101, "小红", 13}, 23}
	fmt.Println(stu)
	stu.score = 123
	fmt.Println(stu)
	s1 := SStudent{PPerson{name: "小明", id: 12}, 12}
	fmt.Println(s1)

	//var s2=new(SStudent)
	var s3 SStudent
	fmt.Println(s3)

}
