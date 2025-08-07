package dao

import (
	"task4/domain"
)

func QueryUserCount(username string) bool {
	var row int64
	db.Model(&domain.User{}).Where("username = ?", username).Count(&row)
	return row > 0
}

func CreateUser(user *domain.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func QueryUserByName(username string) (*domain.User, error) {
	var user domain.User
	err := db.Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := db.First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
