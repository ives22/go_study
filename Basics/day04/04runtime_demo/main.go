package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func f(){
	// 获取程序当前执行的信息
	pc, file, line, ok := runtime.Caller(0)
	// pc, file, line, ok := runtime.Caller(1)
	// pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f1(){
	f()
}

func main() {
	f1()
}
