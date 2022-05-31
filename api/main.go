package main

import (
	"database/sql"
	"os"
	"wgcapi/controllers"
	"wgcapi/logging"
	"wgcapi/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var db *sql.DB

func main() {
	logging.Info("Setup Starting...")
	Environment()
	DatabaseConnection()
	//Docs()
	MapRoutes()
}

func MapRoutes() {
	logging.Info("Mapping Routes")

	/* Define the router */
	router := gin.Default()

	// Index page with version?

	/* Wrestlers */
	router.GET("/wrestlers", controllers.GetAllWrestlers)

	router.Run(":8080")
}

func Docs() {

}

func DatabaseConnection() {
	logging.Info("Checking Database Connection")

	sqlDB := models.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	logging.Info("Database Connected!")

	/*

		cfg := mysql.Config{
			User:                 os.Getenv("DB_USERNAME"),
			Passwd:               os.Getenv("DB_PASSWORD"),
			Net:                  "tcp",
			Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
			DBName:               os.Getenv("DB_NAME"),
			AllowNativePasswords: true,
		}
		// Get a database handle.
		var err error
		db, err = sql.Open(os.Getenv("DB_CONNECTION"), cfg.FormatDSN())

		if err != nil {
			logging.Fatal(err.Error())
		}

		logging.Info("Database Connection Specifics are Ok!")

		pingErr := db.Ping()
		if pingErr != nil {
			logging.Fatal(pingErr.Error())
		}

	*/
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
