package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) GetAllTweets() ([]entity.Tweet, error) {
	var tweets []entity.Tweet
	if err := r.db.Find(&tweets).Error; err != nil {
		return []entity.Tweet{}, err
	}

	return tweets, nil
}
