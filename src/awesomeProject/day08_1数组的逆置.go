package main

import "fmt"

func main03() {
	a := [...]int{-1, -2, -3, -4, -5}
	max, min := a[0], a[0]
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
		if max < a[i] {
			max = a[i]
		}
		if min > a[i] {
			min = a[i]
		}
	}

	fmt.Println("max=", max)
	fmt.Println("min=", min)
	fmt.Println("sum=", sum)
	fmt.Println("avg=", float64(sum)/float64(len(a)))

}

//数组的逆置
func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	j := len(a) - 1
	for i < j {
		//if i>=j{
		//	break
		//}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	fmt.Println(a)
}
