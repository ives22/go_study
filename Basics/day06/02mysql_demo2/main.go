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

// queryMultiRowDemo 多行记录查询
func queryMultiRowDemo(n int) {
	// 1. SQL语句
	sqlStr := `select id, name, age from user where id > ?;`
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 3. 一定要关闭这个rows !!!
	defer rows.Close()
	// 4. 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// insertRowDemo 插入数据
func insertRowDemo() {
	// 1. 写SQL语句
	sqlStr := `insert into user(name, age) values(?, ?);`
	// 2. Exec
	ret, err := db.Exec(sqlStr, "小花", 28)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// updateRowDemo 更新操作
func updateRowDemo(newAge, id int) {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRowDemo() {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// prepareInsertDemo 预处理方式插入多条数据
func prepareInsertDemo() {
	sqlStr := `insert into user(name, age) values(?,?)`
	stmt, err := db.Prepare(sqlStr) // 把SQL语句先发给MySQL预处理以下
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿到stmt去执行一些操作
	var m = map[string]int{ // 造一些测试数据
		"小王": 20,
		"小张": 25,
		"小何": 30,
		"小刘": 30,
	}
	for k, v := range m {
		_, err := stmt.Exec(k, v) // 后续只需要传值
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
}

// prepareQueryDemo 预处理查询
func prepareQueryDemo() {
	sqlStr := `select id, name, age from user where id > ?;`
	stmt, err := db.Prepare(sqlStr) // 把SQL语句发给MySQL预处理一下
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	// 单条记录查询
	// queryRowDemo(2) // u1:main.user{id:2, name:"小黑", age:20}

	// 多行记录查询
	// queryMultiRowDemo(0)

	// 插入数据
	// insertRowDemo()

	// 更新数据
	// updateRowDemo(20, 3) // 更新id为3的数据年龄为20

	// 删除数据
	// deleteRowDemo()

	// 预处理插入数据：
	// prepareInsertDemo()

	// 预处理查询数据
	prepareQueryDemo()
}
