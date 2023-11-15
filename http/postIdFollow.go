package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (h *Http) PostIdFollow(c *gin.Context) {
	id := c.Param("id")

	var followRequest entity.FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userToFollowID := followRequest.UserIDToFollow

	currentUser, err := h.service.GetAccountId(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener el usuario actual"})
		return
	}

	userToFollow, err := h.service.GetAccountId(userToFollowID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario a seguir no encontrado"})
		return
	}
	result, err := h.service.PostIdFollow(currentUser, userToFollow)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al seguir al usuario"})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario seguido correctamente", "result": result})
}
