package repository

import (
	"errors"

	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (r *Repository) PostIdUnfollow(currentUser, userToUnfollow entity.User) (entity.User, error) {

	isFollowing, err := r.IsFollowingInDB(r.db, currentUser, userToUnfollow)
	if err != nil {
		return entity.User{}, err
	}

	if !isFollowing {
		return entity.User{}, errors.New("No sigues a este usuario")
	}

	tx := r.db.Begin()

	if err := tx.Model(&currentUser).Association("Following").Delete(&userToUnfollow); err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	if err := tx.Save(&currentUser).Error; err != nil {
		tx.Rollback()
		return entity.User{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entity.User{}, err
	}

	return currentUser, nil
}
