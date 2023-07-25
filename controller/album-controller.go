package controller

import (
	"bytes"
	"fmt"
	"gin/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Albums
// @Description Get The List Of Albums
// @Tags Album
// @Accept mpfd
// @Produce json
// @Success 200 {array} model.Album
// @Router /album [get]
// Note: @Param        q    query     string  false  "name search by q"  Format(email)
func GetAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAllAlbums())
}

// @Summary Create Album
// @Description Create An Album
// @Tags Album
// @Accept json
// @Produce json
// @Param album body model.Album true "Album Data"
// @Success 200 {object} model.Album
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
// @Router /album [post]
func CreateAlbum(c *gin.Context) {
	bb, _ := io.ReadAll(c.Request.Body) // Important: c.Request.Body will be empty if we read it
	c.Request.Body = io.NopCloser(bytes.NewReader(bb))
	fmt.Println(string(bb))

	// data := struct{}{} // Note: Inline struct

	c.IndentedJSON(http.StatusCreated, service.CreateAlbum(c))
}

// @Summary Get Album
// @Description Get An Album By Id
// @Tags Album
// @Produce json
// @Param id path int true "Album Id"
// @Success 200 {object} model.Album
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
// @Router /album/{id} [get]
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := service.GetAlbumByID(id)

	fmt.Println(album)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}
