package main

import (
	"fmt"
	"io"
	"os"
)

//
//
func main() {
	var srcFliename string
	var copyFilename string
	fmt.Println("源文件名称:")
	fmt.Scan(&srcFliename)
	fmt.Println("目的文件名称:")
	fmt.Scan(&copyFilename)

	if srcFliename == copyFilename {
		fmt.Println("文件名不能重名")
		return
	}

	//读文件
	sf, err1 := os.Open(fmt.Sprintf("./src/awesomeProject/%s", srcFliename))
	if err1 != nil {
		fmt.Println("打开失败")
		return
	}
	//新建目的文件
	df, err2 := os.Create(fmt.Sprintf("./src/awesomeProject/%s", copyFilename))
	if err2 != nil {
		fmt.Println("新建copy文件失败", err2)
		return
	}
	//延时关闭
	defer sf.Close()
	defer df.Close()

	//读取 写入
	buf := make([]byte, 1024*2)
	for {
		n, err := sf.Read(buf)
		//fmt.Println(string(buf))
		if err != nil {
			//fmt.Println("读取失败2",err)
			if err == io.EOF {
				break
			}
		}
		df.Write(buf[:n])

	}
	fmt.Println("复制完成")
}
