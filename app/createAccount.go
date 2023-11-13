package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) CreateAccount(user entity.User) (entity.User, error) {

	result, err := s.repo.CreateAccount(user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
