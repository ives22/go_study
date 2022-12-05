package main

import (
	"fmt"
	"strings"
)

// 闭包
// 闭包是什么？
// 闭包是一个函数，这个函数包含了它外部作用域的一个变量
/*
底层的原理：
1、函数可以作为返回值
2、函数内部查找变量的顺序，先在自己内部找，找不到往外层找
*/

// 要求：
// f1(f2)   在f1中调用执行f2参数
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 解决：定义一个函数，对f2进行包装  闭包
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		// fmt.Println(x + y)
		// 调用f2
		f(x, y)
	}
	return tmp
}

// 闭包2
func adder1() func(int) int {
	var x int = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包3 判断文件后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	ret := f3(f2, 100, 200)
	fmt.Printf("%T\n", ret)
	f1(ret)

	// ret2 := adder1()
	// ret3 := ret2(200)
	// fmt.Println(ret3)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(jpgFunc("呵呵.jpg"))
	fmt.Println(txtFunc("test")) //test.txt

}
