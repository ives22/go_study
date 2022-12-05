package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数
func main() {
	// // 创建一个标志位参数
	// name := flag.String("name", "小白", "请输入名字")
	// age := flag.Int("age", 9000, "请输入真实年龄")
	// married := flag.Bool("married", false, "结婚了么")
	// delay := flag.Duration("delay", time.Second, "结婚多久了")

	// // 此时name、age、married、delay均为对应类型的指针。
	// // 使用flag
	// flag.Parse() // 调用flag.Parse()来对命令行参数进行解析。
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*delay)
	// fmt.Printf("%T\n", *delay)

	var (
		name    string
		age     int
		married bool
		delay   time.Duration
	)
	flag.StringVar(&name, "name", "小白", "请输入名字")
	flag.IntVar(&age, "age", 9000, "请输入真实年龄")
	flag.BoolVar(&married, "married", false, "结婚了么")
	flag.DurationVar(&delay, "delay", time.Second, "结婚多久了")

	// 使用flag.Parse() 解析命令行参数
	flag.Parse()
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(delay)
	fmt.Printf("%T\n", delay)

	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
}
