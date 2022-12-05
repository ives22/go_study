package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context？

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("疫情")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出？
	exitChan <- true
	wg.Wait()

}
