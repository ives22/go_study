package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// SQL注入

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

// initDB 初始化数据库函数
func initDB() (err error) {
	dsn := "root:admin123@tcp(124.71.33.240:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// sql注入示例
func sqlInjectDemo(name string) {
	// 自己拼接SQL语句的字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	sqlInjectDemo("xxx' or 1=1#")
	sqlInjectDemo("xxx' union select * from user #")
	sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}
