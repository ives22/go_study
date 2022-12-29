package main

import (
	"fmt"
	"gin_demo/dao"
	"gin_demo/models"
	"gin_demo/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err)
	}

	r := routers.SetupRouter()

	r.Run(":8080")

}
