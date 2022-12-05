package main

import (
	"fmt"
)

// for 循环
// for 初始语句; 条件表达式; 结束语句{
// 		循环语句
// }

func main() {
	// 基本格式	一般常用这种格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 第二种格式，省略初始语句
	var j = 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 第三种格式，省略初始语句和结束语句
	var k = 5
	for k < 10 {
		fmt.Println(k)
		k++
	}

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range 循环
	s := "Hello 小白"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
		// fmt.Println(i, v)
	}

	//	匿名变量使用示例（哑元变量，不想用到的都直接扔给它）
	s1 := "沙河有沙河"
	for _, v := range s1 {
		// fmt.Println(v)
		fmt.Printf("%c\n", v)
	}

	// 流程控制之跳出for循环 ---break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 流程控制之跳出for循环 ---container
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
