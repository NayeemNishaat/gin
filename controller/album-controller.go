package controller

import (
	"bytes"
	"fmt"
	"gin/model"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	bb, _ := io.ReadAll(c.Request.Body) // Important: c.Request.Body will be empty if we read it
	c.Request.Body = io.NopCloser(bytes.NewReader(bb))
	fmt.Println(string(bb))

	// data := struct{}{} // Note: Inline struct

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return
	}

	// Add the new album to the slice.
	model.Albums = append(model.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range model.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
