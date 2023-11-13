package app

import "github.com/robertobouses/twitter_ejercicio/entity"

func (s *Service) GetFollowingTweets() ([]entity.Tweet, error) {
	result, err := s.repo.GetFollowingTweets()
	if err != nil {
		return []entity.Tweet{}, err
	}

	return result, nil
}
