#   1.安装gorm 框架
go get -u gorm.io/gorm

#   2.安装mysql驱动
go get -u gorm.io/driver/mysql

#   3.安装gin 框架
go get -u github.com/gin-gonic/gin

#   4.集成 JWT
go get -u github.com/golang-jwt/jwt/v5

#   5.Swagger 集成
## 安装 swag 命令行工具
go install github.com/swaggo/swag/cmd/swag@latest

## 安装 Gin 的 Swagger 集成包
go get -u github.com/swaggo/gin-swagger

## 初始化Swagger文档
swag init

## 指定端口运行 9001
http://localhost:9001/docs/index.html 可访问swagger文档
