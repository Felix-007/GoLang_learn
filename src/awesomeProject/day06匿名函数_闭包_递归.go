package main

import "fmt"

/**
匿名函数
	func(参数)(返回值){
	}(参数)

//闭包的使用
//函数的递归
*/
func main() {
	i := 10
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	//匿名函数
	fmt.Println("匿名函数")
	num := 9
	f := func() {
		num++
		fmt.Println(num)
	}
	f()
	fmt.Println(num)

	//func(){
	//	num++
	//	fmt.Println(num)
	//}()
	//fmt.Println(num)

	//type fun1 func()
	//var f1 fun1
	//f1=f
	//f1()

	//匿名函数参数的传递
	fmt.Println("匿名函数参数的传递")
	func(a, b int) {
		count := a + b
		fmt.Println(count)
	}(3, 4)
	ff := func(a, b int) {
		count := a + b
		fmt.Println(count)
	}
	ff(3, 4)

	//有返回值
	max, min := func(c, d int) (max, min int) {
		if c > d {
			max = c
			min = d
		} else {
			max = d
			min = c
		}
		return
	}(10, 20)
	fmt.Println("max=", max, "min=", min)

	//闭包 所有匿名函数都是闭包
	fmt.Println("闭包")
	a1 := test_1()
	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(a1())

	fmt.Println(test_2())

	//递归的使用
	fmt.Println("递归的使用")
	test_3(0)
	fmt.Println("1-100的和-", test_4(10))
}

func test_1() func() int {
	//x++
	x := 0
	//var x int
	//f=
	return func() int {
		x++
		return x
	}

}

func test_2() int {
	//z=10
	return 2
}

//递归函数
//1:需要跳出条件
func test_3(a int) {
	if a == 3 {
		fmt.Println(a)
		return
	}
	test_3(a + 1)
	fmt.Println(a)
}

func test_4(d int) int {
	if d == 1 {
		return 1
	}
	return d + test_4(d-1)
}
