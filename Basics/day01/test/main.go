package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
goroutine 和channel 练习题：
使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
    1. 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
    2. 开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
    3. 主 goroutine 从resultChan取出结果并打印到终端输出
*/

type Job struct {
	Value int64
}

type Result struct {
	Job *Job
	Sum int64
}

var jobChan = make(chan *Job, 50)
var resultChan = make(chan *Result, 50)
var wg sync.WaitGroup

func createInt(jobCh chan<- *Job) {
	// 循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &Job{
			Value: x,
		}
		jobCh <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func readCount(ch1 <-chan *Job, resultChan chan<- *Result) {
	// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for {
		job := <-ch1
		sum := int64(0)
		n := job.Value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			Job: job,
			Sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go createInt(jobChan)
	// 开启24个 goroutine
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go readCount(jobChan, resultChan)
	}
	// 主 goroutine 从resultChan取出结果并打印到终端输出
	for x := range resultChan {
		fmt.Printf("value: %d sum: %d \n", x.Job.Value, x.Sum)
	}
	wg.Wait()
}
