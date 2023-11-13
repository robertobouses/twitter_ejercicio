package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) GetFollowingTweets() ([]entity.Tweet, error) {
	var user entity.User
	if err := r.db.Preload("Following").First(&user, 1).Error; err != nil {
		return []entity.Tweet{}, err
	}

	var tweets []entity.Tweet
	if err := r.db.Where("user_id IN (?)", user.Following).Find(&tweets).Error; err != nil {
		return []entity.Tweet{}, err
	}

	return tweets, nil
}
