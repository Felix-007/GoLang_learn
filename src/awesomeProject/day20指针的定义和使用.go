package main

import "fmt"

/**
指针的定义
作用:直接改变值比较方便
	数据比较大的变量 可以直接传地址修改 不用传大的
p:= &str	&变量  取地址操作 引用运算符
*p=57 //对该地址的值进行操作    *指针变量 取值操作 解引用运算符
*/
func main() {
	var i int
	i = 100
	fmt.Printf("%d\n", i)
	fmt.Printf("%p\n", &i)
	//fmt.Printf("%v\n",&i)  //扩展

	//指针定义 指针默认值  <nil>
	var p *int = nil //可以设置默认值
	fmt.Println(p)
	p = &i
	fmt.Println(p)
	fmt.Printf("%T", p)

	str := "abc"
	var p1 *string
	p1 = &str
	fmt.Println(p1)

	//自动推导
	p2 := &str
	fmt.Println(p2)
	//* 通过地址修改值
	*p2 = "ZXV"
	fmt.Printf("p2=%s,p2=%p\n", str, p2)
	fmt.Printf("p2=%s,p2=%v\n", str, *p2)

	//野指针  指针指向一个未知地址
	//var p3 * int
	//*p3=56 	//没有指向关系
	//fmt.Println(p3)

	var a int
	p4 := &a
	*p4 = 56
	fmt.Println(p4, *p4)

	//new函数
	var p5 *int
	//p5:=new(int)
	p5 = new(int) //开辟一个空间 所以指针p5不是野指针
	*p5 = 157
	fmt.Println(*p5)

}
