package main

import (
	"fmt"
	"os"
)

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

func MakeSheepHungry(s Sheep) {
	s.Say()
}

type Sayer interface {
	Say()
}

func MakeHungry(s Sayer) {
	s.Say()
}

type Payer interface {
	Pay(int64)
}

type ZhiFuBao struct {
	// 支付宝
}

type WeChat struct {
	// 微信
}

func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝支付了: %.2f元。\n", float64(amount/100))
}

func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信支付了: %.2f元。\n", float64(amount/100))
}

func CheckOut(obj Payer) {
	obj.Pay(100)
}

// func CheckoutWithZFB(obj *ZhiFuBao) {
// 	obj.Pay(100)
// }

// func CheckoutWithWX(obj *WeChat) {
// 	obj.Pay(100)
// }

func f1()  {
	var fileObj *os.File
	var err error
	fileObj, err =  os.Open("./main.go")
	defer fileObj.Close()
	if err != nil{
		fmt.Println("open file failed, err:", err)
		return
	}
	defer fileObj.Close()
	
}



func main() {
	// c := Cat{}
	// c.Say()
	// d := Dog{}
	// d.Say()
	// s := Sheep{}
	// s.Say()

	// var c1 Cat
	// MakeHungry(c1)

	CheckOut(&ZhiFuBao{})
	CheckOut(&WeChat{})

	var studentInfo = make(map[string]interface{}, 0)
	studentInfo["name"]="小白"
	studentInfo["age"] = 17
	studentInfo["marrid"]=map[int]string{11:"小白"}
	fmt.Println(studentInfo)

	var a interface{}
	a = 100
	// 如何判断a保存的值的具体类型是什么？
	// 类型断言
	// x.(T)
	v1, ok := a.(int8)
	if ok{
		fmt.Println("猜对了，a是int8", v1)
		return
	}else{
		fmt.Println("猜错了")
	}

	switch v2 := a.(type){
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("滚")
	}
}
