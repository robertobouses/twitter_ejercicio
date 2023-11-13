package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) ConfigurePassword(id string, user entity.User) (entity.User, error) {

	result, err := s.repo.ConfigurePassword(id, user)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
