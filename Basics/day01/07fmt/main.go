package main

import "fmt"

// fmt

func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	// 查看变量的值
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello!"
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字符串: %v\n", s)
	fmt.Printf("字符串: %#v\n", s)

	fmt.Print("沙河有沙")
	fmt.Print("又有河")
	fmt.Println("=========")
	fmt.Println("沙河有沙")
	fmt.Println("又有河")

	// Pringf(”格式化字符串“，值)
	// %T：查看类型
	// %d：十进制数
	// %b：二进制数
	// %o：八进制数
	// %x：十六进制数，使用a-f
	// %X：十六进制数，使用A-F
	// %c：字符
	// %s：字符串
	// %p：指针
	// %v：值
	// %#v：值的Go语言表示
	// %f：浮点数
	// %t：布尔值
	// %%：打印百分比数字

	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	fmt.Printf("%v\n", m1)  // map[理想:100]
	fmt.Printf("%#v\n", m1) // map[string]int{"理想":100}

	printBaifenbi(80) // 80%

	// fmt.Printf("%s\n", 100)	// %!s(int=100)
	fmt.Printf("%v\n", 100)

	// 整数->字符
	fmt.Printf("%q\n", 65) // 'A'
	// 浮点数和复数
	fmt.Printf("%b\n", 3.14159265354697)

	// 字符串
	fmt.Printf("%q\n", "李想有理想") // "李想有理想"

	// 获取输入
	var s1 string
	fmt.Scan(&s1)
	fmt.Println("用户输入的内容是：", s1)

	var (
		name string
		age int 
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)

}

func printBaifenbi(num int) {
	fmt.Printf("%d%%\n", num)
}

// func getInput(str string){

// }
