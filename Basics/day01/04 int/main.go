package main

import "fmt"

// 整型

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)  // 101
	fmt.Printf("%b\n", i1)	// 把十进制数转换为二进制  1100101
	fmt.Printf("%o\n", i1)	// 把十进制数转换为八进制  145
	fmt.Printf("%x\n", i1)	// 把十进制数转换为十六进制  65

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)  // 63
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)  // 19088743
	// 查看变量的类型
	fmt.Printf("%T\n", i3)  // int

	// 声明int8类型的变量
	i4 := int8(9)	// 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)  // int8

	// 声明int64类型的变量
	i5 := int64(30)
	fmt.Printf("%T\n", i5)  // int64
}
