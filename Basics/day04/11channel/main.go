package main

import (
	"fmt"
	"sync"
)

// channel 是一个引用类型，需要通过make初始化才能使用，需要开辟空间
var ch1 chan int // 需要指定同道中元素的类型

var wg sync.WaitGroup // 声明全局等待组变量

// 无缓冲的通道
func noBufChannel() {
	wg.Add(1) // 登记1个goroutine
	ch1 = make(chan int)
	go func() {
		defer wg.Done() // 告知当前goroutine完成
		x := <-ch1      // 从通道中取值，并赋值给x
		fmt.Println("从通道中ch1中获取了值:", x)
	}()
	ch1 <- 10 // 往通道ch1中存入10
	fmt.Println("往通道ch1中存入了10")
	wg.Wait() // 阻塞等待登记的goroutine完成
}

func bufChannel() {
	ch1 = make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch1 <- 10
	fmt.Println("往通道ch1中存入了10")
	x := <-ch1
	fmt.Println("从通道中ch1中获取了值:", x)
}

func main() {
	/*
		fmt.Println(ch1)         // <nil> 不初始化则为nil
		ch1 = make(chan int)     // 通道的初始化, 不带缓冲区的通道的初始化
		ch1 = make(chan int, 16) // 带缓冲区的通道的初始化
		fmt.Println(ch1)         // 0xc0000260c0  内存地址

		// 通道的操作 <-
		// 1. 发送：ch1 <- 10
		// 2. 接收：x := <- ch1
		// 3. 关闭：close()
	*/

	noBufChannel()
	bufChannel()
}
