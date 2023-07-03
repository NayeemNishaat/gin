package controller

import (
	"bytes"
	"fmt"
	"gin/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAllAlbums())
}

// postAlbums adds an album from JSON received in the request body.
func CreateAlbum(c *gin.Context) {
	bb, _ := io.ReadAll(c.Request.Body) // Important: c.Request.Body will be empty if we read it
	c.Request.Body = io.NopCloser(bytes.NewReader(bb))
	fmt.Println(string(bb))

	// data := struct{}{} // Note: Inline struct

	c.IndentedJSON(http.StatusCreated, service.CreateAlbum(c))
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
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
