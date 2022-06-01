package routes

import (
	"os"
	"wgcapi/controllers"
	"wgcapi/logging"

	"github.com/gin-gonic/gin"
)

func MapRoutes() *gin.Engine {
	logging.Info("Mapping Routes")

	/* Define the router */
	r := gin.Default()

	v1 := r.Group(os.Getenv("APP_VERSION"))
	{
		v1.GET("/wrestlers", controllers.GetAllWrestlers)
		v1.GET("/books", controllers.ListBook)
	}

	return r
}
