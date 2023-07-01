package main

import (
	"brookatlas/urlShortenerApi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.POST(
		"createShortURL", urlShortenerApi.CreateShortURLController,
	)
	router.GET("/r/:hashedPath", urlShortenerApi.RedirectController)

	router.Run()
}
