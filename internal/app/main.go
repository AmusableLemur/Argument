package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(config Config) *gin.Engine {
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
