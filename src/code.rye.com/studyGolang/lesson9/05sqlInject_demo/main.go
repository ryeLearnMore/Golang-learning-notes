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

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed: %v\n", err)
		return
	}

	// sql注入的几种示例
	sqlInjectDemo("张迪")

	sqlInjectDemo("xxx' or 1=1 #")

	sqlInjectDemo("xxx' union select * from user #")
	sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}
