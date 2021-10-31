package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skl98/web-service-gin/model"
)

func main() {
	router := gin.Default()
	router.GET("/albums", model.GetAlbums)
	router.GET("/albums/:id", model.GetAlbumById)
	router.POST("/albums", model.PostAlbums)

	router.Run("localhost:8080")
}
