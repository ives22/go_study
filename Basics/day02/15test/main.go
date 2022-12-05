package main

import (
	"encoding/json"
	"fmt"
)

type dongWu struct {
	name string
}

type dog struct {
	feet uint8
	dongWu
}

func (d dongWu) move() {
	fmt.Printf("%s会动\n", d.name)
}

func (d dog) wang() {
	fmt.Printf("%s叫: 汪汪汪~~~\n", d.name)
}

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	d1 := dog{feet: 3, dongWu: dongWu{name: "白狗"}}
	d1.move()
	d1.wang()


	p1 := person{
		Name: "小白",
		Age: 18,
	}
	value, err := json.Marshal(p1)
	if err != nil{
		fmt.Println(err)
		return
		
	}
	fmt.Printf("%v\n", string(value))

	// 反序列化
	st := `{"name":"小白","age":18}`
	var p2 person
	json.Unmarshal([]byte(st), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Name)
}
