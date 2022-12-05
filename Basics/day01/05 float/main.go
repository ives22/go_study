package main

import (
	"fmt"
	"math"
)

// 浮点数

func main() {
	fmt.Println(math.MaxFloat32) // float32 最大值  3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) // float64 最大值  1.7976931348623157e+308
	fmt.Println(math.Pi)         // 获取Pi的值  3.141592653589793

	f1 := 1.2324
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型

	f2 := float32(1.2344)  // 显示声明float32类型
	fmt.Printf("%T\n", f2) // float32

	// f1 = f2  // float32类型不能直接赋值给float64类型
}
