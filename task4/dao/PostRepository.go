package dao

import (
	"task4/domain"
	"time"
)

func CreatePost(post *domain.Post) error {
	err := db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(id uint, post *domain.Post) error {
	err := db.Where("id = ?", id).Updates(post).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int) error {
	err := db.Delete(&domain.Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryPostExists(id any) bool {
	var rows int64
	db.Model(&domain.Post{}).Where("id = ?", id).Count(&rows)
	if rows > 0 {
		return true
	}
	return false
}

func DetailPost(id int) (*DPostVo, error) {
	var post domain.Post
	err := db.Preload("User").First(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	postVo := DPostVo{
		PostID:    post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:     post.Title,
		Content:   post.Content,
		UserVo: UserVo{
			UserID:   post.User.ID,
			Username: post.User.Username,
			Email:    post.User.Email,
		},
	}
	return &postVo, nil
}

func PostList() *[]PostVo {
	var posts []domain.Post
	db.Preload("User").Order("created_at desc").Find(&posts)

	postVos := make([]PostVo, 0)
	for _, post := range posts {
		postVo := PostVo{
			PostID:    post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Title:     post.Title,
			Content:   post.Content,
		}
		postVos = append(postVos, postVo)
	}

	return &postVos
}

type UserVo struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PostVo struct {
	PostID    uint      `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

type DPostVo struct {
	PostID    uint      `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserVo
}
