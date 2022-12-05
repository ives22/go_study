package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	// str := `<h1 style="color: red">hello 成都！</h1>`
	// w.Write([]byte(str))

	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都是放在URL上（query param）, 请求体中是没有数据的。
	queryParam := r.URL.Query()  // 自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	
	fmt.Println(r.URL)  // 获取客户端请求URL
	fmt.Println(r.Method)  // 获取请求方法
	fmt.Println(ioutil.ReadAll(r.Body))  // 获取客户端请求的body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/go/15", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
