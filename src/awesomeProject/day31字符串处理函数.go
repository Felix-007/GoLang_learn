package main

import (
	"fmt"
	"strings"
)

/**
使用strings处理字符串
	//查找
	1:bool 类型 :=strings.Contains(被查找的字符,查找字符串)
	2:int 类型 :=strings.Index(被查找的字符,查找字符串)
	//分割
	[]string 切片类型:=strings.Split(切割字符串,标志)
	//组合
	string 类型:=strings.Join(切片,标志)
	//重复
	 string 类型:=strings.Repeat(字符串,次数)
	//替换
	 string 类型:=strings.Replace(原字符串,oldstring,newstring,次数)
	//去掉内容(只能替换头尾)
	string 类型:=strings.Trim(字符串,去掉的内容)
	[]string 切片类型:=strings.Fields(字符串)//去掉空格
*/
func main() {
	str1 := "hello world"
	str2 := "l"
	//Contains
	bool := strings.Contains(str1, str2)
	fmt.Println(bool)

	//Join 切片字符串连接
	s1 := []string{"123", "456", "789"}
	str3 := strings.Join(s1, ",")
	fmt.Println(s1)
	fmt.Println(str3)

	//Index 查找下标 第一次出现的位置
	index := strings.Index(str1, str2)
	fmt.Println(index) //-1没找到

	//Repeat 一个字符串 重复N此
	str4 := strings.Repeat(str1, 3)
	fmt.Println(str4)

	//Replace 字符串的替换
	str5 := "字符串的替换,字符串的替换"
	str6 := strings.Replace(str5, "替换", "XX", 1) //n表示替换次数 如果n<0  全部替换
	fmt.Println(str6)

	//Split
	str7 := "139-2345-6574"
	s2 := strings.Split(str7, "-")
	fmt.Println(s2)
	fmt.Println(strings.Join(s2, ""))

	//Trim	去除字符串头尾的内容
	str8 := "======a====u===ok===="
	str9 := strings.Trim(str8, "=")
	fmt.Println(str9)

	//Fields 去除空格 返回切片 一般用于统计字符串个数
	str10 := "  are  you  ok  "
	s3 := strings.Fields(str10)
	fmt.Println(s3)
	fmt.Println(strings.Join(s3, ""))

}
