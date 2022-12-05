package main

import (
	"fmt"
	"time"
)

// time 时间

func formatDemo() {
	// 格式时间，把语言中时间对象，转换成字符串类型的时间

	now := time.Now() // 获取当前时间
	// 2022-11-07
	fmt.Println(now.Format("2006-01-02"))
	// 2022/11/07 20:30:42
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2022/11/07 08:31:40 PM
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	// 2022/11/07 20:32:36.482
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2022-11-07")
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

func f1() {
	now := time.Now() // 获取当前时间
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())


	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 纳秒

	// time.Unix()
	ret := time.Unix(1667823489, 0)
	fmt.Println(ret) // 2022-11-07 20:18:09 +0800 CST
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)

	// now + 24小时   获取明天什么时候
	fmt.Println("时间间隔-------------")
	fmt.Println(now.Add(24 * time.Hour))
	// sub 两个时间相减法
	nextYear, err := time.Parse("2006-01-02", "2022-11-08")
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	d := now.Sub(nextYear)
	d1 := nextYear.Sub(now)
	fmt.Println(d)  // -11h15m32.923751s
	fmt.Println(d1) // 11h15m32.923751s
	fmt.Println("时间间隔-------------")

	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer{
	// 	fmt.Println(t)   // 1秒钟执行一次
	// }

	// sleep
	time.Sleep(100)
	n := 100
	time.Sleep(time.Duration(n) * time.Second)
	time.Sleep(5 * time.Second)
}

// 时区
func f2() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2022-11-08 20:59:50")
	// 按照东八区的时区和格式去解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed, err:", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-11-08 20:59:50", loc)
	if err != nil {
		fmt.Println("parse time failed, err:", err)
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td) // 23h55m9.944108s
}

func main() {

	// f1()

	// 时间格式化
	// formatDemo()

	// f2()
	
	now := time.Now().Hour()
	fmt.Println(now)
	fmt.Printf("%T\n", now)
	currentTime := time.Now()
    m, _ := time.ParseDuration("-1h")
    result := currentTime.Add(m)
    fmt.Println(result.Format("2006010215"))
	// time.ParseDuration("-1h")
	fmt.Println(time.Now().Add(m).Format("2006010215"))

}
