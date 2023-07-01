package urlShortenerApi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectController(c *gin.Context) {
	hashedPath := c.Param("hashedPath")
	destinationUrl, err := getDestinationUrlByHash(hashedPath)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unknown error",
		})
	}
	c.Redirect(http.StatusFound, destinationUrl)
}

func CreateShortURLController(c *gin.Context) {
	var requestBody createShortURLRequest
	err := c.BindJSON(&requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "improper request",
		})

		return
	}

	isValidURL := validateURL(requestBody.Url)

	if isValidURL != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "url is invalid",
				"error":   isValidURL.Error(),
			},
		)

		return
	}

	shortenedRecord, err := createUrlShortenerRecord(requestBody.Url)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	shortedUrl := fmt.Sprintf("https://myShortener.xyz/%s", shortenedRecord.hash)

	c.JSON(
		http.StatusOK,
		gin.H{
			"url": shortedUrl,
		},
	)

	return
}
