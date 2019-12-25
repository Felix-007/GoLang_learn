package main

import "fmt"

func main22() {
	arr := [3]int{1, 2, 3}
	//数组指针
	var p *[3]int
	p = &arr
	p1 := &arr
	fmt.Println(p, p1)
	fmt.Printf("type=%T,p=%p,p1=%p\n", p, p, p1)

	//通过指针操作数据
	(*p)[0] = 111 //(*指针变量)[下标]
	p[1] = 222    //指针变量[下标]
	fmt.Println(*p)

}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	p := &arr
	p2 := &arr[0]
	fmt.Println(p, p2)
	fmt.Printf("p=%p,p2=%p\n", p, p2)
	fmt.Printf("p=%T,p2=%T\n", p, p2)
	test22(p)
	fmt.Println(p, p2)

}

func test22(p *[5]int) {
	p[4] = 222
}
