package main

import "fmt"

// panic 和 recover
// recover()必须搭配defer使用
// defer一定要在可能引用panic的语句之前定义

func funcA() {
	fmt.Println("Func A")
}

func funcB() {
	// 刚刚打开了数据库连接
	defer func() {
		err := recover()	// 通过recover尝试修复错误，
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误！！！ ")	// 程序崩溃退出
	fmt.Println("Func B")
}

func funcC() {
	fmt.Println("Func C")
}

func main() {
	funcA()
	funcB()
	funcC()
}

