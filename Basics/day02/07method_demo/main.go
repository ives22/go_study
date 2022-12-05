package main

import "fmt"

// 方法

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用方法具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
	// fmt.Println("汪汪汪~")
}

// 使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

// 使用指针接收者：传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱！")
}

func main() {
	d1 := newDog("黄狗")
	d1.wang()

	p1 := newPerson("小白", 18)
	fmt.Println(p1.age) // 18
	p1.guonian()
	fmt.Println(p1.age) // 18
	p1.zhenguonian()
	fmt.Println(p1.age) // 19
	p1.dream()

}
