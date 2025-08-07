package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

// InitTables 初始化表结构
func InitTables(db *gorm.DB) {
	creatTable(db, &User{})
	creatTable(db, &Post{})
	creatTable(db, &Comment{})
}

func creatTable(db *gorm.DB, dst interface{}) {
	err := db.AutoMigrate(dst)
	if err != nil {
		panic(err)
	}
}
