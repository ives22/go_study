package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主 goroutine去执行
func main() {
	startTime := time.Now()
	// fmt.Println(startTime)

	// go hello(i) // 开启一个单独的goroutine去执行hello函数（任务）
	for i := 0; i < 100; i++ {
		go func(i int){
			fmt.Println(i)	// 用的是函数参数的i
		}(i)
	}

	fmt.Println("main")
	time.Sleep(time.Second)
	// main函数结束了 由main函数启动的goroutine也都结束了
	endTime := time.Now()
	fmt.Println("耗时：", startTime, endTime)
}

