package main

import (
	"io/ioutil"
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
		})
	})

	return r
}

func main() {
	stream, err := ioutil.ReadFile("config.toml")

	if err != nil {
		panic(err)
	}

	config = LoadConfig(stream)

	setupRouter().Run(":8080")
}
