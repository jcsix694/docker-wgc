package main

import (
	"os"
	"wgcapi/config"
	"wgcapi/logging"
	"wgcapi/routes"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logging.Info("Setup Starting...")
	Environment()
	config.DBInit()
	//Docs()
	r := routes.MapRoutes()

	// running
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
