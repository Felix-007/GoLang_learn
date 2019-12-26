package main

import "fmt"

type person7 struct {
	id   int
	name string
	age  int
}

type student7 struct {
	person7
	class int
}

func (p person7) PrintInfo1() {
	fmt.Printf("Info1%p,%v\n", &p, p)
}
func (p *person7) PrintInfo2() {
	fmt.Printf("Info2%p,%v\n", p, p)
}

func main() {
	p := person7{101, "mike", 22}
	p.PrintInfo1() //0xc0000044c0,{101 mike 22}
	p.PrintInfo2() //0xc000006030,&{101 mike 22}

	fmt.Println(p.PrintInfo1)
	fmt.Println(p.PrintInfo2)
	fmt.Printf("%T\n", p.PrintInfo1)
	fmt.Printf("%T\n", p.PrintInfo2)

	//方法值		隐式传递 隐藏的是p接收者  绑定对象
	//var pfunc =func()	//该方式用的少
	pfunc1 := p.PrintInfo1
	//pfunc1=p.PrintInfo2	//都是func()类型 可以相互赋值
	pfunc1()                   //p.PrintInfo1()
	fmt.Printf("%T\n", pfunc1) //func()

	//方法表达式	需要进行显示传参
	fmt.Println("方法表达式")
	pfunc11 := person7.PrintInfo1
	pfunc11(p)
	pfunc22 := (*person7).PrintInfo2
	//pfunc11=(*person7).PrintInfo2=pfunc22()	//类型不用 不能相互赋值
	pfunc22(&p)
	fmt.Printf("%T\n", pfunc11)
	fmt.Printf("%T\n", pfunc22)

}
