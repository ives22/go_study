package main

import (
	"fmt"
	"sync"
)

// 锁  互斥锁

var x = 0
var wg sync.WaitGroup
// 声明一个全局的锁
var lock sync.Mutex	

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()	// 加锁
		x = x + 1
		lock.Unlock()  // 释放锁
	}
	wg.Done()

}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
