package main

import (
	"fmt"
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

// Config specifies the availabel configuration values
type Config struct {
	Title    string
	Database struct {
		Username string
		Password string
		Name     string
		Host     string
		Port     int
		URI      string
	}
}

// LoadConfig loads a config from a file
func LoadConfig(f string) Config {
	stream, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	return SetupConfig(stream)
}

// SetupConfig converts a byte stream to a Config object
func SetupConfig(stream []byte) Config {
	config := Config{}
	toml.Unmarshal(stream, &config)

	if "" == config.Title {
		config.Title = "Argument"
	}

	config.Database.URI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	return config
}
