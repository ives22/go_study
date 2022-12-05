package main

import (
	"fmt"
	"strconv"
)

// strconv

func f1() {
	// 把数字转换成字符串类型
	i := int32(97)
	ret1 := string(i)
	fmt.Println(ret1)            // "a"
	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Println(ret2)

	// 把字符串转换为int类型
	str := "10000"
	ret3, err := strconv.ParseInt(str, 10, 64) // strconv.ParseInt(需要转的字符串, 进制, 位数)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%T, %v\n", ret3, ret3) // int64, 10000
	// 把字符串转为int类型
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%T, %v\n", retInt, retInt) // int, 10000

	// 把int类型转换为 string类型
	
	retStr := strconv.Itoa(20)
	fmt.Printf("%T, %v\n", retStr, retStr) // string, 20

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _:=strconv.ParseBool(boolStr)
	fmt.Printf("%T, %v\n", boolValue, boolValue) // bool, true

	// 从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _:=strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%T, %v\n", floatValue, floatValue) // float64, 1.234
}

func main() {
	f1()
}
