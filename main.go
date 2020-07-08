package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var config Config

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": config.Title,
			"posts": GetPosts(),
		})
	})

	return r
}

func main() {
	config = LoadConfig("config.toml")

	Connect(config.Database.URI)
	setupRouter().Run(":8080")
}
