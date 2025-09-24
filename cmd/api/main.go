package main

import (
	"strconv"

	"canonflow-golang-backend-template/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogrus(viperConfig)
	validate := config.NewValidator()
	app := config.NewGin(viperConfig, log)
	db := config.NewDatabase(viperConfig, log)

	// TODO: Bootstrap all configs
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	// TODO: Get the web port
	webPort := viperConfig.GetInt("WEB_PORT")

	if viperConfig.GetString("ENV") == "production" {
		log.Info("ON Production Environment")
	}

	err := app.Run(":" + strconv.Itoa(webPort))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
