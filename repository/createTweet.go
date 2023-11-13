package repository

import "github.com/robertobouses/twitter_ejercicio/entity"

func (r *Repository) CreateTweet(tweet entity.Tweet) (entity.Tweet, error) {

	if err := r.db.Create(&tweet).Error; err != nil {
		return entity.Tweet{}, err
	}

	return tweet, nil
}
