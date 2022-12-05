package main

import "fmt"

// 引出接口的实例
// 接口是一种类型，是一种特殊的类型，它规定了变量的有哪些方法

// 接口的定义
/*
type 接口名 interface {
	方法名1(参数1, 参数2...)(返回值1, 返回值2...)
	方法名2(参数1, 参数2...)(返回值1, 返回值2...)
	...
}

用来给变量\参数\返回值等设置类型.
*/

// 接口的实现
/*
一个变量如果实现了接口中规定的所有的方法，那么这个变量就是实现了这个接口，可以称为为这个接口类型的变量。
*/


// 定义一个能叫的类型
type speaker interface{
	speak()  // 只要实现了speak方法的变量都是speaker类型, 方法签名
}

type cat struct{}

type dog struct{}

type person struct{}


func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak(){
	fmt.Println("啊啊啊~")
}

func da(x speaker){
	// 接收一个参数，传进来什么，我就打什么
	x.speak()  // 挨打了就要叫
}

// 在编程中会遇到以下场景：
// 我不关心这一个变量是什么类型，我只关心能调用它的什么方法。

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

	var ss speaker   // 定义一个接口类型：speaker 的变量：ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
