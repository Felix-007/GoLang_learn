package main

import (
	"fmt"
	"unsafe"
)

/*
常量
常量iota
基础数据类型bool int float
===一个汉字占三个字符
复数类型
*/
func main() {
	//var a,b int
	//a=1
	//b=2
	//fmt.Println(a,b)
	//
	//c,d:=10,20
	//fmt.Println(c,d)
	//
	//var(
	//	e int=1
	//	f float64=4.12
	//)
	//fmt.Println(e,f)
	//
	//var(
	//	g=12
	//	h=32.1
	//)
	//fmt.Println("###",g,h)
	//
	//u,q:=123,"asd"
	//fmt.Println(u,q)
	//
	//const (
	//	i=14
	//	j=25
	//)
	//fmt.Println(i,j)
	//
	//fmt.Printf("%T",i)

	//iota 枚举
	//1.iota 常量自动生成器,每行自动+1
	//2.iota 遇到const,重置为0

	const (
		xx = 1
		a  = iota
		b  = iota
		c  = iota
	)

	fmt.Println(a, b, c) //0,1,2

	const d = iota
	fmt.Println(d)
	//3.iota 同一个()内声明 可以只写一个iota
	const (
		a1 = iota
		b1
		c1
		d1 = 100
		e1
		f1 = iota
	)
	fmt.Println(a1, b1, c1, d1, e1, f1)
	//4.如果iota在同一行 值是一样的
	const (
		z    = iota       //0
		x, v = iota, iota //1,1
		k    = iota       //2
	)
	fmt.Println(z)
	fmt.Println(x, v)
	fmt.Println(k)

	qq := 2 == 2
	fmt.Println(qq)

	var ccc int32 = 10
	fmt.Println(ccc)
	fmt.Println(unsafe.Sizeof(ccc))

	fmt.Println('\t')

	//var l byte='A'
	l := 'A'
	fmt.Println(l)
	fmt.Printf("%T\n", l)
	h := 'h'
	fmt.Printf("%c\n", h)
	str := "abc"
	fmt.Println(str)
	fmt.Printf("%T\n", str)
	fmt.Println(len(str))
	fmt.Println(len("真TM的帅"))
	fmt.Println(str[1])
	fmt.Printf("%c\n", str[1])
	str3 := "真的很牛逼A"
	fmt.Printf("%c\n", str3[5])

	//复数  实部+虚部
	var cc1 complex128
	cc1 = 2.1 + 3.14i
	fmt.Println(cc1)

	cc2 := 2.2 + 3.11i
	fmt.Println(cc2)
	fmt.Printf("%T\n", cc2)
	//分别取实部 虚部
	fmt.Println(real(cc1), imag(cc1))
}
