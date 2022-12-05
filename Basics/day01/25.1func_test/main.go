package main

import "fmt"

// 函数的练习

func f1(x int, y int) (ret int) {
	ret = x + y
	return
}

func f2(x, y int)int {
	return x + y
}

func f3(x int, y...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

func f4(x string, y...int) {
	fmt.Println(x, y)
	fmt.Printf("%T\n", y)
}

func f5(x,y int)(sum, sub int){
	sum = x +y
	sub = x -y
	return 
}


func main() {
	r := f1(1, 3)
	fmt.Println(r)

	r2 := f2(1,4)
	fmt.Println(r2)

	f3(1, 2, 3, 4, 5, 6)
	f4("小白", 3,4,5,6,6)
	sum, sub := f5(5, 8)
	fmt.Println(sum, sub)
}

