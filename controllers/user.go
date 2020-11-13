package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todoProject/models"
	"github.ru/GeorgVartanov/todoProject/storage"
)

// CreateUserController ...
func CreateUserController(c *gin.Context) {
	var newAPIUser models.User
	if err := c.ShouldBindJSON(&newAPIUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := newAPIUser.ValidateFields(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDB, err := storage.CreateUser(newAPIUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &userDB})
	return

}
