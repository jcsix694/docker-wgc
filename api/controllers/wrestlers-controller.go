package controllers

import (
	"wgcapi/helpers"
	"wgcapi/models"
	"wgcapi/requests"
	"wgcapi/structures"

	"github.com/gin-gonic/gin"
)

func CreateWrestler(c *gin.Context) {
	// Validate input
	var input requests.CreateWrestlerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		// Change error to show an array?
		helpers.RespondJSON(c, 400, nil, err.Error(), nil)
		return
	}

	wrestler, err := models.CreateWrestler(&input)

	if err != nil {
		helpers.RespondJSON(c, 400, wrestler, "Error", nil)
		return
	} else {
		helpers.RespondJSON(c, 200, wrestler, "Success", nil)
		return
	}
}

func GetAllWrestlers(c *gin.Context) {
	var wrestler []structures.Wrestler
	wrestler, err := models.GetAllWrestlers(&wrestler)
	if err != nil {
		helpers.RespondJSON(c, 400, wrestler, "Error", nil)
		return
	} else {
		helpers.RespondJSON(c, 200, wrestler, "Success", nil)
		return
	}
}
