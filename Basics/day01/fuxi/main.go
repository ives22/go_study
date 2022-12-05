package main

import "fmt"

// 课程学习复习
func main() {
	// 切片

	// var s1 []int
	// // s1 = []int{1, 2, 3}
	// fmt.Println(s1)
	// fmt.Println(s1 == nil)

	// // make 初始化，分配内存
	// s2 := make([]bool, 2, 4)
	// fmt.Println(s2)
	
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3 == nil)


	// s1 := []int{1,2,3}	// [1 2 3]
	// s2 := s1
	// var s3 = make([]int, 3, 3)
	// copy(s3, s1)
	// fmt.Println(s2)	// [1 2 3]
	// s2[1] = 100
	// fmt.Println(s1)	//[1 100 3]
	// fmt.Println(s2)	// [1 100 3]
	// fmt.Println(s3)	// []

	// var s1 []int	// nil 必须初始化才能使用
	// // s1 = make([]int, 1, 2)
	// // s1[0] = 100
	// // fmt.Println(s1)
	// s1 = append(s1, 1)	// append会自动初始化
	// fmt.Println(s1)

	// 指针
	// Go里面的指针只能读,不能修改指针变量指向的地址
	// addr := "北京"
	// addrP := &addr
	// fmt.Println(addrP)
	// fmt.Printf("%T\n", addrP)
	// addrR := *addrP
	// fmt.Println(addrR)	

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["lixiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"])	// 如果key不存在返回的是value对应类型的零值
	score, ok := m1["ji"]
	if !ok{
		fmt.Println("不存在这个人")
	}else{
		fmt.Println("这个人的分时是：", score)
	}
	delete(m1, "lixiang")
	fmt.Println(m1)

}


