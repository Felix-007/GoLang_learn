package main

import (
	"fmt"
	"math"
	"net"
)

var sf bool
var hp = 100

var attack = 40
var defence = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate

func main() {
	str, ai := "123a", 123

	fmt.Println("Hello, World!")
	fmt.Println(sf)
	fmt.Println(hp)
	fmt.Println(damageRate)
	fmt.Println(damage)
	fmt.Printf("%.2f\n", math.Pi)

	fmt.Println(str)
	fmt.Println(ai)

	fmt.Println("运行测试方法")
	test()

}

var (
	conn net.Conn
	err  error
)

func test() {

	// conn, err := net.Dial("tcp", "127.0.0.1:8080")
	// fmt.Println(conn)
	// fmt.Println(err)

	a, b := 100, 200
	a = a ^ b
	b = b ^ a
	a = a ^ b

	fmt.Println(a, b)

	c, _ := GetData()
	_, d := GetData()
	fmt.Println(c, d)

	const str = `第一行
	第二行
第三行\r\n
`
	fmt.Println(str)
}

//GetData 123
func GetData() (int, int) {
	return 100, 200
}
