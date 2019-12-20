package main

import "fmt"

func main() {
	var arr [2][3]int
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))

	//1:全部初始化
	arr1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)

	var c [2][3]int = [2][3]int{{1, 2, 3}}
	fmt.Println(c)
}
