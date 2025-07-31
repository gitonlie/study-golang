package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task3/scripts"
)

func main() {
	//连接数据库
	dsn := "root:admin123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("建立数据库连接...")

	//基本CRUD操作
	scripts.RunStudent(db)

	//事务语句
	scripts.RunTransaction(db)

	//Sqlx入门 使用SQL扩展库进行查询
	scripts.RunExtend()

	//Sqlx入门 实现类型安全映射
	scripts.RunTypeSafeMap()

	//进阶gorm
	scripts.CreateUserPostComment(db)
}
