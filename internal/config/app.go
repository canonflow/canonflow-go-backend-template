package config

import (
	"canonflow-golang-backend-template/internal/controllers"
	"canonflow-golang-backend-template/internal/middlewares"
	"canonflow-golang-backend-template/internal/repositories"
	"canonflow-golang-backend-template/internal/routers"
	"canonflow-golang-backend-template/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *gin.Engine
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// TODO: Setup all repositories
	userRepository := repositories.NewUserRepository(config.Log)

	// TODO: Setup all services
	userService := services.NewUserService(config.DB, config.Log, config.Validate, config.Config, userRepository)

	// TODO: Setup all controllers
	userContoller := controllers.NewUserController(userService, config.Log)

	// TODO: Setup all middlewares
	authMiddleware := middlewares.AuthMiddleware(config.Config.GetString("JWT_SECRET_KEY"))

	routerConfig := &routers.RouterConfig{
		App:            config.App,
		UserController: userContoller,
		AuthMiddleware: &authMiddleware,
	}

	routerConfig.Setup()
}
