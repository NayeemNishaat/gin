package service

import (
	"errors"
	"fmt"
	"gin/model"

	"github.com/gin-gonic/gin"
)

func GetAllAlbums() []model.Album {
	return model.Albums
}

func CreateAlbum(c *gin.Context) model.Album {
	// Call BindJSON to bind the received JSON to newAlbum.
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return newAlbum
	}

	// Add the new album to the slice.
	model.Albums = append(model.Albums, newAlbum)
	return newAlbum
}

func GetAlbumByID(id string) (album model.Album, err error) {
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range model.Albums {
		if album.ID == id {
			return album, err
		}
	}

	return album, errors.New("not found")
}
