package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (h *Http) PublishTweet(c *gin.Context) {
	var tweet entity.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.PublishTweet(tweet)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
