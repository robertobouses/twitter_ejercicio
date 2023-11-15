package app

import (
	"errors"

	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (s *Service) PostIdFollow(currentUser, userToFollow entity.User) (entity.User, error) {
	// Llama a la función IsFollowingInDB y maneja los valores devueltos
	isFollowing, err := s.repo.IsFollowingInDB(s.db, currentUser, userToFollow)
	if err != nil {
		return entity.User{}, err
	}

	// Verifica si el usuario ya sigue al otro
	if isFollowing {
		return entity.User{}, errors.New("Ya sigues a este usuario")
	}

	// Agrega el usuario a la lista de seguidos
	currentUser.Following = append(currentUser.Following, &userToFollow)

	// Llama a la función PostIdFollow y maneja los valores devueltos
	result, err := s.repo.PostIdFollow(currentUser, userToFollow)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
