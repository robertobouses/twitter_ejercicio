package repository

import (
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (r *Repository) GetFollowingTweets(id string) ([]entity.Tweet, error) {
	var user entity.User
	if err := r.db.Preload("Following").First(&user, id).Error; err != nil {
		return []entity.Tweet{}, err
	}

	var followingUserIDs []uint
	for _, followingUser := range user.Following {
		followingUserIDs = append(followingUserIDs, followingUser.ID)
	}

	var tweets []entity.Tweet
	if err := r.db.Where("user_id IN (?)", followingUserIDs).Find(&tweets).Error; err != nil {
		return []entity.Tweet{}, err
	}

	return tweets, nil
}
