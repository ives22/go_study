package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x = 0
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	// lock.Lock()  // 加互斥锁
	rwlock.RLock()  // 加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()  // 释放互斥锁
	rwlock.RUnlock()   // 释放读锁
}

func write() {
	defer wg.Done()
	// lock.Lock()   // 加互斥锁
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock() // 释放互斥锁
	rwlock.Unlock()
}

func main() {
	startTime := time.Now()
	for i:=0; i<10;i++{
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i:=0; i<1000;i++{
		wg.Add(1)
		go read()
	}
	wg.Wait()
	endTime:=time.Now()
	fmt.Println(endTime.Sub(startTime))
}
 