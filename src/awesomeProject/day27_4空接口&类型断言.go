package main

import (
	"fmt"
)

/**
空接口 interface{}
里面存任何东西
判断取出的东西通过类型断言:
一:value,ok:=elemet.(T)
二:switch测试
*/
func main27_4() {
	//空接口	可以接收任意数据类型
	//	fmt.Printf 参数就是接口类型的
	var i interface{}
	fmt.Println(i) //<nil>
	//fmt.Printf("%V\n",i)
	i = 10
	fmt.Printf("%T\n", i) //int
	i = 3.14
	fmt.Printf("%T\n", i) //float64
	i = "abc"
	fmt.Printf("%T\n", i) //sting

}

func test27() {
	fmt.Println("我是个函数")
}
func main_27_4() {
	//空接口类型的切片
	var s []interface{}
	fmt.Printf("%T\n", s)
	s = append(s, 10, 2, "aaa", test27)
	fmt.Println(s)

	for i, v := range s {
		fmt.Println(i, "是", v)
	}

}

func main() {
	//接口断言
	var i interface{}
	i = 10
	//value,ok:=elemet.(T)
	//值,值得判断:=接口变量.(数据类型)
	v, ok := i.(int)
	fmt.Println(v, ok)

	i = test27
	fmt.Printf("%T\n", i)
	value, ok := i.(func())
	//value()
	fmt.Println(value, ok)

	//空接口切片
	var s []interface{}
	s = append(s, 10, 3.12, "aaa", test27)
	for _, v := range s {
		if data, ok := v.(int); ok {
			fmt.Println("整型", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("浮点型", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("字符串", data)
		} else if data, ok := v.(func()); ok {
			//fmt.Println("函数",data)
			data()
		}
	}

	for _, v := range s {
		switch value := v.(type) {
		case int:
			fmt.Println("整型", v)
		case float64:
			fmt.Println("浮点型", v)
		case string:
			fmt.Println("字符串", v)
		case func():
			value()

		}
	}

}
