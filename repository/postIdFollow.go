package repository

import (
	"errors"

	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (r *Repository) PostIdFollow(currentUser, userToFollow entity.User) (entity.User, error) {

	isFollowing, err := r.IsFollowingInDB(r.db, currentUser, userToFollow)
	if err != nil {
		return entity.User{}, err
	}

	if isFollowing {
		return entity.User{}, errors.New("Ya sigues a este usuario")
	}

	tx := r.db.Begin()

	if err := tx.Model(&currentUser).Association("Following").Append(&userToFollow); err != nil {
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
