package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) GetOwnTweets() ([]entity.Tweet, error) {
	result, err := s.repo.GetOwnTweets()
	if err != nil {
		return []entity.Tweet{}, err
	}

	return result, nil
}
