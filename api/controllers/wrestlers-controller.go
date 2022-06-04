package controllers

import (
	"net/http"
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
		helpers.RespondJSON(c, http.StatusBadRequest, nil, err, nil)
		return
	}

	// Creates the Wrestler
	wrestler, err := models.CreateWrestler(&input)

	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, nil, err, nil)
	} else {
		helpers.RespondJSON(c, http.StatusCreated, wrestler, nil, nil)
	}
}

func GetAllWrestlers(c *gin.Context) {
	var wrestler []structures.Wrestler
	wrestler, err := models.GetAllWrestlers(&wrestler)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, wrestler, err, nil)
	} else {
		helpers.RespondJSON(c, http.StatusOK, wrestler, nil, nil)
	}
}
