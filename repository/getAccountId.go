package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) GetAccountId(id string) (entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
