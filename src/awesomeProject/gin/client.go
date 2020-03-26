package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// 调用最基本的POST,并获得返回值
	resp, _ := http.Post("http://0.0.0.0:8888/test", "", strings.NewReader("[{\"carNum\": \"京A88888\",\"yldm\": \"123\"}, {\"carNum\": \"京A88888\",\"yldm\":\"123\"}]"))
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("body:", string(body))
}
