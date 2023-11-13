package http

import "github.com/gin-gonic/gin"

func publishTweet(c *gin.Context) {
	var tweet Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tweet.UserID = 1 // Hardcoded user ID for simplicity
	db.Create(&tweet)
	c.JSON(200, tweet)
}
