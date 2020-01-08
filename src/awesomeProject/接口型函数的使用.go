package main

import "fmt"

type Handler interface {
	Do(v interface{})
}

type welcome string

func (w welcome) Do(v interface{}) {
	fmt.Printf("我叫%s\n", v)
}

//接口型函数
type FuncIn func(interface{})

func (f FuncIn) Do(v interface{}) {
	f(v)
}

func selfInfo(v interface{}) {
	fmt.Printf("我叫%s\n", v)
}

func Each(m []interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for _, v := range m {
			h.Do(v)
		}
	}
}

//去掉强转
func EachFunc(m []interface{}, f func(interface{})) {
	Each(m, FuncIn(f))
}

func main() {
	persons := make([]interface{}, 0)
	persons = append(persons, "A", "B", "C")

	//var w welcome = "大家好"

	//Each(persons, w)

	/*
		1.	因为必须要实现Handler接口，Do这个方法名不能修改，不能定义一个更有意义的名字
		2.	必须要新定义一个类型，才可以实现Handler接口，才能使用Each函数
		所以解决以上问题 可以使用接口型函数
	*/
	Each(persons, FuncIn(selfInfo))

	//var vv Handler
	//vv = FuncIn(func(v interface{}) {
	//	fmt.Println("!@3")
	//})
	//vv.Do(123)
	//fmt.Printf("%T",vv)

	//再次包装 去掉强装
	EachFunc(persons, selfInfo)
}
