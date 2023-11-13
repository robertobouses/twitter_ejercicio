package app

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (s *Service) PublishTweet(tweet entity.Tweet) (entity.Tweet, error) {

	result, err := s.repo.CreateTweet(tweet)
	if err != nil {
		return entity.Tweet{}, err
	}

	return result, nil
}
