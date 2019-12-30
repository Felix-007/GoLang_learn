package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Open 只读打开
	//file,err:=os.Open("./src/awesomeProject/file.txt")

	//os.OpenFile(文件名,打开方式,权限)
	file, err := os.OpenFile("./src/awesomeProject/file.txt", os.O_RDWR, 7)

	if err != nil {
		fmt.Println("打开失败")
		return
	}
	//覆盖写
	file.WriteString("你好")

	defer file.Close()
}
