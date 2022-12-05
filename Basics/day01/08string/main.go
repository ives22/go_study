package main

import (
	"fmt"
	"strings"
)

// 字符串
// Go语言中字符串是用双引号包裹的！！！
// Go语言中单引号包裹的是字符！！！

func main() {
	// 字符串
	s1 := "Hello 小白"
	// 单独的字母、汉字、符号表示一个字符
	c1 := 'h'
	c2 := '1'
	c3 := '沙'

	fmt.Println("s1:", s1) // s1: Hello 小白
	fmt.Println("c1:", c1) // c1: 104
	fmt.Println("c2:", c2) // c2:  49
	fmt.Println("c3:", c3) // c3:  27801

	// \ 本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "D:\\Go\\src\\code"
	path1 := "\"D:\\Go\\src\\code\""
	fmt.Println(path)  // D:\Go\src\code
	fmt.Println(path1)  // "D:\Go\src\code"

	s2 := "I' m ok"
	fmt.Println(s2)

	// 多行的字符串
	s3 := `
		111
		222
		333
	`
	fmt.Println(s3)

	s4 := `D:\\Go\\src\\code`
	fmt.Println(s4)

	// 字符串的常用操作
	fmt.Println(len(s4))

	// 字符串拼接，方式一
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	// 字符串拼接，方式二
	ss1 := fmt.Sprintf("%s%s", name, world)
	// fmt.Printf("%s%s", name, world)
	fmt.Println(ss1)

	// 字符串的分割  strings.Split(要分割的字符串，分割符)
	ret := strings.Split(path, "\\")
	fmt.Println(ret)        // [D: Go src code]
	fmt.Printf("%T\n", ret) // 切割后的类型：[]string

	// 判断是否包含
	fmt.Println(strings.Contains(ss, "理性")) // false
	fmt.Println(strings.Contains(ss, "理想")) // true

	// 前缀和后缀判断
	// 判断前缀，是否以什么开头
	fmt.Println(strings.HasPrefix(ss, "理想")) // true
	// 判断后缀，是否以什么结尾
	fmt.Println(strings.HasSuffix(ss, "理想")) // false

	// 判断子串的位置
	s5 := "abcdeb"
	fmt.Println(strings.Index(s5, "c"))     // 2   判断字符串c在s5中的位置
	fmt.Println(strings.LastIndex(s5, "b")) // 5

	// 拼接
	s6 := []string{"D:", "Go", "src", "code"}
	fmt.Println(strings.Join(s6, "+")) // D:+Go+src+code

}
