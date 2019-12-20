package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main04() {

	var a [9]int = [9]int{4, 2, 8, 6, 5, 7, 1, 3, 9}
	fmt.Println(a)
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}

func main05() {
	//随机数种子  用time来变换
	rand.Seed(time.Now().Unix())
	//0-9
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))

}

//双色球
func main() {
	rand.Seed(time.Now().UnixNano())
	var red [6]int
	for i := 0; i < len(red); i++ {
		v := rand.Intn(33) + 1
		for j := 0; j < i; j++ {
			//数据重复
			if v == red[j] {
				v = rand.Intn(33) + 1
				j = -1
			}
		}
		red[i] = v
	}
	fmt.Println("红球=", red, "篮球=", rand.Intn(16)+1)
}
