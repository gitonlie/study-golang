package scripts

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

// 创建表
func createBookTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(80) NOT NULL,
		author VARCHAR(50) NOT NULL,
		price DECIMAL(10,2) NOT NULL
	)`
	_, err := db.Exec(schema)
	return err
}

// 创建数据 (Create)
func createBook(db *sqlx.DB, book *Book) (int64, error) {
	res, err := db.NamedExec(
		`INSERT INTO books (title, author, price) VALUES (:title, :author, :price)`, book)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func RunTypeSafeMap() {
	db, err := sqlx.Open("mysql", "root:admin123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}
	defer db.Close()

	//先建表
	createBookTable(db)

	books := []Book{
		{Title: "GO技术手册", Author: "Benjamin J. Evans", Price: 79.00},
		{Title: "GO技术内幕1", Author: "Benjamin J. Evans", Price: 39.00},
		{Title: "GO技术内幕2", Author: "Benjamin J. Evans", Price: 49.00},
		{Title: "GO技术内幕3", Author: "Benjamin J. Evans", Price: 102.00},
		{Title: "GO技术内幕4", Author: "Benjamin J. Evans", Price: 25.00},
	}

	//插入数据
	for _, v := range books {
		idKey, _ := createBook(db, &v)
		fmt.Println("LastInsertId:", idKey)
	}

	//查询
	var book []Book
	sql := "select id, title, author, price from books where price > :price"
	stmt, _ := db.PrepareNamed(sql)
	args := map[string]interface{}{"price": 50}
	if err1 := stmt.Select(&book, args); err1 != nil {
		fmt.Errorf("查询失败: %w", err1)
	}
	fmt.Println(book)

}
