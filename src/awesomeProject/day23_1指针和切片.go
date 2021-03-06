package main

import "fmt"

//切片的指针  *[]int
func main_23() {
	/*
				--------------------------------------
				|		|        | s 	|       |
				|1 2 3  |		 |0xff00|		|0xee00
				|____________________________________
				0xff00				0xee00       0x10ee
		                                         p:=&s
	*/
	s := []int{1, 2, 3}
	p := &s //切片指针
	fmt.Println(p)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

	//通过指针操作元素
	(*p)[1] = 222
	fmt.Println(s)

	test_23(p)

	fmt.Println(s)
}

func test_23(a *[]int) {
	(*a)[2] = 333
	*a = append(*a, 4, 5, 6)
}

//指针切片:切片里面是指针
func main() {
	a := 10
	b := 20
	c := 30

	p := []*int{&a, &b, &c}
	fmt.Println(p)
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

}
