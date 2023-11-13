package app

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
	"github.com/robertobouses/twitter_ejercicio/repository"
)

type APP interface {
	ConfigurePassword(id string, user entity.User) (entity.User, error)
	CreateAccount(user entity.User) (entity.User, error)
	GetFollowingTweets() ([]entity.Tweet, error)
	GetOwnTweets() ([]entity.Tweet, error)
	PublishTweet(tweet entity.Tweet) (entity.Tweet, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
