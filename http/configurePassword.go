package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (h *Http) ConfigurePassword(c *gin.Context) {
	id := c.Param("id")
	var user entity.User

	result, err := h.service.ConfigurePassword(id, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
