// repository/repository.go
package repository

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (r *Repository) CreateAccount(user entity.User) (entity.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
