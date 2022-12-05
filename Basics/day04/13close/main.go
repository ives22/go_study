package main

import "fmt"

// 关闭通道
func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }

	x, ok := <-ch1
	fmt.Println(x, ok) // 10 true
	x, ok = <-ch1
	fmt.Println(x, ok) // 20 true
	x, ok = <-ch1      // 如果没有值了，还是可以取到，为对应类型的0值，ok为false
	fmt.Println(x, ok) // 0 false
}
