package main

import "fmt"

// defer
// defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接）
// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值复制
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func(){
		x++	 // 修改的是X不是返回值
	}()
	return x // 1. 返回值赋值，2. defer，3. 真正的RET指令
}


func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")	// defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("呵呵呵")	// 一个函数中可以有多个defer语句
	defer fmt.Println("哈哈哈")	// 多个defer语言按照先进后出（后出先进）的顺序延迟执行
	fmt.Println("end")

	fmt.Println(f1())
}

func main() {
	deferDemo()
}

