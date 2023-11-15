package entity

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}

func (Tweet) TableName() string {
	return "tweet"
}

//evito que gorm ponga "tweets" como nombre
