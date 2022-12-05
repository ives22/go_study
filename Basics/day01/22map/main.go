package main

import "fmt"

// map

func main() {
	var m1 map[string]int
	fmt.Println(m1)	// map[]
	fmt.Println(m1 == nil)	// true 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10)	// 进行初始化，要估算好该map容量，避免在程序运行期间再动态扩容
	m1["理想"] = 900
	m1["小白"] = 18
	fmt.Println(m1)
	fmt.Println(m1["理想"])

	// 定义方式2
	m2 := make(map[string]int)
	fmt.Println(m2)

	// 拿取不存在的值
	fmt.Println(m1["娜扎"])	// 如果不存在这个key拿到对应值类型的零值

	// 约定成俗用ok接收返回的布尔值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	}else{
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1{
		fmt.Println(k, v)
	}

	// 只是遍历key
	for k := range m1{
		fmt.Println(k)
	}

	// 只想遍历value
	for _, v := range m1{
		fmt.Println(v)
	}

	// 删除map中某个键值对，使用delete()函数的格式
	delete(m1, "小白")
	fmt.Println(m1)
	// 删除不存在的
	delete(m1, "沙河")	// 删除不存在的key， 不会报错
}
