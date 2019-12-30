package main

import (
	"fmt"
	"io"
	"os"
)

/**
新建文件(会自动打开)
*/
func main33() {
	file, err := os.Create("./src/awesomeProject/file.txt")
	if err != nil {
		//1.不存在
		//2.文件权限
		//3.程序打开文件上限
		fmt.Println("文件创建失败")
		return
	}
	//写文件  windows \r\n linux \n
	file.WriteString("hello world \r\n")
	file.WriteString("鬼魅幻影")

	//多用defer
	defer file.Close()
}

func main33_2() {
	file, err := os.Create("./src/awesomeProject/file.txt")
	if err != nil {
		//1.不存在
		//2.文件权限
		//3.程序打开文件上限
		fmt.Println("文件创建失败")
		return
	}
	//s:=[]byte{'h','i','g','o'}

	//count,err1:=file.Write(s)
	count, err1 := file.Write([]byte("hello go"))
	if err1 != nil {
		fmt.Println(err1)
		return
	} else {
		fmt.Println("写入成功", count)
	}

	defer file.Close()
}

func main() {
	file, err := os.Create("./src/awesomeProject/file.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	//获取文件光标流位置  一般使用io包里面的
	count, _ := file.Seek(0, io.SeekEnd)
	fmt.Println(count)
	//指定位置写入
	file.WriteAt([]byte("aaaaaaaa"), count)
	file.WriteAt([]byte("bbb"), count)
	file.WriteAt([]byte("dd"), 7)
	fmt.Println(count)
	defer file.Close()

}
