package main

import (
	"database/sql"
	"log"
	"os"
	"wgcapi/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var db *sql.DB

func main() {
	Environment()
	DatabaseConnection()
	//Docs()
	MapRoutes()
}

func MapRoutes() {
	/* Define the router */
	router := gin.Default()

	/* Wrestlers */
	router.GET("/wrestlers", controllers.GetAllWrestlers)

	router.Run(":8080")
}

func Docs() {

}

func DatabaseConnection() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open(os.Getenv("DB_TYPE"), cfg.FormatDSN())

	if err != nil {
		log.Fatal("err")
	}

	logger.Info("Database Connection Specifics are Ok!")

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	logger.Info("Database Connected!")
}

func Environment() {
	/* Load the environment file */
	envErr := godotenv.Load()

	environment := os.Getenv("APP_ENVIRONMENT")

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Info("Environment Configured",
		// Structured context as strongly typed Field values.
		zap.String("appEnvironment", environment),
	)
}
