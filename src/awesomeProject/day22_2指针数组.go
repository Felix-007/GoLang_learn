package main

import "fmt"

//指针数组:数组里面都是指针
func main_22_2() {
	a := 10
	b := 20
	c := 30
	var p [3]*int = [3]*int{&a, &b, &c}
	fmt.Println(p)
	for i, v := range p {
		fmt.Printf("i%d=%p\n", i, v)
	}
	*p[1] = 200 //*(p[1])=200

	fmt.Println(b)

}

func main22_3() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [3]int{7, 8, 9}

	//p:=&a	//数组指针
	//fmt.Println(p)
	p := [3]*[3]int{&a, &b, &c} //指针数组
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//运算优先级
	(*p[1])[1] = 500
	fmt.Println(b)

}

//拓展的  数组指针(数组里面是切片)
func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{7, 8, 9}

	//p:=&a	//数组指针
	//fmt.Println(p)
	p := &[3][]int{a, b, c} //指针数组
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//运算优先级
	(*p)[1][1] = 500
	fmt.Println(b)
}
