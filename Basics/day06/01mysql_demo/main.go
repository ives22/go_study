package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL示例

func main() {
	// 连接数据库
	// 用户名:密码@tcp(ip:端口)/数据库
	dsn := "root:admin123@tcp(124.71.33.240:3306)/testdb1"

	// 连接数据库
	db, err := sql.Open("mysql", dsn)  // 不会校验用户名和密码是否正确
	if err != nil {  // dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	
	err = db.Ping()  // 尝试与数据库建立连接
	if err != nil {
		fmt.Printf("connection %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功!")
}
