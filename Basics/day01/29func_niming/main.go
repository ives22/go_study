package main

import "fmt"

// 匿名函数
// 匿名函数就是没有名字的函数

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {

	// 函数内部没有办法声明带名字的函数
	// 匿名函数
	f1 := func (x, y int)  {
		fmt.Println(x + y)
	}
	f1(10, 20)  // 通过变量调用匿名函数

	// 如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int){
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)
}



