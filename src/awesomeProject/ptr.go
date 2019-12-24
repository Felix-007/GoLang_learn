package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("xy", "", "process mode")

//指针的操作
func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p \n", &cat, &str)

	house := "malibu point 10088,90257,t1es41t"

	ptr := &house

	fmt.Printf("ptr type:%T\n", ptr)
	fmt.Printf("ptr type:%p\n", ptr)

	value := *ptr
	fmt.Printf("value type:%T\n", value)
	fmt.Printf("value type:%s\n", value)

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	flag.Parse()

	fmt.Println(*mode)
	fmt.Println(mode)
	fmt.Printf("%T\n", mode)

	ss := new(string)
	fmt.Printf("a123:%p\n", ss)
	fmt.Println(*ss)
	*ss = "啊啊啊啊"
	fmt.Println(*ss)

	al := "123"
	fmt.Println(al)

	fmt.Printf("cal=%d\n", calc(1, 3))
}

func swap(a, b *int) {

	*a, *b = *b, *a
	// mid := *a
	// *a = *b
	// *b = mid

}

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}
