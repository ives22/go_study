package main

import "fmt"

// pointer 指针

func main() {
	// 1. &:获取地址
	n := 18
	p := &n		
	fmt.Println(p)	
	fmt.Printf("%T\n", p)	// *int int类型的指针
	// 2. *:根据地址取值
	m := *p		// 根据内存地址取值
	fmt.Println(m)	
	fmt.Printf("%T\n", m)



	// 错误的写法
	// var a * int // nil pointer
	// *a = 100
	// fmt.Println(*a)

	// new函数申请一个内存地址
	var a2 = new(int)
	fmt.Println(*a2)	// 0
	*a2 = 100
	fmt.Println(*a2)	// 100
}
