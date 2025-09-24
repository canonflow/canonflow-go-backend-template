package routers

import (
	"canonflow-golang-backend-template/internal/controllers"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	App            *gin.Engine
	UserController *controllers.UserController
	AuthMiddleware *gin.HandlerFunc
}

func (c *RouterConfig) Setup() {
	c.SetupGuestRouter()
	c.SetupAuthRouter()
}

func (c *RouterConfig) SetupGuestRouter() {
	// TODO: Setup Login Endpoint

	auth := c.App.Group("auth")
	{
		auth.POST("/login", c.UserController.Login)
	}
}

func (c *RouterConfig) SetupAuthRouter() {
	// TODO: Declare all endpoints with Authentication
	c.App.Use(*c.AuthMiddleware)

	// TODO: Setup Auth endpoints
	auth := c.App.Group("auth")
	{
		auth.POST("/logout", c.UserController.Logout)
	}
}
