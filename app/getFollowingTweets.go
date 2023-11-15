package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) GetFollowingTweets(id string) ([]entity.Tweet, error) {
	result, err := s.repo.GetFollowingTweets(id)
	if err != nil {
		return []entity.Tweet{}, err
	}

	return result, nil
}
