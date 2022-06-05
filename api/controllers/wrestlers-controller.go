package controllers

import (
	"fmt"
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
		fmt.Println(err)
		helpers.RespondJSON(c, http.StatusBadRequest, nil, err, nil)
		return
	}

	// Check if the wrestler already exists

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

func GetWrestler() {

}
