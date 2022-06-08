package controllers

import (
	"github.com/google/uuid"
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

	// Check if the wrestler already exists by name

	// Creates the Wrestler
	wrestler, err := models.CreateWrestlerData(&input)

	if err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, nil, err, nil)
	} else {
		helpers.RespondJSON(c, http.StatusCreated, wrestler, nil, nil)
	}
}

func GetAllWrestlers(c *gin.Context) {
	var wrestler []structures.Wrestler
	wrestler, err := models.GetAllWrestlersData(&wrestler)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, wrestler, err, nil)
	} else {
		helpers.RespondJSON(c, http.StatusOK, wrestler, nil, nil)
	}
}

func GetWrestlerByUuid(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("uuid"))

	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, nil, err, nil)
		return
	}

	mapData := requests.GetWrestlerRequest{UUID: uuid}
	wrestler, err := models.GetWrestlerData(mapData)

	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, nil, err, nil)
	} else {
		helpers.RespondJSON(c, http.StatusOK, wrestler, nil, nil)
	}
}
