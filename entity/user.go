package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Tweets    []Tweet
	Following []*User `gorm:"many2many:user_followers;association_jointable_foreignkey:follower_id"`
}

func (User) TableName() string {
	return "account"
}
