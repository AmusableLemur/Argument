package app

import (
	"net/http"

	"github.com/AmusableLemur/Argument/internal/config"
	"github.com/AmusableLemur/Argument/internal/database"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	if config.Conf.Test {
		r.LoadHTMLGlob("../../templates/*")
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": config.Conf.Title,
			"posts": database.GetPosts(),
		})
	})

	r.POST("/", func(c *gin.Context) {
		p := database.Post{Title: c.PostForm("title")}

		database.SavePost(p)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": config.Conf.Title,
			"posts": database.GetPosts(),
		})
	})

	return r
}
