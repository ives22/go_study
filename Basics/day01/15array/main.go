package main

import (
	"fmt"
)

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度是数据类型的一部分
// var 数组变量名 [元素数量]T

func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1:%T a2: %T\n", a1, a2) // a1:[3]bool a2: [4]bool

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串：""）
	fmt.Println(a1, a2)
	// 1. 初始化方式1  使用指定的值来完成初始化
	a1 = [3]bool{true, true, true}
	fmt.Println(a1) // [true true true]
	// 2. 初始化方式2
	// 根据初始值自动推断数组的长度是多少
	a100 := [...]int{0, 1, 2, 3, 4, 4, 6}
	fmt.Println(a100) // [0 1 2 3 4 4 6]
	// 3. 初始化方式3
	// 根据索引进行初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3) // [1 0 0 0 2]


	fmt.Println("\n数组的遍历=======")
	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}	// 索引：0~2 citys[0], citys[1], citys[2]
	// 1. 方式1，根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2. 方式2，fo range遍历
	for i, v := range citys{
		fmt.Println(i, v)
	}


	fmt.Println("\n多维数组=======")
	// 多维数组
	// [[1,2] [3,4] [5,6]]
	var a11 [3][2]int
	a11 = [3][2]int {
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println(a11)

	fmt.Println("\n多维数组的遍历=======")
	// 多为数组的遍历
	for _, v1 := range a11{
		fmt.Println(v1)
		for _, v2 := range v1{
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1	// [1 2 3]
	b2[0] = 100	// b2: [100 2 3]
	fmt.Println(b1, b2)	// [1 2 3] [100 2 3]

	fmt.Println("\n数组练习题=======")
	// 数组练习题
	// 1.求数组[1,3,5,7,8]所有元素的和
	ss1 := [...]int{1,3,5,7,8}
	sum := 0
	for _, v := range ss1{
		sum += v;
	}
	fmt.Println(sum)

	// 2.找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找到和为8的两个元素的下标分别为(0,3)和(1,2)
	ss2 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(ss2); i++ {
		for j := i+1; j < len(ss2); j++ {
			if ss2[i] + ss2[j] == 8 {
				fmt.Printf("(%d %d)\n", i, j)
			}
		}
	}
}