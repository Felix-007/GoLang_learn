package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50}
	//截取 s[起始位置:结束位置:max]
	//起始位置 包含改位置
	//结束位置 不包含这个位置
	//len =结束位置-起始位置
	//max 容量  不能超过原有切片 一般不用
	//[:] 全部复制 s1=s
	//[:2] 0-2
	//[2:] 2-len(s)

	//截取后的切片还是原始切片的一部分
	//修改截取后的切片 原始切片也会被修改

	s1 := s[0:3:5]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := s2[2:5]
	s3[2] = 999
	fmt.Println(s3)
	fmt.Println(s2)

	//切片的 copy 复制的切片和原切片是两个独立的内容
	//var s4 []int
	s4 := make([]int, len(s1))
	copy(s4, s1)
	fmt.Println("切片的复制")
	fmt.Println(s1)
	fmt.Println(s4)

}
