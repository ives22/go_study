package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  int
}

type dog struct{
	name string
}

// 构造函数：约定成俗用new开头
// 构造函数是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}


func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("小白", 9000)
	fmt.Println(p1, p2)
	p2.name = "小黄"
	fmt.Println(p2)
	d1 := newDog("黄")
	fmt.Println(d1)
}
