package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/recepturker/accounting-app/internal/models"
	"github.com/recepturker/accounting-app/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı kaydedilemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla kaydedildi"})
}
