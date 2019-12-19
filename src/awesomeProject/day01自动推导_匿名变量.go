package main

import "fmt"

/*
自动推导
_匿名变量
*/
func main() {
	fmt.Println("hell Fuck")
	var size int
	fmt.Println(size)

	/*
		自动推导类型
		:=左边必须是没使用过的变量名

	*/
	var a = 1
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := 20.3
	fmt.Println(b)

	fmt.Printf("%T\n", b)

	c := 3

	a, c = c, a
	fmt.Println(a, c)

	//println和printf的区别

	//_匿名变量 丢弃数据不处理
	tmp, _ := 10, 20
	fmt.Println(tmp)

	_, e, f := test1()
	fmt.Println(e, f)

}

//可以返回多个字
func test1() (x, y, z int) {
	return 1, 2, 3
}
