package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/entity"
)

func (h *Http) PostIdUnfollow(c *gin.Context) {
	id := c.Param("id")

	var unFollowRequest entity.FollowRequest
	if err := c.ShouldBindJSON(&unFollowRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userToFollowID := unFollowRequest.UserIDToFollow

	currentUser, err := h.service.GetAccountId(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener el usuario actual"})
		return
	}

	userToFollow, err := h.service.GetAccountId(userToFollowID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario a dejar de seguir no encontrado"})
		return
	}
	result, err := h.service.PostIdUnfollow(currentUser, userToFollow)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al dejar de seguir al usuario"})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario dejado de seguir correctamente", "result": result})
}
