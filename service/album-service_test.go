package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllAlbums(t *testing.T) {
	albums := GetAllAlbums()

	if len(albums) == 0 {
		t.Errorf("No Albums Found")
	}
}

func TestCreateAlbum(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	album, err := json.Marshal(map[string]any{"ID": "4", "Title": "Custom Albm", "Artist": "ABC", "Price": 30, "URL": "Not Available"})

	if err != nil {
		panic(err)
	}

	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/album", bytes.NewBuffer(album)) // Important: Initially c.Request is nil. So first we need to create request with http package.
	// c.Request.Body = io.NopCloser(bytes.NewBuffer(album)) // Note: Setting a new body!

	createdAlbum := CreateAlbum(c)
	fmt.Println(createdAlbum)
}
