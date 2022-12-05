package main

import "fmt"

// 构造体
type person struct {
	name string
	age  int
}

// 为什么要有构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age: age,
	}
	
}

func main() {
	var p1 = person{
		name: "小白",
		age:  18,
	}
	fmt.Println(p1)

	p2 := person{
		"小白1",
		22,
	}
	fmt.Println(p2)

	p3 := newPerson("小白2", 22)
	fmt.Println(p3)
}
