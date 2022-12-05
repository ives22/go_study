package main

import (
	"fmt"
	"os"
)

// 1.文件对象的类型
// 2.获取文件对象的详细信息


// func  f1(a int, b ...int64)  {
// 	fmt.Println(a, b, len(b))
	
// }


func main() {
	fileObj, err := os.Open("/Users/liyanjie/Documents/07 云腾/01 项目文档/14 新希望TCE/show_last_failed.py")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 1. 文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 2.获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("get file info failed, err:", err)
	}
	// fmt.Println(fileInfo)
	fmt.Println(fileInfo.Size())
	fmt.Printf("文件大小是:%dB\n", fileInfo.Size())
	fmt.Printf("文件名是:%v\n", fileInfo.Name())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Sys())
	// fmt.Println(fileInfo.ModTime())
}
