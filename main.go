package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()

	initialMigration()

	r.POST("/createAccount", createAccount)
	r.PUT("/configurePassword/:id", configurePassword)
	r.POST("/publishTweet", publishTweet)
	r.GET("/getOwnTweets", getOwnTweets)
	r.GET("/getFollowingTweets", getFollowingTweets)

	r.Run(":8000")
}
