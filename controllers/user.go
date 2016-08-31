package controllers

import (
	"github.com/faerulsalamun/golang-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) GetAll(c *gin.Context) {
	user, err := userModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": user})
	return
}
