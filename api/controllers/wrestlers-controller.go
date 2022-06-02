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
		// Change error to show an array of errors
		helpers.RespondJSON(c, 400, nil, err.Error(), nil)
		return
	}

	// Creates the Wrestler
	wrestler, err := models.CreateWrestler(&input)

	if err != nil {
		helpers.RespondJSON(c, 400, wrestler, "Error", nil)
	} else {
		helpers.RespondJSON(c, 200, wrestler, "Success", nil)
	}
}

func GetAllWrestlers(c *gin.Context) {
	var wrestler []structures.Wrestler
	wrestler, err := models.GetAllWrestlers(&wrestler)
	if err != nil {
		helpers.RespondJSON(c, 400, wrestler, "Error", nil)
	} else {
		helpers.RespondJSON(c, 200, wrestler, "Success", nil)
	}
}
