package main

import "fmt"

type Opter interface {
	Result() int
}

type Operate struct {
	num1 int
	num2 int
}

type add struct {
	Operate
}
type sub struct {
	Operate
}

func (a *add) Result() int {
	return a.num1 + a.num2
}

func (s *sub) Result() int {
	return s.num1 - s.num2
}

func main() {
	add := add{Operate{10, 20}}
	sub := sub{Operate{10, 20}}

	//var op Opter
	//op = &sub
	//fmt.Println(op.Result())

	//多态实现
	//接口作为函数参数
	Result(&add)
	Result(&sub)

}

func Result(opter Opter) {
	fmt.Println(opter.Result())
}
