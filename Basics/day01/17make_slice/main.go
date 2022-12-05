package main

import (
	"fmt"
)

// make()函数创造切片
// 如果需要动态的创建一个切片，我们就需要用到内置的make()函数，格式如下：
// make([]T, size, cap)
// T:切片的元素类型；size：切片中元素的数量；cap：切片的容量
func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 切片的赋值
	s3 := []int{1,3,5}
	s4 := s3	// s3和s4 都指向了同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s4, s4)

	// 切片的遍历
	s5 := []int{1,2,3,4,5}
	// 索引遍历
	for i := 0; i < len(s5); i++ {
		fmt.Println(s5[i])
	}
	// for range 遍历
	for k, v := range s5{
		fmt.Println(k, v)
	}
}
