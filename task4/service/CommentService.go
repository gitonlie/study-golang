package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task4/common"
	"task4/dao"
	"task4/domain"
)

type CommentDTO struct {
	Content string `binding:"required" json:"content"`
	PostID  uint   `binding:"required" json:"postID"`
}

// CreateComment go docs
// @Summary      发表评论
// @Description  发表评论
// @Tags		 评论管理
// @Security 	 BearerAuth
// @Accept       json
// @Produce      json
// @Param		 request body  CommentDTO true "请求体"
// @Success      200  {object}  instance.SuccessResponse
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/createComment [post]
func CreateComment(c *gin.Context) {
	commentDTO := &CommentDTO{}
	err1 := c.ShouldBind(commentDTO)
	if err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	//获取当前用户
	realm := common.GetCurrentUser(c)
	userId := realm.UserId

	comment := &domain.Comment{
		Content: commentDTO.Content,
		PostID:  commentDTO.PostID,
		UserID:  userId,
	}

	if exists := dao.QueryPostExists(comment.PostID); !exists {
		panic(common.Failed("文章不存在"))
	}

	if err2 := dao.CreateComment(comment); err2 != nil {
		panic(err2)
	}

	c.JSON(http.StatusOK, common.OK("发表评论成功"))
}

// CommentList go docs
// @Summary      所有评论列表
// @Description  所有评论列表
// @Tags		 评论管理
// @Security 	 BearerAuth
// @Accept       plain
// @Produce      json
// @Param		 postId query  string true "文章ID" default(10)
// @Success      200  {object}  instance.SuccessResponse{dao=dao.CommentVo}
// @Failure      400  {object}  instance.FailedResponse
// @Failure      401  {object}  instance.UnauthorizedResponse
// @Failure      500  {object}  instance.ErrorResponse
// @Router       /auth/commentList [get]
func CommentList(c *gin.Context) {
	pid := c.Query("postId")
	if pid == "" {
		panic(common.Failed("缺少必填参数"))
	}
	var postId int
	var err1 error
	if postId, err1 = strconv.Atoi(pid); err1 != nil {
		panic(common.Failed(err1.Error()))
	}

	comments := dao.GetCommentByPostId(postId)
	c.JSON(http.StatusOK, common.OK(comments))
}
