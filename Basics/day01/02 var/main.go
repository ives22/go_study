package main

// 导入语句
import "fmt"

// Go语言中推荐使用驼峰式命名	var studentName string
// 全局变量可以声明不使用，局部变量声明后，必须使用，否则编译不过去

//  声明变量
// var name string
// var age int
// var isOk bool

// 批量声明（全局变量声明推荐这种方式，一个var下面定义多个变量）(常用方式)
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)



func main() {
	// 变量初始化
	name = "理想"
	age = 18
	isOk = true
	// Go语言中非全局变量声明后必须使用，不使用就编译不下去
	fmt.Print(isOk)                // 在终端输出要打印的内容
	fmt.Println()                  // 打印换行
	fmt.Printf("name: %s\n", name) // %s：占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)               // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "小白"
	fmt.Println(s1)

	// 类型推导（根据值判断变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	var ss2 = 10
	fmt.Println(ss2)

	// 简短变量声明（只能在函数里面使用），可以根据变量的值判断是什么类型	(常用方式)
	s3 := "嘿嘿"
	fmt.Println(s3)

	// var s1 := "100"		// 同一个作用域中不能重复声明同名的变量

}

// 注意事项：
// 1.函数外的每个语句都必须以关键字开始（var、const、func等）
// 2. :=不能使用在函数外
// 3. _多用于占位符，表示忽略值
