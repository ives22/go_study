package main

import (
	"fmt"
	"io"
	"os"
)

// 在文件中间位置添加内容
func f2() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err:= os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, err:%v\n", err)
		return
	}
	defer tmpFile.Close()
	// 读取源文件，写入到临时文件中
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	// 写入到临时文件
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	// 紧接着把源文件的后续内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF{
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
		}
		tmpFile.Write(x[:n])
	}
	// 源文件后续的也写入了临时文件中
	// 把
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")

}

func main() {
	f2()
}
