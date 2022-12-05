package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go实现MySQL事务示例

// 定义一个全局的对象db
var db *sql.DB

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

func transactionDemo() {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
	}
	// 执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id=1`
	// sqlStr2 := `update user set age=age+2 where id=2`
	sqlStr2 := `update xxx set age=age+2 where id=2`
	// 执行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	// 执行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	// 上面两步SQL都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Printf("事务执行成功\n")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}

	transactionDemo()
}
