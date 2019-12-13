package main

import "fmt"

/**
流程控制
if if支持一个初始化语句 初始化语句和条件用;分隔
switch
case支持多条件匹配
switch后面支持一个初始化语句 初始化语句和条件用;分隔
switch后面可以不写条件 case里面写条件
switch后面可以只写初始化语句+;  caes后面写条件
1.case 1,2,3
2.case 默认有break
3.fallthrough 会执行下一个case 这个有点变态

for循环
range
//两个变量接收 一个下标 一个data
//可以一个变量接收(下标) 可以使用匿名变量接收  丢弃一个
for index,data:=range str

*/
func main() {
	var score int
	//fmt.Scan(&score)
	if score == 700 {
		fmt.Println("恭喜")
	} else {
		fmt.Println("挫败")
	}

	//if支持一个初始化语句 初始化语句和条件用;分隔
	if soc := 200; soc == 700 {
		fmt.Println("恭喜2")
	} else if soc < 700 {
		fmt.Println("<700")
	}

	i := 1
	switch i {
	case 1:
		fmt.Println("A")
		fallthrough
	case 2, 3: //cast支持多条件匹配
		fmt.Println("B")
		fallthrough
	case 4:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	switch j := 1; j {
	case 1:
		fmt.Println("A")
		fallthrough
	case 2, 3: //cast支持多条件匹配
		fmt.Println("B")
		fallthrough
	case 4:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	switch j := 3; {
	case j == 1:
		fmt.Println("AA")
	case j == 2, j == 3: //cast支持多条件匹配
		fmt.Println("BA")
	default:
		fmt.Println("DA")
	}

	/**
	for循环
	*/
	sum := 0
	for k := 1; k <= 10; k++ {
		sum += k
	}
	fmt.Println(sum)

	str:="abcd"

	for index,data:=range str{
		fmt.Printf("i%d=%c\t",index,data)
	}


}
