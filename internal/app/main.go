package app

import (
	"net/http"

	"github.com/AmusableLemur/Argument/internal/config"
	"github.com/AmusableLemur/Argument/internal/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config config.Config) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": config.Title,
			"posts": database.GetPosts(),
		})
	})

	return r
}
