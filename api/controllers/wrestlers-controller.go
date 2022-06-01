package controllers

import (
	"wgcapi/helpers"
	"wgcapi/models"
	"wgcapi/structures"

	"github.com/gin-gonic/gin"
)

func GetAllWrestlers(c *gin.Context) {
	var wrestler []structures.Wrestler
	wrestler, err := models.GetAllWrestlers(&wrestler)
	if err != nil {
		helpers.RespondJSON(c, 400, wrestler, "Error", nil)
	} else {
		helpers.RespondJSON(c, 200, wrestler, "Success", nil)
	}
}
