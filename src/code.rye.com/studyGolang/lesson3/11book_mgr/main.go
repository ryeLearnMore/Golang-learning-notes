package main
import (
	"fmt"
	"os"
)
// 需求
// 使用函数实现一个简单的图书馆里系统
// 每本书有书名，作者，价格，上架信息

// 需求分析
// 0. 定义结构体
type book struct {
	title string
	author string
	price float32
	publish bool
}

// 定义创建新书的构造函数
func newBook(title, author string, price float32, publish bool) *book {
	return &book{
		title: title,
		author: author,
		price: price,
		publish: publish,
	}
}

// 定义一个存放book指针的切片，用来存储所有的书籍
var allBooks = make([]*book, 0, 200)

// 1. 打印菜单
func showMenu()  {
	fmt.Println("欢迎登陆BMS!")
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
}
// 2. 等待用户输入菜单选项
func userInput() *book {
	var (
		title string
		author string
		price float32
		publish bool
	)
	// 3.1 获取用户输入
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名：")
	fmt.Scanln(&title)
	fmt.Print("请输入作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入价格：")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架[true|false]：")
	fmt.Scanln(&publish)
	fmt.Println(title, author, price, publish)
	// 3.2 创建一本新书
	book := newBook(title, author, price, publish)
	return book
}
// 3. 添加书籍的函数
func addBook()  {
	book := userInput()
	// 3.3 把书添加到allBooks切片中
	// 3.3.1 判断新增的书名是否已经存在
	for _, b := range allBooks{
		if b.title == book.title {
			fmt.Printf("《%s》已经存在", book.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Println("添加书籍成功！")
	fmt.Println("----------------------------")
}

// 4. 修改书籍的函数
func updateBook()  {
	book := userInput()
	// 4.1 遍历所有的书籍根据书名找到要修改的那本书，更新信息
	for index, b := range allBooks {
		// 如果遍历到的这本书名正好等于book.title 就更新
		if b.title == book.title {
			allBooks[index] = book
			fmt.Printf("《%s》更新成功！\n", book.title)
			return
		}
	}
	fmt.Printf("《%s》不存在！\n", book.title)
	fmt.Println("----------------------------")
}

// 5. 展示书籍的函数
func showBooks()  {
	if len(allBooks) == 0 {
		fmt.Println("啥也没有！")
		return
	}

	for _, b := range allBooks {
		fmt.Printf("《%s》作者：%s 价格：%.2f 是否上架：%t\n", b.title, b.author, b.price, b.publish)
	}
}

// 6. 退出 os.Exit(0)

func main()  {
	for {
		showMenu()
		// 2. 等待用户输入菜单选项
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addBook()
		case 2:
			updateBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)
		}
	}
}