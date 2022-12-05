package main

import (
	"fmt"
	"os"
)

// os.Args 获取命令行参数
func main() {
	if len(os.Args) > 0 {
		fmt.Printf("%#v\n", os.Args)
		fmt.Printf("%T\n", os.Args) // []string
		fmt.Println(os.Args[1])     // 获取第一个参数
		// 循环获取所有参数
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%s\n", index, arg)
		}
	}
}
