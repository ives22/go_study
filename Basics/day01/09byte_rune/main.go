package main

import (
	"fmt"
)

// 字符

func main() {
	// byte	uint8的别名	只能表示常见的 ASCII码
	// rune	int32的别名	代表一个UTF-8字符

	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	s := "Hello沙河"
	// len()求得是byte字节的数量
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i]) // %c: 字符
	}

	fmt.Println()
	for _, r := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", r)
	}

	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	// 字符串修改
	s2 := "白萝卜"      // => '白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	d1 := "红"
	d2 := '红'                            // rune类型（int32）
	fmt.Printf("d1:%T  d2:%T\n", d1, d2) // c1:string  c2:int32
	d3 := "H"                            // string
	d4 := 'H'                            // int32
	fmt.Printf("d3:%T  d4:%T\n", d3, d4)
	fmt.Printf("%d\n", d4) // 72

	// 类型转换
	n := 10 // int
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Printf("%T\n", f) // float64
}
