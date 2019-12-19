package main

import "fmt"

/*
函数 func 函数名(参数) 返回值{
		函数体
	}

不定参数  args ...int
 1.放在其他固定参数后面
 2.不定参数在调用的时候可以不传值

函数的嵌套使用
不定参数在 多函数传递
args...  arg[:]...  arg[0:2]...

函数返回值 (a,b,c int)

*/
func main() {
	//函数名
	run(1, 2)
	sum(3, 4)

	sum2(1, 2, 3, 4, 5, 6, 7, 8)
	sum2()
	fmt.Println("不定参数传递")
	sum3(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println("函数的返回值")
	fmt.Println(fanghui(1, 2, 3, 4, 5))
	fmt.Println(fanghui2(1, 2, 3, 4, 5))
	result1, _, result2 := fanghui2(2, 3, 4, 5)
	fmt.Println(result1, result2)

}

func fanghui(args ...int) int {
	i, sum := 0, 0

	for _, v := range args {
		sum += v
		i++
	}
	return sum
}

//							这里申名了返回值 return 可以不写返回的名字
func fanghui2(args ...int) (str string, i, sum int) {
	i, sum = 0, 0
	for _, v := range args {
		sum += v
		i++
	}
	str = "avd"
	return
}

func run(a, b int) {
	fmt.Printf("a=%d,b=%d\n", a, b)
}

func sum(c, d int) {
	fmt.Println(c + d)
}

func sum2(paras ...int) {
	fmt.Println(paras)
	sum := 0
	for _, v := range paras {
		sum += v
	}
	fmt.Println(sum)
}

func sum3(args ...int) {
	sum2(args...)
	sum2(args[2:4]...)
	sum2(args[:]...)
}
