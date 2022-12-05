package main

import "fmt"

// if 条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	}else{
		fmt.Println("未成年，不允许进入, 回家写暑假作业去")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")	
	}else if age > 18 {
		fmt.Println("青年")
	}else {
		fmt.Println("好好学习")
	}

	// 特殊写法	（好处，减少内存的占用）
	// 作用域， age1 变量此时只在if条件判断语句中生效
	if age1 := 19; age1 > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	}else{
		fmt.Println("未成年，不允许进入，回家写暑假作业去")
	}

	// 判断sst中的中文汉字的个数 
	stt := "hello 沙河小王子"
	fmt.Println(stt)
}
