package main

import (
	"fmt"
	"sync"
)

// Channel 练习
// 1. 创建一个goroutne向通道ch1中存100个值
// 2. 创建一个goroutne从通道ch1中读取值，并计算平方，存入到ch2中
// 3. 打印最终计算结果

var wg sync.WaitGroup
var once sync.Once

// func f1(ch1 chan int) {
func f1(ch1 chan<- int) {   // 单向通道，在chan<- 表示只能存值
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// func f2(ch1 chan int, ch2 chan int) { // ch1 和 ch2 即能存值也能取值
func f2(ch1 <-chan int, ch2 chan<- int) { // 单向通道，ch1只能取值，ch2只能存值
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {close(ch2)})	// 确保某个操作只是执行一次
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for x := range b {
		fmt.Println(x)
	}
}
