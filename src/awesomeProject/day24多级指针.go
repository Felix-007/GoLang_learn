package main

import "fmt"

func main() {
	a := 10
	p := &a
	b := 200

	*p = 12
	fmt.Println("a1=", a)
	fmt.Println(p)
	fmt.Printf("一级指针:%T\n", p)

	//  ***ppp=**pp=*p=a

	pp := &p
	*(*pp) = 24
	fmt.Println("a2=", a)
	fmt.Println(pp)
	fmt.Printf("二级指针:%T\n", pp)

	ppp := &pp
	*(*(*ppp)) = 48
	fmt.Println("a3=", a)
	fmt.Println(ppp)
	fmt.Printf("三级指针:%T\n", ppp)

	pppp := &ppp
	****pppp = 96
	fmt.Println("a4=", a)
	fmt.Println(pppp)
	fmt.Printf("四级指针:%T\n", pppp)

	//通过二级指针 连接修改一级指针的值
	*pp = &b
	fmt.Println("改变后的一级指针=", *p)
	//ppp=&pp
	fmt.Println("改变后的一级指针=", **pp)

}
