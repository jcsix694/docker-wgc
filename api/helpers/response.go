package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func RespondJSON(w *gin.Context, status int, payload interface{}, message string, meta interface{}) {
	var res ResponseData

	res.Status = status
	res.Meta = meta
	res.Data = payload
	res.Message = message

	w.JSON(status, res)
}
