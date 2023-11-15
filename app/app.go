package app

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
	"github.com/robertobouses/twitter_ejercicio/repository"
)

type APP interface {
	ConfigurePassword(id string, user entity.User) (entity.User, error)
	CreateAccount(user entity.User) (entity.User, error)
	GetAccountId(id string) (entity.User, error)
	GetAllTweets() ([]entity.Tweet, error)
	GetFollowingTweets(id string) ([]entity.Tweet, error)
	GetOwnTweets() ([]entity.Tweet, error)
	PostIdFollow(currentUser, userToFollow entity.User) (entity.User, error)
	PostIdUnfollow(currentUser, userToUnfollow entity.User) (entity.User, error)
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
