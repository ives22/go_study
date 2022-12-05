package main


// 导入语句
import "fmt"

// 函数外只能放置标识符（变量、常量、函数、类型）的声明
// fmt.Println("测试下")	非法，这是不行的，只能放在函数里面

// 单行注释

/*
多行注释
*/

// Go语言函数外的语句必须以关键字开头

// 如果要编译可执行文件，必须要有main包和main函数， main函数是入口函数，它没有参数，也没有返回值
func main()  {
	// 函数内部定义的变量必须使用
	fmt.Println("Hello World")
}


