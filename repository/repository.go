package repository

import (
	_ "embed"

	"github.com/jinzhu/gorm"
	"github.com/robertobouses/twitter_ejercicio/entity"
)

type REPOSITORY interface {
	ConfigurePassword(id string, user entity.User) (entity.User, error)
	CreateAccount(user entity.User) (entity.User, error)
	CreateTweet(tweet entity.Tweet) (entity.Tweet, error)
	GetFollowingTweets() ([]entity.Tweet, error)
	GetOwnTweets() ([]entity.Tweet, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) REPOSITORY {
	return &Repository{db: db}
}
