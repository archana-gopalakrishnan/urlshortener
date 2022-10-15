package handlers

import (
	"gin/shortener"
	storage "gin/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define shortening request model to contain longurl
type UrlShortenRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var shortenRequest UrlShortenRequest
	if err := ctx.ShouldBindJSON(&shortenRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(shortenRequest.LongUrl)
	storage.SaveUrlMapping(shortUrl, shortenRequest.LongUrl)

	host := "http://localhost:8080/"
	ctx.JSON(200, gin.H{
		"message":   "url shortened successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := storage.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
