package main

import "fmt"

type hero struct {
	name  string
	age   int
	power int
}

//结构体map
func main_9() {
	m1 := make(map[int]hero)
	fmt.Println(m1)
	m2 := map[int]hero{
		1: {"超人1", 18, 100},
		2: {"超人2", 29, 239}}
	fmt.Println(m2)

	m1[101] = hero{name: "钢铁侠", age: 30, power: 100}
	m1[103] = hero{name: "美队", age: 130, power: 110}
	fmt.Println(m1)

	delete(m2, 1)
	fmt.Println(m2)
}

//结构体切片
func main() {
	//value类型是一个切片
	m := make(map[int][]hero)

	m[101] = []hero{
		{name: "钢铁侠", age: 30, power: 100},
		{name: "火箭人", age: 30, power: 100}}

	m[101] = append(m[101], hero{name: "蜘蛛侠", age: 30, power: 20})
	m[102] = []hero{{"美队", 30, 90}}
	m[102] = append(m[102], hero{"冬兵", 30, 80})
	fmt.Println(m)
}
