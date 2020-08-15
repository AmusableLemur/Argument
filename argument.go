package main

import (
	"github.com/AmusableLemur/Argument/internal/app"
	"github.com/AmusableLemur/Argument/internal/config"
)

var conf config.Config

func main() {
	conf, _ := config.LoadConfig("config.toml")
	app.SetupRouter(conf).Run(":8080")
}
