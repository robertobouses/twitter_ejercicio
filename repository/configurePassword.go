package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) ConfigurePassword(id string, user entity.User) (entity.User, error) {
	if err := r.db.First(&user, id).Error; err != nil {
		return entity.User{}, err
	}

	if err := r.db.Save(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
