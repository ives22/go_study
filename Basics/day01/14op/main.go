package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a ++ // 单独的语句，不能放在=的右边 赋值  ==> a = a+1
	b -- // 单独的语句，不能放在=的右边 赋值  ==> b = b-1
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b)	// Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)	// 不等于
	fmt.Println(a >= b)	// 大于等于
	fmt.Println(a > b)	// 大于
	fmt.Println(a <= b)	// 小于等于
	fmt.Println(a < b)	// 小于

	// 逻辑运算符
	age := 22
	// 如果年龄大于18岁，并且年龄小于60岁
	if age > 18 && age < 60{
		fmt.Println("苦逼上班的！")
	}else{
		fmt.Println("不用上班！")
	}
	// 如果年龄小于18岁，或者年龄大于60岁
	if age < 18 || age > 60{
		fmt.Println("不用上班！")
	}else{
		fmt.Println("苦逼上班的！")
	}
	// not 取反，原来为真就为假，原来为假就为真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)


	fmt.Println("\n位运算========")
	// 位运算：针对的是二进制
	// 5的二进制表示：101
	// 2的二进制表示：10
	// &：按位与（两位均为1才为1）
	fmt.Println(5 & 2)	// 000
	// |：按位或（两位有1个位1就为1）
	fmt.Println(5 | 2)	/// 111 => 7
	// ^：按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2)	/// 111 => 7
	// <<：将二进制位左移指定位数
	fmt.Println(5 << 1)	// 1010 => 10
	fmt.Println(1 << 10)	// 10000000000 => 1024
	// >>：将二进制位右移指定位数
	fmt.Println(5 >> 1)	// 10 => 2


	// 赋值运算符
	fmt.Println("\n赋值运算符========")
	var x int 
	x = 10
	x += 1	// x = x + 1
	x -= 1	// x = x - 1
	x *= 2	// x = x * 1
	x /= 2	// x = x / 1
	x %= 2	// x = x % 1

	x <<= 2 // x = x << 2
	x &= 2 // x = x & 2
	x |= 2 // x = x | 2
	x ^= 2 // x = x ^ 2
	x >>= 2 // x = x >> 2
	fmt.Println(x)


}
