package main

import (
	"database/sql/driver"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局的对象db
var db *sqlx.DB

// initDB 初始化数据库函数
func initDB() (err error) {
	dsn := "root:admin123@tcp(124.71.33.240:3306)/sql_test"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u user) Value()(driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// queryRowDemo 查询单条记录
func queryRowDemo() {
	sqlStr := `select id, name, age from user where id=?`
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("u:%#v\n", u)
}

// queryMultiRowDemo 查询多条记录
func queryMultiRowDemo() {
	sqlStr := `select id, name, age from user where id>?`
	// 定义一个切片存放所有数据
	var userList []user
	err := db.Select(&userList, sqlStr, 1)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)

	// 循环获取每条数据
	for _, v := range userList {
		fmt.Printf("id:%d, name:%s, age:%d\n", v.Id, v.Name, v.Age)
	}
}

// insertRowDemo 插入数据
func insertRowDemo() {
	sqlStr := `insert into user(name, age) values(?, ?)`
	ret, err := db.Exec(sqlStr, "小王", 28)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 获取插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}

// updateRowDemo 更新数据
func updateRowDemo() {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, 40, 8)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 获取受影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// deleteRowDemo 删除数据
func deleteRowDemo() {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, 8)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	// 获取受影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// insertUserDemo 通过DB.NamedExec方法以结构体的方式插入数据
func insertUserDemo() {
	sqlStr := `insert into user(name, age) values(:name,:age)`
	// 早一个结构体类型切片的测试数据
	users := []user{
		{Name: "小花1", Age: 18},
		{Name: "小花2", Age: 20},
		{Name: "小花3", Age: 30},
	}
	// 插入数据
	ret, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("insert for namedexec failed, err:%v\n", err)
	}
	n, err := ret.RowsAffected() // 获取受影响的行，即插入了多少行
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, affected rows:%d\n", n)
}

// insertUserByMapDemo 通过DB.NamedExec方法以map的方式插入数据
func insertUserByMapDemo() {
	sqlStr := `insert into user(name, age) values(:name,:age)`
	// 早一个map类型切片的测试数据
	users := []map[string]interface{}{
		{"name": "小菜1", "age": 20},
		{"name": "小菜2", "age": 30},
		{"name": "小菜3", "age": 40},
	}
	// 插入数据
	ret, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("insert for namedexec failed, err:%v\n", err)
	}
	n, err := ret.RowsAffected() // 获取受影响的行，即插入了多少行
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, affected rows:%d\n", n)
}

// queryByStructDemo 使用DB.NameQuery通过结构体做命名查询
func queryByStructDemo() {
	sqlStr := `select id, name, age from user where id >:id`
	u := user{
		Id: 10,
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err := db.NamedQuery(sqlStr, u)
	fmt.Println(sqlStr)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环处理查询的结果
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

// queryByMapDemo 使用DB.NameQuery通过map做命名查询
func queryByMapDemo() {
	sqlStr := `select id, name, age from user where id >:id`
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"id": 10})
	fmt.Println(sqlStr)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环处理查询的结果
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

// transactionDemo 事务示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()
	sqlStr1 := `update user set age=20 where id=?`
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != 1 {
		fmt.Println("exec sqlStr1 failed")
	}

	sqlStr2 := `update user set age=20 where id=?`
	rs, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != 1 {
		fmt.Println("exec sqlStr2 failed")
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}

	// 查询单条记录
	// queryRowDemo()

	// 查询多条记录
	// queryMultiRowDemo()

	// 插入数据
	// insertRowDemo()

	// 更新数据
	// updateRowDemo()

	// 删除数据
	// deleteRowDemo()

	// 通过DB.NamedExec insert with structs
	// insertUserDemo()

	// 通过DB.NamedExec insert with maps
	// insertUserByMapDemo()

	// 通过DB.NameQuery query with structs
	// queryByStructDemo()

	// 通过DB.NameQuery query with maps
	// queryByMapDemo()

	// 事务示例
	transactionDemo()
}
