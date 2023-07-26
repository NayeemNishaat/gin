package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/controller"
	"gin/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetAllAlbums(t *testing.T) {
	r := GetRouter()

	r.GET("/api/v1/album", controller.GetAllAlbums)
	req, _ := http.NewRequest("GET", "/api/v1/album", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	// w.Body.Bytes()
	// jsonValue, _ := json.Marshal(responseData)
	var albums []model.Album
	json.Unmarshal(responseData, &albums)
	fmt.Println(len(albums))

	assert.NotEqual(t, albums, nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateAlbum(t *testing.T) {
	r := GetRouter()

	r.POST("/api/v1/album", controller.CreateAlbum)

	album := model.Album{ID: "4", Title: "Custom Albm", Artist: "ABC", Price: 30, URL: "Not Available"}
	stringifiedAlbum, _ := json.Marshal(album)

	req, _ := http.NewRequest("POST", "/api/v1/album", bytes.NewBuffer(stringifiedAlbum))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
