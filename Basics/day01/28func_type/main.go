package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello func f1")

}

func f2() int {
	fmt.Println("Hello func f2")
	return 10

}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) // func()
	b := f2
	fmt.Printf("%T\n", b) // func() int

	f3(f2)

	f7 := f5(f2)
	fmt.Printf("%T\n", f7)

}
