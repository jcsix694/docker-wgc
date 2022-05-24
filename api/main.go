package main

import (
	"database/sql"
	"os"
	"wgcapi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	/* Load the environment file */
	envErr := godotenv.Load()

	environment := os.Getenv("APP_ENVIRONMENT")

	if environment == "local" {
		if envErr != nil {
			panic("Error loading .env file")
		}
	}

	/* Connecting to the database */
	conn, dbErr := sql.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if dbErr != nil {
		panic("Error connecting to database")
	}

	defer conn.Close()

	mapRoutes()
}

func mapRoutes() {
	/* Define the router */
	router := gin.Default()

	/* Wrestlers */
	router.GET("/wrestlers", controllers.GetAllWrestlers)

	router.Run(":8080")
}
