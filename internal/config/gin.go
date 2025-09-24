package config

import (
	"fmt"
	"net/http"

	"canonflow-golang-backend-template/internal/models/web"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func PanicRecovery(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%s \"%s\" - Panic Occured: %+v", c.Request.Method, c.Request.URL, err)

				//! Return a unified error response
				c.JSON(http.StatusInternalServerError, web.ErrorResponse{
					Code:   http.StatusInternalServerError,
					Status: "Internal Server Error",
					Error:  fmt.Sprintf("%v", err),
				})

				c.Abort()
			}
		}()

		c.Next()
	}
}

func NewGin(config *viper.Viper, log *logrus.Logger) *gin.Engine {
	if viper.GetString("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	// TODO: CORS config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // your FE URL
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// TODO: Add Panic Middleware
	app.Use(PanicRecovery(log))

	// TODO: Make Not Found ENDPOINT Response
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, web.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Error:  "Endpoint Not Found",
		})
	})

	return app
}
