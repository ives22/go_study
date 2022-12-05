package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL示例

// 定义一个全局的对象db
var db *sql.DB

// initDB 初始化数据库函数
func initDB() (err error) {
	// 连接数据库
	// 用户名:密码@tcp(ip:端口)/数据库
	dsn := "root:admin123@tcp(124.71.33.240:3306)/sql_test"

	// 连接数据库
	// 注意！！！这里不要使用:=，上面是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // dsn格式不正确的时候会报错
		return err
	}

	// 尝试与数据库建立连接 （校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}

	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 单条记录查询
func queryRowDemo(id int) {
	// 声明一个user对象
	var u1 user
	// 1. 写一个查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?`
	// 2. 执行并获取结果
	// 从连接池中拿一个连接出来去数据库查询单条记录， 必须对rowObj对象调用Scan方法，因为该方法会释放数据库连接
	err := db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	// 3. 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}

	sqlStr := `select id,name,age from user where id=?`
	// 这里模拟查询11次，超出设置的最大连接数
	for i := 0; i < 11; i++ {
		db.QueryRow(sqlStr, 1)
		fmt.Printf("已经进行了第%d次查询\n", i)
	}
}
