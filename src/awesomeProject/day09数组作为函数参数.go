package main

import "fmt"

//数组作为函数值传递
//形参和实参是不同的地址 不影响值
func BubboleSort(a [10]int) [10]int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	//b=a
	fmt.Println(a)
	return a
}
func main() {
	a := [...]int{9, 1, 2, 3, 6, 4, 6, 8, 3, 10}
	fmt.Println(BubboleSort(a))
}
