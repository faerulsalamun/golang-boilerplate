package server

import (
	"github.com/gin-gonic/gin"
	"github.com/faerulsalamun/golang-boilerplate/middlewares"
	"github.com/faerulsalamun/golang-boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		userController := new(controllers.UserController)

		v1.GET("/", userController.GetAll)

	}

	return router
}
