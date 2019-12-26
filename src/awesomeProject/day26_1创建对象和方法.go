package main

import "fmt"

/**
对象的方法  func (对象)方法 (参数列表) 返回值列表{
				代码体
			}
*/
type cat struct {
	name string
	age  int
}

func (c cat) show() {
	fmt.Println("喵喵喵", c.name)
}

type dog struct {
	name string
	age  int
}

//  地址传递
// 方法调用中 方法的接收者 是指针类型
// 指针对象类型 和 普通类型表示的是相同的对象
func (d *dog) show() {
	fmt.Println("旺旺旺")
	d.name = "嘿嘿"
	fmt.Println(d) //&{嘿嘿 3}
}

func main() {
	c := cat{"小花", 2}
	fmt.Println(c)
	c.show()

	d := dog{"小黑", 3}
	fmt.Println(d)
	d.show()
	fmt.Println(d)

	//e := cat{"m名字", 2}
	//fmt.Println(e)
	//e.show()

}
