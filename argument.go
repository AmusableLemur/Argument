package main

import (
	"github.com/AmusableLemur/Argument/internal/app"
	"github.com/AmusableLemur/Argument/internal/config"
)

var conf config.Config

func main() {
	app.SetupRouter().Run(":8080")
}
