package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	// "math/rand" 随机数
	rand.Seed(time.Now().UnixNano()) // 添加随机数的种子， 保证每次执行的时候都不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // 返回一个int64类型的随机数
		r2 := rand.Intn(10) // 0<= x < 10
		// fmt.Println(r1, r2)
		fmt.Println(0-r1, 0-r2) // 拿到随机的负数
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	// f()

	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个goroutne，添加一个计数器
		go f1(i)
	}

	// ？ 如何知道这个10个goroutine都结束了
	wg.Wait() // 等待wg的计数器减为0

}
