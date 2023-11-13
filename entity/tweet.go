package entity

type Tweet struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}
