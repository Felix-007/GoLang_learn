package main

import "fmt"

//map 地址传递
//和切片相似  形参的修改会影响到实参
func main() {
	m := map[int]string{}
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"
	mapdemo(m)
	fmt.Println(m)
	fmt.Printf("%p\n", m)

	var m2 map[int]string = mapdemo2(m)
	fmt.Println("m2=", m2)

}

func mapdemo(m map[int]string) {
	m[4] = "e"
	fmt.Println(m)
	fmt.Printf("%p\n", m)
}

func mapdemo2(m map[int]string) map[int]string {
	m[6] = "g"
	fmt.Println(m)
	fmt.Printf("%p\n", m)
	return m
}
