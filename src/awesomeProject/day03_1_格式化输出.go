package main

import "fmt"

/**
格式化输出
%T 数据类型
%c 字符串
%d 整形
%s 字符串
%f 浮点数   %.2f 保留小数点两位
%v 自动匹配格式
%p 指针地址  &取地址 *取地址的值

用户输入 Scan(&a)

类型转换 float64(a)

类型别名 type bigint int64 -int64去别名叫bigint
*/
func main() {
	a := 10
	b := "abc"
	c := 'd'
	d := 3.13532
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)
	fmt.Printf("%d,%s,%c,%.2f\n", a, b, c, d)
	fmt.Printf("%v,%v,%v,%v\n", a, b, c, d)

	//用户输入
	//var v string
	//fmt.Println("请输入数据:")
	//fmt.Scanf("%d", &v)
	//fmt.Println(v)

	//fmt.Scan(&v)
	//fmt.Println(v)

	q, w, e := 50, 30, 24
	sum := q + w + e
	fmt.Println(float64(sum) / 3.0)

	type bigint int64
	var big bigint = 213
	fmt.Println(big)
	fmt.Printf("%T\n",big)

	type (
		long int64
		char byte
	)
	var rr long=123
	var yy char ='a'
	fmt.Println(rr,yy)
	fmt.Printf("%T,%T\n",rr,yy)

}
