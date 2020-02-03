package main

// 数据库事务的操作

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:1234@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // dsn格式不正确的时候会报错
		return
	}
	err = db.Ping() // 检测用户名和密码是否正确，尝试连接数据库
	if err != nil {
		return
	}

	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

func transactionDemo() {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed: %v\n", err)
		return
	}
	// 2. 执行多个SQL操作
	sqlStr1 := "update user set age = age - 2 where id = 1;"
	sqlStr2 := "update user set age = age + 2 where id = 2;"
	// 执行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL1出错，要回滚。")
		return
	}
	// 执行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行SQL2出错，要回滚。")
		return
	}
	// 上面两步SQL都成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("提交本次事务出错，要回滚。")
		return
	}
	fmt.Println("事务执行成功！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed: %v\n", err)
	}
	fmt.Println("连接数据库成功！")

	transactionDemo()
}
