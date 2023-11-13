package http

import (
	entity "command-line-argumentsC:\\Users\\Roberto\\Desktop\\Workspace\\twitter_ejercicio\\entity\\tweet.go"
	"github.com/gin-gonic/gin"
)

func configurePassword(c *gin.Context) {
	id := c.Param("id")
	var user entity.User
	db.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(200, user)
}
