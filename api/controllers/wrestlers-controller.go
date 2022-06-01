package controllers

import (
	"wgcapi/helpers"
	"wgcapi/logging"
	"wgcapi/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAllWrestlers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func ListBook(c *gin.Context) {
	var book []models.Book
	logging.Info("Trying to get all books")
	book, err := models.GetAllBook(&book)
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}
