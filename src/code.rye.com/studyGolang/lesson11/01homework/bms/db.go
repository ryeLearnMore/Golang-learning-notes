package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 跟数据库相关的操作
var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:1234@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 设置数据库连接池最大空闲连接数
	db.SetMaxIdleConns(16)
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(100)
	return
}

// 查询数据
func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id, title, price from book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍信息失败！")
		return
	}
	return
}

// day 11
// 查询单个数据
func queryBookByID(id int64) (book Book, err error) {
	// sqlStr := "select id, title, price from book where id = ?;"
	sqlStr :=
		`
	select
		book.id, book.title, book.price, publisher.province, publisher.city
	from
		book
	inner join
		publisher
	on book.publisher_id = publisher.id
	where
		book.id = ?
		`
	err = db.Get(&book, sqlStr, id)
	if err != nil {
		fmt.Println("查询该书籍信息失败！")
		return
	}
	return
}

// 插入数据
func insertBook(title string, price float64) (err error) {
	sqlStr := "insert into book(title, price) values(?, ?);"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入书籍信息失败！")
		return
	}
	return
}

// 删除数据
func deleteBook(id int64) (err error) {
	sqlStr := "delete from book where id = ?;"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除书籍信息失败！")
		return
	}
	return
}

// day 11
// 编辑数据
func editBook(title string, price float64, id int64) (err error) {
	sqlStr := "update book set title = ?, price = ? where id = ?;"
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		fmt.Println("编辑书籍信息失败！")
		return
	}
	return
}
