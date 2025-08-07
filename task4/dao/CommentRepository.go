package dao

import (
	"task4/domain"
	"time"
)

func CreateComment(comment *domain.Comment) error {
	err := db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCommentByPostId(postId int) *[]CommentVo {
	var comments []domain.Comment
	db.Order("created_at desc").Where("post_id = ?", postId).Find(&comments)
	commentVos := make([]CommentVo, 0)
	for _, comment := range comments {
		commentVos = append(commentVos, CommentVo{
			CommentID: comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})
	}
	return &commentVos
}

type CommentVo struct {
	CommentID uint      `json:"commentId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
