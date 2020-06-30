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

func setupConfig(f string) Config {
	stream, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	return LoadConfig(stream)
}

func main() {
	config = setupConfig("config.toml")
	setupRouter().Run(":8080")
}
