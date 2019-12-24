package main

import "fmt"

/**
结构体map 地址传递 形参改变实参
结构体map数组 地址传递 形参改变实参
结构体map切片 地址传递 形参改变实参
结构体   值传递  形参不改变实参
*/

type superhero struct {
	name  string
	age   int
	power int
}

func main() {
	m1 := make(map[int]superhero) //结构体map
	m1[101] = superhero{name: "钢铁侠", age: 30, power: 100}
	m1[102] = superhero{name: "美队", age: 130, power: 110}

	m2 := make(map[int][2]superhero) //结构体数组
	m2[101] = [2]superhero{
		{name: "钢铁侠", age: 30, power: 100},
		{name: "美队", age: 130, power: 110}}

	m3 := make(map[int][]superhero) //结构体切片
	m3[101] = []superhero{{name: "钢铁侠", age: 30, power: 100},
		{name: "美队", age: 130, power: 110}}

	m4 := superhero{"结构体", 20, 48}

	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)
	fmt.Println("m3=", m3)
	fmt.Println("m4=", m4)

	test_m1(m1)
	test_m2(m2)
	test_m3(m3)
	test_m4(m4)

	fmt.Println("方法后m1=", m1)
	fmt.Printf("%p\n", m1)
	fmt.Println("方法后m2=", m2)
	fmt.Printf("%p\n", m2)
	fmt.Println("方法后m3=", m3)
	fmt.Printf("%p\n", m3)
	fmt.Println("方法后m4=", m4)
	fmt.Printf("%p\n", &m4)

}

//结构体map作为函数参数
func test_m1(m1 map[int]superhero) {
	//err
	//m1[102].power=98
	temp := m1[102]
	temp.power = 98
	m1[102] = temp
	fmt.Println("方法内m1=", m1)
	fmt.Printf("%p\n", m1)
}

//结构体数组
func test_m2(m2 map[int][2]superhero) {
	//err
	//m1[102].power=98
	temp := m2[101]
	temp[0].power = 98
	m2[101] = temp
	fmt.Println("方法内m2=", m2)
	fmt.Printf("%p\n", m2)
}

//结构体切片
func test_m3(m3 map[int][]superhero) {
	//err
	//m1[102].power=98
	temp := m3[101]
	temp[0].power = 98
	m3[101] = temp
	fmt.Println("方法内m3=", m3)
	fmt.Printf("%p\n", m3)
}

//结构体作为函数参数
func test_m4(m4 superhero) {
	m4.power = 98
	fmt.Println("方法内m4=", m4)
	fmt.Printf("%p\n", &m4)
}
