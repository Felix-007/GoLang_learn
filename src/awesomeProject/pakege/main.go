package main

import (
	"awesomeProject/pakege/pUserinfo"
	"fmt"
)

func main() {
	fmt.Println("main")
	test()
	//调用其他文件夹的
	//包名.函数名
	pUserinfo.Userlogin()
}
