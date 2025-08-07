package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"task4/common"
	"task4/dao"
	"task4/domain"
)

type PostDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UPostDTO struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreatePost go docs
// @Summary      发表文章
// @Description  发表文章
// @Tags		 文章管理
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param		 request body  PostDTO true "请求体"
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/createPost [post]
func CreatePost(c *gin.Context) {
	postDTO := &PostDTO{}
	err1 := c.ShouldBind(postDTO)
	if err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	//获取当前用户
	realm := common.GetCurrentUser(c)
	userId := realm.UserId

	post := &domain.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		UserID:  userId,
	}
	if err2 := dao.CreatePost(post); err2 != nil {
		panic(err2)
	}

	c.JSON(http.StatusOK, common.OK(post.ID))
}

// UpdatePost go docs
// @Summary      修改文章
// @Description  修改文章
// @Tags		 文章管理
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param		 request body  UPostDTO true "请求体"
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/updatePost [post]
func UpdatePost(c *gin.Context) {
	postDTO := &UPostDTO{}
	err1 := c.ShouldBind(postDTO)
	if err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	//获取当前用户
	realm := common.GetCurrentUser(c)
	userId := realm.UserId

	post := &domain.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		UserID:  userId,
	}
	if err2 := dao.UpdatePost(postDTO.ID, post); err2 != nil {
		panic(err2)
	}
	c.JSON(http.StatusOK, common.OK("更新成功"))
}

// DeletePost go docs
// @Summary      删除文章
// @Description  删除文章
// @Tags		 文章管理
// @Security 	 BearerAuth
// @Accept       plain
// @Produce      json
// @Param		 id query  string true "请求体" default(10)
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/deletePost [post]
func DeletePost(c *gin.Context) {
	pid := c.Query("id")
	if pid == "" {
		panic(common.Failed("缺少必填参数"))
	}
	var postId int
	var err1 error
	if postId, err1 = strconv.Atoi(pid); err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	if ok := dao.QueryPostExists(postId); !ok {
		panic(common.Failed("文章不存在"))
	}

	if err2 := dao.DeletePost(postId); err2 != nil {
		panic(err2)
	}
	c.JSON(http.StatusOK, common.OK("删除成功"))
}

// DetailPost go docs
// @Summary      查询文章详情
// @Description  查询文章详情
// @Tags		 文章管理
// @Security 	 BearerAuth
// @Accept       plain
// @Produce      json
// @Param		 id query  string true "请求体" default(10)
// @Success      200  {object}  instance.SuccessResponse{data=dao.DPostVo}
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/detailPost [post]
func DetailPost(c *gin.Context) {
	rid := c.Query("id")
	if rid == "" {
		panic(common.Failed("缺少必填参数"))
	}
	var postId int
	var err1 error
	if postId, err1 = strconv.Atoi(rid); err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	postVo, err2 := dao.DetailPost(postId)
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			panic(common.Failed("文章不存在"))
		}
		panic(err2)
	}
	c.JSON(http.StatusOK, common.OK(postVo))
}

// PostList go docs
// @Summary      查询文章列表
// @Description  查询文章列表
// @Tags		 文章管理
// @Security 	 BearerAuth
// @Accept       plain
// @Produce      json
// @Success      200  {object}  instance.SuccessResponse{data=dao.PostVo}
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/postList [get]
func PostList(c *gin.Context) {
	posts := dao.PostList()
	c.JSON(http.StatusOK, common.OK(posts))
}
