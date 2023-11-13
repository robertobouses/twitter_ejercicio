package http

import "github.com/gin-gonic/gin"

func getOwnTweets(c *gin.Context) {
	var tweets []Tweet
	db.Where("user_id = ?", 1).Find(&tweets) // Hardcoded user ID for simplicity
	c.JSON(200, tweets)
}
