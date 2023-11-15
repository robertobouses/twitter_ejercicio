package repository

import (
	"errors"

	"github.com/robertobouses/twitter_ejercicio/entity"
	"gorm.io/gorm"
)

func (r *Repository) IsFollowingInDB(db *gorm.DB, currentUser, userToFollow entity.User) (bool, error) {
	if currentUser.ID == 0 || userToFollow.ID == 0 {
		return false, errors.New("ID de usuario inválido")
	}

	var count int64
	if err := db.Model(&currentUser).Where("id = ? AND user_following_id = ?", currentUser.ID, userToFollow.ID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
