package main

import "fmt"

/**
for循环嵌套
break 跳出当前循环 多层 跳出最近的
continue  跳过本次循环 继续下次循环
goto 跳转  goto tag
			tag自定义
*/
func main() {

	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			count++
		}
	}
	fmt.Println(count)

	str := "abc"
	for _, vaule := range str {
		fmt.Printf("index:=%c\n", vaule)
	}

	//cock  hen chicken
	//var cock,hen,chicken int32=0,0,0
	var (
		cock = 0
	)
	//cock:=0
	//hen:=0
	//chicken:=0
	for ; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			if (100-cock-hen)%3 == 0 && cock*5+hen*3+(100-cock-hen)/3 == 100 {
				fmt.Println(cock, hen, 100-cock-hen)
			}
		}
	}

	//var z int= 0
	var (
		z = 0
	)
	//z:=0
	for ; z < 10; z++ {
		fmt.Printf("%d \n", z)
	}

	fmt.Println("aaa")
	goto HAI
	fmt.Println("bbb")
HAI:
	fmt.Println("ccc")

	fmt.Println("aaa")

	fmt.Println("bbb")
OneMoreTime: //形成死循环
	fmt.Println("ccc")
	goto OneMoreTime
}
