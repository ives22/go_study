package main

import "fmt"

// 结构体是值类型
type person struct {
	name, gender string
}

// go 语言中函数传参数永远传的是拷贝
func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女"  // 根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "女" // 语法糖，自动根据指针找到对应的变量
}

func main() {
	var p person
	p.name = "小白"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男

	f2(&p)                // 0x1241ac4
	fmt.Println(p.gender) // 女

	// 1. 结构体指针1
	var p2 = new(person)
	p2.name = "理想"
	p2.gender = "保密"
	fmt.Printf("%T\n", p2) // *main.person
	fmt.Println(p2)
	fmt.Printf("%p\n", p2)  // p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // 求p2的内存地址

	// 2. 结构体指针2
	// 2.1 key-value初始化
	var p3 = person{
		name:   "小",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	// 2.2 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}
