package main

import (
	"fmt"
)

// 函数
/*
Go语言中定义函数使用func关键字，具体格式如下：
func 函数名(参数)(返回值){
	函数体
}

函数存在的意义？
函数是一段代码的封装
把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用即可
使用函数能够让代码结构更清晰、更简洁

// Go语言中函数没有默认参数这个概念
*/

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值的函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数, 没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数，但有返回值的
func f3() int {
	ret := 3
	return ret
}

// 返回值可以命名，也可以不命名
// 命名的返回值就相当于在函数中声明了一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 如果使用的是命名的返回值，return后面可以不写
}

// 多个返回值
func f5() (int, string) {
	return 1, "深圳"
}

// 参数的类型简写：当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数，这里的y可以传，也可以不传
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片 []int
}

// 变量的作用域
var x = 100 // 定义一个全局变量
func f8() {
	// x := 10
	// 函数中查找变量的顺序
	// 1.先在函数内部查找，
	// 2.如果找不到，就往函数的外面查找，一直找到全局
	fmt.Println(x)
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	m, n := f5()
	fmt.Println(m, n)

	f7("下雨了", 1, 2, 3, 4, 5, 6)
	f7("下雨了")

	f8()
}
