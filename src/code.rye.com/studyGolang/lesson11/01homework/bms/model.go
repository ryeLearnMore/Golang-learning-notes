package main

// 专门用来定义与数据对应的结构体

type Book struct {
	ID    int64   `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
	Publisher
}

type Publisher struct {
	ID       int64  `db:"id"`
	Province string `db:"province"`
	City     string `db:"city"`
}
