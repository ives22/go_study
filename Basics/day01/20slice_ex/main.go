package main

import "fmt"

// 切片

func main() {
	var a = make([]int, 5, 10)	// 创建切片，长度为5，容量为10
	for i := 0; i < 10; i++ {
		a = append(a, i)	
	}
	fmt.Println(a)	// [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}
