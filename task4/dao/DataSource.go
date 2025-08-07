package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task4/domain"
	"time"
)

var db *gorm.DB

func init() {
	//建立数据库连接
	var err error
	dsn := "root:admin123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 配置连接池参数（通过底层 *sql.DB）
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("GORM[MySQL]数据库连接成功")
	//初始化表结构
	domain.InitTables(db)
}
