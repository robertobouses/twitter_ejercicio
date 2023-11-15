package app

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (s *Service) PostIdFollow(currentUser, userToFollow entity.User) (entity.User, error) {

	currentUser.Following = append(currentUser.Following, &userToFollow)

	result, err := s.repo.PostIdFollow(currentUser, userToFollow)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
