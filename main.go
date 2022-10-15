package main

import (
	"fmt"
	handler "gin/handlers"
	"gin/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "Hi! You have reached the URL Shortener API"}) })

	r.POST("/create-short-url", func(ctx *gin.Context) { handler.CreateShortUrl(ctx) })

	r.GET("/:shortUrl", func(ctx *gin.Context) { handler.HandleShortUrlRedirect(ctx) })

	// Note that store initialization happens here
	storage.InitializeStore()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
