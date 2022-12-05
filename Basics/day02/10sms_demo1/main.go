package main

import (
	"fmt"
	"os"
)

var (
	allStudent = make(map[int64]*student, 50)
)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号: %d  姓名: %v\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的名字: ")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	value, ok := allStudent[id]
	if !ok {
		allStudent[id] = newStu
	} else {
		fmt.Printf("学号为 %v 的学生已存在!!!\n", value.id)
	}
}

func deleteStudent() {
	var (
		deleteID int64
	)
	fmt.Print("请输入需要删除学生的学号: ")
	fmt.Scanln(&deleteID)
	_, ok := allStudent[deleteID]
	if !ok {
		fmt.Printf("学号为 %v 的学生不存在!!!\n", deleteID)
	} else {
		delete(allStudent, deleteID)
	}
}

func main() {
	for {
		fmt.Println("欢迎访问学生管理系统")
		fmt.Println(`
	1. 查看所有学生
	2. 新增学生
	3. 删除学生
	4. 退出
	`)
		var (
			choice int64
		)
		fmt.Print("你想要干啥:")
		fmt.Scanln(&choice)
		fmt.Println("你选择了:", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚~~~")
		}
	}
}
