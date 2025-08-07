package control

import (
	"github.com/gin-gonic/gin"
	"task4/common"
	"task4/service"
)

func Action(router *gin.Engine) {
	router.GET("/hello", func(c *gin.Context) {
		panic("hello world")
	})
	//公共组
	public := router.Group("/public")
	public.POST("/register", service.Register)
	public.POST("/login", service.Login)
	//public.POST("/logout", service.Logout)
	//鉴权组
	auth := router.Group("/auth").Use(common.JWTAuth())

	//文章CURD
	auth.POST("/createPost", service.CreatePost)
	auth.POST("/updatePost", service.UpdatePost)
	auth.POST("/deletePost", service.DeletePost)
	auth.POST("/detailPost", service.DetailPost)
	auth.GET("/postList", service.PostList)

	//评论操作
	auth.POST("/createComment", service.CreateComment)
	auth.GET("/commentList", service.CommentList)
}
