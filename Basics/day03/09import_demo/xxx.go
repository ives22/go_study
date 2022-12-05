package main

import (
	"fmt"

	calc "../08calc"
)

var x = 100
const pi = 3.14

func init(){
	fmt.Println("自动执行!")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(1, 3)
	fmt.Println(ret)
}
