package entity

"github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	Tweets    []Tweet
	Following []*User `gorm:"many2many:user_followers;association_jointable_foreignkey:follower_id"`
}
