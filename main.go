package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/laulfred/go-url-shortener/handler"
	"github.com/laulfred/go-url-shortener/store"
)

func main() {
	r := gin.Default()

	//FRONTEND
	//CORS
	r.Use(cors.Default())
	//index.html
	r.StaticFile("/", "./frontend/index.html")

	/*r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API!",
		})
	})*/

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
