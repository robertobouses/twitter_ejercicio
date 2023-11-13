package http

func getFollowingTweets(c *gin.Context) {
	var user User
	db.Preload("Following").First(&user, 1) // Hardcoded user ID for simplicity
	var tweets []Tweet
	db.Where("user_id IN (?)", user.Following).Find(&tweets)
	c.JSON(200, tweets)
}
