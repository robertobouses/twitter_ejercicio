package http

import "github.com/gin-gonic/gin"

func (h *Http) GetAllTweets(c *gin.Context) {

	result, err := h.service.GetAllTweets()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
