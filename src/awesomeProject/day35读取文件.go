package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main_35() {
	file, err := os.Open("./src/awesomeProject/file.txt")
	if err != nil {
		fmt.Println("打开失败")
	}
	//var buf []byte
	buf := make([]byte, 10)
	//读取长度  块读取
	n, _ := file.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf))

	n, _ = file.Read(buf)
	fmt.Println(string(buf))

	//for 循环读取
	for {
		n2, err := file.Read(buf)
		//文件结尾
		if err == io.EOF {
			//fmt.Println("读完了")
			break
		}
		fmt.Println(string(buf[:n2]))
	}

	defer file.Close()
}

//行读取
func main() {
	file, err := os.Open("./src/awesomeProject/file.txt")
	if err != nil {
		fmt.Println("打开失败")
	}
	defer file.Close()
	//创建文件缓存区
	r := bufio.NewReader(file)
	//行读取	 \n
	//s,_:=r.ReadBytes('\n')
	//fmt.Println(string(s))
	//s,_=r.ReadBytes('\n')
	//fmt.Println(string(s))
	//for {
	//	//'\n' 结束
	//	s, err1 := r.ReadBytes('\n')
	//	fmt.Println(string(s))
	//	if err1 != nil {
	//		if err1 == io.EOF {
	//			break
	//		}
	//		fmt.Println(err1)
	//	}
	//}

	for {
		str, err2 := r.ReadString('\n')
		fmt.Println(str)
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
		}
	}
}
