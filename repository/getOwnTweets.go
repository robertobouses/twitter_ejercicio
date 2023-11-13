package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) GetOwnTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	if err := r.db.Where("user_id = ?", 1).Find(&tweets).Error; err != nil {
		return []entity.Tweet{}, err
	}

	return tweets, nil
}
