package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换成 字符切片
	str := "hello world"
	slice := []byte(str)
	fmt.Println(slice)
	fmt.Printf("%c\n", slice[2])

	//字符切片转成 字符串
	slice2 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(slice2)
	fmt.Println(string(slice2))

	//其他类型转成字符串 strconv.Format
	str2 := strconv.FormatBool(true)
	fmt.Println(str2)
	fmt.Printf("类型:%T\n", str2)
	//整型   进制 2-36(10个数字+26个字母)
	str3 := strconv.FormatInt(322, 10)
	str4 := strconv.FormatInt(322, 2)
	fmt.Println(str3)
	fmt.Println(str4)
	//		输出方式  小数    类型
	str5 := strconv.FormatFloat(3.14, 'f', 4, 64)
	fmt.Println(str5)

	//整型转字符串
	str6 := strconv.Itoa(123)
	fmt.Println(str6)

	iint, err := strconv.Atoi("123")
	fmt.Println(iint)

	//字符串转其他类型
	bb, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(bb)
	}

	s := make([]byte, 0)
	s = strconv.AppendBool(s, true)
	s = strconv.AppendInt(s, 123, 2)
	s = strconv.AppendFloat(s, 3.14159, 'f', 4, 64)
	//字符串
	s = strconv.AppendQuote(s, "abc")
	fmt.Println(s)
	fmt.Println(string(s))

}
