package repository

import (
	_ "embed"
	"errors"

	"github.com/robertobouses/twitter_ejercicio/entity"
	"gorm.io/gorm"
)

type REPOSITORY interface {
	ConfigurePassword(id string, user entity.User) (entity.User, error)
	CreateAccount(user entity.User) (entity.User, error)
	CreateTweet(tweet entity.Tweet) (entity.Tweet, error)
	GetAccountId(id string) (entity.User, error)
	GetAllTweets() ([]entity.Tweet, error)
	GetFollowingTweets(id string) ([]entity.Tweet, error)
	GetOwnTweets() ([]entity.Tweet, error)
	PostIdFollow(currentUser, userToFollow entity.User) (entity.User, error)
	PostIdUnfollow(currentUser, userToUnfollow entity.User) (entity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &Repository{db: db}, nil
}
