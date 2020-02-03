package main

// 数据库事务的操作

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:1234@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sqlx.Connect("mysql", dsn)
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
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed: %v\n", err)
		return
	}
	sqlStr1 := "select id, name, age from user where id = 1;"
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := "select id, name, age from user;"
	db.Select(&userList, sqlStr2)
	fmt.Printf("userList: %#v\n", userList)
}
