package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Http) GetFollowingTweets(c *gin.Context) {
	// Llamar a la lógica de negocio
	result, err := h.service.GetFollowingTweets()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
