package app

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (s *Service) PostIdUnfollow(currentUser, userToUnfollow entity.User) (entity.User, error) {

	var updatedFollowing []*entity.User
	for _, u := range currentUser.Following {
		if u.ID != userToUnfollow.ID {
			updatedFollowing = append(updatedFollowing, u)
		}
	}
	currentUser.Following = updatedFollowing

	result, err := s.repo.PostIdUnfollow(currentUser, userToUnfollow)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
