package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"log"
	"os"
	"task4/common"
	"task4/control"
	_ "task4/docs"
)

// 初始化Swagger配置
// @title Gin Web API
// @version 1.0
// @description 个人博客系统 API 文档
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:9001
// @BasePath /
func main() {
	// 设置日志记录到文件。
	f, _ := os.Create("gin.log")
	// 需要同时将日志写入文件和控制台。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 同时输出到控制台（可选）
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	router := gin.New()

	// 内置日志中间件（带颜色）
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// 使用统一异常中间件
	router.Use(common.CustomRecovery())

	//配置swagger.json路径
	router.StaticFile("/swagger.json", "./docs/swagger.json")

	// 添加docs路由
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//控制中心
	control.Action(router)

	//启动服务
	err := router.Run(":9001")
	if err != nil {
		panic(err)
	}

}
