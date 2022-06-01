package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	status int
	meta   interface{}
	data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	var res ResponseData

	res.status = status
	res.data = payload

	w.JSON(200, res)
}
