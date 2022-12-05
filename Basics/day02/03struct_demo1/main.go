package main

import "fmt"

// 结构体

// 定义一个 person （人）的结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量per
	var per person
	// 通过字段赋值
	per.name = "小白"
	per.age = 9000
	per.gender = "男"
	per.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(per)
	// 访问变量per的字段
	fmt.Printf("%T\n", per)
	fmt.Println(per.name)

	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type: %T  value: %v\n", p2, p2)

	// 匿名的结构体：多用于一些临时场景
	var s struct{
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type: %T  value: %v\n", s, s)
}