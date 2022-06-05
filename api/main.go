package main

import (
	"os"
	"wgcapi/database"
	"wgcapi/helpers"
	"wgcapi/logging"
	"wgcapi/routes"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logging.Info("Setup Starting...")
	Environment()
	database.Connect()
	database.AutoMigrate()
	r := routes.MapRoutes()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		helpers.ConfigureValidationJsonFields(v)
	}

	//Docs()
	r.Run(":" + os.Getenv("APP_PORT"))
}

func Docs() {

}

func Environment() {
	logging.Info("Checking Environment")

	/* Load the environment file */
	envErr := godotenv.Load()

	environment := os.Getenv("APP_ENVIRONMENT")

	if envErr != nil {
		logging.Fatal("Error loading .env file")
	}

	logging.Info("Environment Configured",
		// Structured context as strongly typed Field values.
		zap.String("appEnvironment", environment),
	)
}
