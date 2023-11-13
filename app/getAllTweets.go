package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) GetAllTweets() ([]entity.Tweet, error) {

	result, err := s.repo.GetAllTweets()
	if err != nil {
		return []entity.Tweet{}, err
	}

	return result, nil
}
