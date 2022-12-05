package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http Client
func main() {
	resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=18")
	if err != nil {
		fmt.Println("get url failed, err:%v", err)
		return
	}
	// 从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:%v", err)
		return
	}
	fmt.Println(string(b))
}
