package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) GetAccountId(id string) (entity.User, error) {

	result, err := s.repo.GetAccountId(id)
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}
