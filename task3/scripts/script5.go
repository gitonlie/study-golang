package scripts

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User 用户
type User struct {
	ID    int
	Name  string
	Posts []Post
	Num   int
}

// Post 文章
type Post struct {
	ID       int
	Title    string
	UserID   int
	Comments []Comment
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	user := User{ID: post.UserID}
	tx.Debug().First(&user)
	tx.Debug().Model(&User{}).Where("id = ?", user.ID).Update("num", user.Num+1)
	fmt.Println("hook create data")
	return nil
}

func (post *Post) BeforeDelete(tx *gorm.DB) (err error) {
	user := User{ID: post.UserID}
	tx.Debug().First(&user)
	tx.Debug().Model(&User{}).Where("id = ?", user.ID).Update("num", user.Num-1)
	fmt.Println("hook delete data")
	return nil
}

// Comment 评论
type Comment struct {
	ID      int
	Content string
	PostID  int
}

func (comment *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	var num int
	tx.Debug().Model(&Comment{}).Select("count(1)").Where("post_id = ?", comment.PostID).First(&num)
	if num == 0 {
		fmt.Println("无评论")
	}
	return nil
}

func CreateUserPostComment(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	user := User{
		Name: "小王",
		Posts: []Post{
			{
				Title: "今日高温",
				Comments: []Comment{
					{
						Content: "今日高温39度",
					},
					{
						Content: "连续高温天气预警",
					},
				},
			}, {
				Title: "六预警齐发",
				Comments: []Comment{
					{
						Content: "中央气象台7月29日18时继续发布暴雨橙色预警",
					},
					{
						Content: "中央气象台7月29日18时发布强对流天气蓝色预警",
					},
				},
			},
		}}

	db.Create(&user)

	var user1 User
	user1.Name = "小王"
	db.Debug().Preload(clause.Associations).Preload("Posts.Comments").First(&user1)
	fmt.Println(user1)

	var post Post
	subQuery := db.Debug().Table("posts p").Select("p.id,count(1) num").Joins("left join comments c on c.post_id = p.id").Group("p.id").Order("num desc").Limit(1)
	db.Debug().Table("posts a").Joins("inner join (?) b on a.id = b.id", subQuery).Preload("Comments").First(&post)
	fmt.Println(post)

	db.Debug().Where("post_id = ?", 1).Delete(&Comment{})

}
