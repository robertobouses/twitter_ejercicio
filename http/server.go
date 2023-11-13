package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/app"
)

type HTTP interface {
	ConfigurePassword(c *gin.Context)
	CreateAccount(c *gin.Context)
	GetFollowingTweets(c *gin.Context)
	GetOwnTweets(c *gin.Context)
	PublishTweet(c *gin.Context)
}

type Http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &Http{
		service: service,
	}
}
