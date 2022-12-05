package main

import "fmt"

func main() {
	// 跳出多层循环
	// var flag = false
	// for i := 0; i < 10; i++ {
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break // 跳出内层的for循环
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)
	// 	}
	// 	if flag {
	// 		break // 跳出for循环（外层的for循环）
	// 	}
	// }

	// goto+label实现跳出多层循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx	// 跳到我指定的那个xx标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	xx:	// label标签
	fmt.Println("over")
}
