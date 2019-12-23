package main

import "fmt"

//map的定义
func main06() {
	//1.
	//map[key类型]值类型
	var m map[int]string
	fmt.Println(m)
	fmt.Println(len(m))

	//2.
	m2 := make(map[int]string, 3)
	fmt.Println(m2)
	fmt.Println(len(m2))
	m2[1] = "张三"
	m2[2] = "李四"
	m2[3] = "王五"
	m2[4] = "张三"
	fmt.Println(m2)

	//3.
	m3 := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	fmt.Println(m3)
	fmt.Printf("%p\n", m3)
}

func main() {
	//map的key一般选用 == != 的基本数据类型
	//m:=make(map[string][3]int)
	m := map[string][3]int{"a": {1, 2, 3}, "b": {1, 2, 3}, "c": {1, 2, 3}}
	fmt.Println(m)
	m["a"] = [3]int{100, 99, 98}
	m["b"] = [3]int{2, 3, 98}
	m["c"] = [3]int{89, 23, 98}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	//判断键值 是否存在是否存在
	value, ok := m["a"]
	fmt.Println(value, ok)

	//删除map中的元素
	delete(m, "a")
	fmt.Println(m)

}
